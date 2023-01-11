package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/simbafs/TUImd/component"
	keymap "github.com/simbafs/TUImd/keyMap"
	"github.com/simbafs/TUImd/util"
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
	mode     component.Mode
	width    int
	height   int
	tab      component.Tab
	source   component.Source
	markdown component.Markdown
	cmd      component.Cmd
}

func NewModel() model {
	tab := component.NewTab()
	source := component.NewSouce("loading file...")
	markdown := component.Markdown("Markdown Editor\nhifdasjf jsdklafjkl ajfklwjefjds")
	cmd := component.NewCmd()

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
	case tea.KeyMsg:
		switch {
		case keymap.Matches(msg, keymap.Quit):
			cmds = append(cmds, func() tea.Msg {
				return component.UpdateMsgMsg("Type  :q  and press <Enter> to exit TUImd")
			})
		// there should be a more elegant way to manage state
		case keymap.Matches(msg, keymap.BeginInsertMode):
			if m.mode == component.NormalMode {
				m.mode = component.InsertMode
				m.source.Mode = component.InsertMode
				m.cmd.Mode = component.InsertMode
			}
		case keymap.Matches(msg, keymap.BeginNormalMode):
			if m.mode == component.InsertMode {
				m.mode = component.NormalMode
				m.source.Mode = component.NormalMode
				m.cmd.Mode = component.NormalMode
			}
		}

	case tea.WindowSizeMsg:
		m.height, m.width = msg.Height, msg.Width
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
	if _, err := tea.NewProgram(NewModel(), tea.WithAltScreen()).Run(); err != nil {
		fmt.Printf("Oops, there's some error: %v\n", err)
		os.Exit(1)
	}
}
