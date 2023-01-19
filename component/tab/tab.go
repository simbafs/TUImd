package tab

import (
	Msg "tuimd/msg"

	tea "github.com/charmbracelet/bubbletea"
)

type Tab struct {
	Text string
}

func (m Tab) Init() tea.Cmd { return nil }
func (m Tab) Update(msg tea.Msg) (Tab, tea.Cmd) {
	// var cmd tea.Cmd
	// var cmds []tea.Cmd = make([]tea.Cmd, 4)

	switch msg := msg.(type) {
	case Msg.FilenameChange:
		m.Text = string(msg)
	}

	return m, nil
	// return m, tea.Batch(cmds...)
}

func (m Tab) View() string { return m.Text }

func New() Tab {
	return Tab{
		Text: "",
	}
}
