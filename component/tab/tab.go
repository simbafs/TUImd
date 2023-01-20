package tab

import (
	"fmt"
	Msg "tuimd/msg"

	tea "github.com/charmbracelet/bubbletea"
)

type Tab struct {
	filename string
	mode     string
}

func (m Tab) Init() tea.Cmd { return nil }
func (m Tab) Update(msg tea.Msg) (Tab, tea.Cmd) {
	switch msg := msg.(type) {
	case Msg.Mode:
		m.mode = string(msg)
	case Msg.Filename:
		m.filename = string(msg)
	case Msg.SaveFile:
		if msg == "" {
			if m.filename == "" {
				return m, Msg.NewCmd[Msg.ShowMsg]("No file name")
			} else {
				return m, Msg.NewCmd[Msg.SaveFile](m.filename)
			}
		}
	}

	return m, nil
}

func (m Tab) View() string {
	return fmt.Sprintf("%s | mode: %s", m.filename, m.mode)
}

func New() Tab {
	return Tab{
		filename: "",
	}
}
