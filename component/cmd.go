package component

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	keymap "github.com/simbafs/TUImd/keyMap"
)

type Cmd struct {
	Textinput textinput.Model
}

func (m Cmd) Init() tea.Cmd { return textinput.Blink }
func (m Cmd) Update(msg tea.Msg) (Cmd, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Textinput.Width = msg.Width
	case tea.KeyMsg:
		if key.Matches(msg, keymap.CommandPrefix) {
			m.Textinput.Focus()
		}
	}

	m.Textinput, cmd = m.Textinput.Update(msg)

	return m, cmd
}
func (m Cmd) View() string { return m.Textinput.Value() }

func NewCmd() Cmd {
	return Cmd{
		Textinput: textinput.New(),
	}
}

var Blink = textinput.Blink
