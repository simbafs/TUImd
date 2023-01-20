package cmd

import (
	keymap "tuimd/keyMap"
	Msg "tuimd/msg"
	"tuimd/util"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Event func(string)

type Cmd struct {
	textinput    textinput.Model
	msg          string
	mode         string
	isCommanding bool
}

func (m Cmd) Init() tea.Cmd { return textinput.Blink }
func (m Cmd) Update(msg tea.Msg) (Cmd, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd = make([]tea.Cmd, 0)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.textinput.Width = msg.Width
	case Msg.ModeChange:
		m.mode = string(msg)
		m.textinput.Blur()
		m.textinput.SetValue("")
		m.msg = ""
		if msg == "normal" {
			m.isCommanding = false
		}
	case Msg.ShowMsg:
		m.msg = string(msg)
	case tea.KeyMsg:
		switch {
		case keymap.Matches(msg, keymap.CommandPrefix):
			m.isCommanding = true
			m.textinput.Focus()
		case keymap.Matches(msg, keymap.EnterCommand):
			m.isCommanding = false
			cmds = append(cmds, util.CmdExec(m.textinput.Value()))
			m.textinput.Blur()
			m.textinput.SetValue("")
		case keymap.Matches(msg, keymap.BeginInsertMode):
			if !m.isCommanding {
				cmds = append(cmds, Msg.ChangeMode("insert"))
			}
		}
	}

	m.textinput, cmd = m.textinput.Update(msg)
	cmds = append(cmds, cmd)

	if m.textinput.Value() == "" {
		m.textinput.Blur()
		m.isCommanding = false
	}

	return m, tea.Batch(cmds...)
}

func (m Cmd) View() string {
	if m.msg != "" {
		return m.msg
	}
	if m.isCommanding {
		return m.textinput.View()
	}
	return ""
}

func New() Cmd {
	text := textinput.New()
	text.Prompt = ""
	return Cmd{
		textinput: text,
		msg:       "",
	}
}
