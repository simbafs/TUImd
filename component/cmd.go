package component

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	keymap "github.com/simbafs/TUImd/keyMap"
	"github.com/simbafs/TUImd/util"
)

type Cmd struct {
	textinput textinput.Model
	msg       string
}

type UpdateMsgMsg string

func (m Cmd) Init() tea.Cmd { return textinput.Blink }
func (m Cmd) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd = make([]tea.Cmd, 0)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.textinput.Width = msg.Width
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keymap.CommandPrefix):
			m.msg = ""
			m.textinput.Focus()
		case key.Matches(msg, keymap.EnterCommand):
			m.textinput.Blur()
			cmd = util.CmdExec(m.textinput.Value())
			m.textinput.SetValue("")
			cmds = append(cmds, cmd)
		}
	case UpdateMsgMsg:
		m.msg = string(msg)
	}

	m.textinput, cmd = m.textinput.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m Cmd) View() string {
	if m.msg != "" {
		// gray
		return m.msg
	}
	return m.textinput.Value()
}

func NewCmd() Cmd {
	return Cmd{
		textinput: textinput.New(),
	}
}

var Blink = textinput.Blink
