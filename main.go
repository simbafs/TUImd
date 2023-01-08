package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/simbafs/TUImd/component"
	keymap "github.com/simbafs/TUImd/keyMap"
	"github.com/simbafs/TUImd/util"
)

// main model
type model struct {
	filename string
	body     string
	width    int
	height   int
	tab      component.Tab
	source   component.Source
	markdown component.Markdown
	cmd      component.Cmd
}

func NewModel() model {
	tab := component.Tab("tab a | tab b")
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

func readFile(filename string) tea.Cmd {
	return func() tea.Msg {
		b, err := ioutil.ReadFile(filename)
		if err != nil {
			return component.SourceMsg("fail to read " + filename)
		}

		return component.SourceMsg(b)
	}
}

func (m model) Init() tea.Cmd {
	var cmds []tea.Cmd = make([]tea.Cmd, 4)
	cmds = append(cmds,
		m.tab.Init(),
		m.source.Init(),
		m.markdown.Init(),
		m.cmd.Init(),
		readFile("./test.md"),
	)

	return tea.Batch(cmds...)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd = make([]tea.Cmd, 4)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keymap.Quit):
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.height, m.width = msg.Height, msg.Width
		// m.tui.UpdateSize(msg)
	}

	m.tab, cmds[0] = m.tab.Update(msg)
	m.source, cmds[1] = m.source.Update(msg)
	m.markdown, cmds[2] = m.markdown.Update(msg)
	m.cmd, cmds[3] = m.cmd.Update(msg)

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
