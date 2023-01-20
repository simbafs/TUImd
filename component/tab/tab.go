package tab

import (
	"fmt"
	Msg "tuimd/msg"

	tea "github.com/charmbracelet/bubbletea"
)

type Tab struct {
	text string
	mode string
}

func (m Tab) Init() tea.Cmd { return nil }
func (m Tab) Update(msg tea.Msg) (Tab, tea.Cmd) {
	switch msg := msg.(type) {
	case Msg.ModeChange:
		m.mode = string(msg)
	case Msg.FilenameChange:
		m.text = string(msg)
	}

	return m, nil
}

func (m Tab) View() string {
	return fmt.Sprintf("%s | mode: %s", m.text, m.mode)
}

func New() Tab {
	return Tab{
		text: "",
	}
}
