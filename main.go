package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"tuimd/component/cmd"
	"tuimd/component/markdown"
	"tuimd/component/source"
	"tuimd/component/tab"
	keymap "tuimd/keyMap"
	Msg "tuimd/msg"
	"tuimd/util"
)

const (
	tabAddr = iota
	sourceAddr
	markdownAddr
	cmdAddr
)

// main model
type model struct {
	filename string
	body     string
	width    int
	height   int
	tab      tab.Tab
	source   source.Source
	markdown markdown.Markdown
	cmd      cmd.Cmd
}

func NewModel() model {
	tab := tab.New()
	source := source.New("loading file...")
	markdown := markdown.New()
	cmd := cmd.New()

	m := model{
		tab:      tab,
		source:   source,
		markdown: markdown,
		cmd:      cmd,
	}

	return m
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		Msg.NewCmd[Msg.Mode]("normal"),
		m.tab.Init(),
		m.source.Init(),
		m.markdown.Init(),
		m.cmd.Init(),
		util.ReadFile("./test.md"),
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd = make([]tea.Cmd, 4)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height, m.width = msg.Height, msg.Width
	case tea.KeyMsg:
		switch {
		case keymap.Matches(msg, keymap.Quit):
			// cmds = append(cmds, tea.Quit)
			cmds = append(cmds, func() tea.Msg {
				return Msg.ShowMsg("Type  :q  and press <Enter> to exit TUImd")
			})
		case keymap.Matches(msg, keymap.BeginNormalMode):
			cmds = append(cmds, Msg.NewCmd[Msg.Mode]("normal"))
		}
	}

	m.tab, cmd = m.tab.Update(msg)
	cmds = append(cmds, cmd)
	m.source, cmd = m.source.Update(msg)
	cmds = append(cmds, cmd)
	m.markdown, cmd = m.markdown.Update(msg)
	cmds = append(cmds, cmd)
	m.cmd, cmd = m.cmd.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	// TODO: switch to github.com/treilik/bubbleboxer

	// TODO: the problem of height of source component
	middle := util.SplitHorizontal(m.height-4, []int{m.width/2 - 1, m.width/2 - 1},
		m.source.View(),
		m.markdown.View(),
	)

	return util.SplitVertical(m.width, []int{1, lipgloss.Height(middle), 1},
		m.tab.View(),
		middle,
		m.cmd.View(),
	)
}

func main() {
	if _, err := tea.NewProgram(NewModel(),
		tea.WithAltScreen(),
		tea.WithMouseAllMotion(),
	).Run(); err != nil {
		fmt.Printf("Oops, there's some error: %v\n", err)
		os.Exit(1)
	}
}
