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
	width    int
	height   int
	nodes    []tea.Model
}

func NewModel() model {
	tab := component.NewTab()
	source := component.NewSouce("loading file...")
	markdown := component.Markdown("Markdown Editor\nhifdasjf jsdklafjkl ajfklwjefjds")
	cmd := component.NewCmd()

	m := model{
		nodes: make([]tea.Model, 4),
	}

	m.nodes[tabAddr] = tab
	m.nodes[sourceAddr] = source
	m.nodes[markdownAddr] = markdown
	m.nodes[cmdAddr] = cmd

	return m
}

func (m model) Init() tea.Cmd {
	var cmd tea.Cmd
	var cmds []tea.Cmd = make([]tea.Cmd, 4)

	for _, v := range m.nodes {
		cmd = v.Init()
		cmds = append(cmds, cmd)
	}

	cmds = append(cmds, util.ReadFile("./test.md"))

	return tea.Batch(cmds...)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd = make([]tea.Cmd, 4)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if keymap.Matches(msg, keymap.Quit) {
			cmds = append(cmds, func() tea.Msg {
				return component.UpdateMsgMsg("Type  :q  and press <Enter> to exit TUImd")
			})
		}
	case tea.WindowSizeMsg:
		m.height, m.width = msg.Height, msg.Width
	}

	for addr, v := range m.nodes {
		m.nodes[addr], cmd = v.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	// TODO: switch to github.com/treilik/bubbleboxer

	middle := util.SplitHorizontal(m.height-4, []int{m.width/2 - 1, m.width/2 - 1},
		m.nodes[sourceAddr].View(),
		m.nodes[markdownAddr].View(),
	)

	return util.SplitVertical(m.width, []int{1, lipgloss.Height(middle), 1},
		m.nodes[tabAddr].View(),
		middle,
		m.nodes[cmdAddr].View(),
	)
}

func main() {
	if _, err := tea.NewProgram(NewModel(), tea.WithAltScreen()).Run(); err != nil {
		fmt.Printf("Oops, there's some error: %v\n", err)
		os.Exit(1)
	}
}
