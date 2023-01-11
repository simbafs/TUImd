package component

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/simbafs/TUImd/util"
)

type Tab struct {
	text string
}

func (m Tab) Init() tea.Cmd { return nil }
func (m Tab) Update(msg tea.Msg) (Tab, tea.Cmd) {
	// var cmd tea.Cmd
	// var cmds []tea.Cmd = make([]tea.Cmd, 4)

	switch msg := msg.(type) {
	case util.FileMsg:
		m.text = msg.Filename
	}

	return m, nil
	// return m, tea.Batch(cmds...)
}
func (m Tab) View() string { return m.text }

func NewTab() Tab {
	return Tab{
		text: "",
	}
}
