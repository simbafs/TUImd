package main

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	keymap "github.com/simbafs/TUImd/keyMap"
	"github.com/simbafs/TUImd/util"
)

type Cmd struct {
	Textinput textinput.Model
	Msg       string
}

type UpdateMsgMsg string

func (m Cmd) Init() tea.Cmd { return textinput.Blink }
func (m Cmd) Update(msg tea.Msg) (Cmd, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd = make([]tea.Cmd, 0)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Textinput.Width = msg.Width
	case Mode:
		if mode == InsertMode {
			m.Textinput.SetValue("")
			m.Textinput.Blur()
		}
	case tea.KeyMsg:
		if mode == NormalMode {
			switch {
			case key.Matches(msg, keymap.CommandPrefix):
				m.Msg = ""
				m.Textinput.Focus()
			case key.Matches(msg, keymap.EnterCommand):
				m.Textinput.Blur()
				cmd = util.CmdExec(m.Textinput.Value())
				cmds = append(cmds, cmd)
				m.Textinput.SetValue("")
			case key.Matches(msg, keymap.BeginInsertMode):
				if m.Textinput.Value() != "" {
					break
				}
				mode = InsertMode
				cmds = append(cmds, func() tea.Msg {
					return InsertMode
				})
			case m.Textinput.Value() == "":
				m.Textinput.Blur()
			}
		}
	case UpdateMsgMsg:
		m.Msg = string(msg)
	}

	m.Textinput, cmd = m.Textinput.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m Cmd) View() string {
	if m.Msg != "" {
		// gray
		return m.Msg
	}
	return m.Textinput.Value()
}

func NewCmd() Cmd {
	return Cmd{
		Textinput: textinput.New(),
	}
}

var Blink = textinput.Blink
