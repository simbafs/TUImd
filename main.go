package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	keymap "github.com/simbafs/TUImd/keyMap"
	"github.com/simbafs/TUImd/util"
)

const (
	tabAddr = iota
	sourceAddr
	markdownAddr
	cmdAddr
)

type Mode int

const (
	NormalMode Mode = iota
	InsertMode
)

var mode = NormalMode

// main model
type model struct {
	filename string
	body     string
	mode     Mode
	width    int
	height   int
	tab      Tab
	source   Source
	markdown Markdown
	cmd      Cmd
}

func NewModel() model {
	tab := NewTab()
	source := NewSouce("loading file...")
	markdown := Markdown("Markdown Editor\nhifdasjf jsdklafjkl ajfklwjefjds")
	cmd := NewCmd()

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
				return UpdateMsgMsg("Type  :q  and press <Enter> to exit TUImd")
			})
			// there should be a more elegant way to manage state
			// case keymap.Matches(msg, keymap.BeginInsertMode):
			// 	if mode == NormalMode {
			// 		mode = InsertMode
			// 		cmds = append(cmds, func() tea.Msg {
			// 			return InsertMode
			// 		})
			// 	}
			// case keymap.Matches(msg, keymap.BeginNormalMode):
			// 	if mode == InsertMode && m.cmd.Cmding == false {
			// 		mode = NormalMode
			// 		cmds = append(cmds, func() tea.Msg {
			// 			return NormalMode
			// 		})
			// 	}
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
