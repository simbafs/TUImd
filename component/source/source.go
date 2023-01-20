package source

import (
	Msg "tuimd/msg"
	"tuimd/util"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

type Source struct {
	Textarea textarea.Model
}

func (m Source) Init() tea.Cmd {
	return textarea.Blink
}

func (m Source) Update(msg tea.Msg) (Source, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd = make([]tea.Cmd, 4)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Textarea.SetHeight(msg.Height - 4)
		m.Textarea.SetWidth(msg.Width/2 - 1)
		cmds = append(cmds, util.RenderMD(m.Textarea.Value(), m.Textarea.Width()))
	case Msg.Body:
		m.Textarea.SetValue(string(msg))
		cmds = append(cmds, util.RenderMD(m.Textarea.Value(), m.Textarea.Width()))
	case Msg.Mode:
		if msg == "insert" {
			m.Textarea.Focus()
		} else if msg == "normal" {
			m.Textarea.Blur()
		}
	case Msg.SaveFile:
		if msg != "" {
			cmds = append(cmds, util.SaveFile(string(msg), m.Textarea.Value()))
		}
	}

	// disable mouse event for textarea
	if _, ok := msg.(tea.MouseEvent); !ok {
		m.Textarea, cmd = m.Textarea.Update(msg)
		cmds = append(cmds, cmd)
	}

	if _, ok := msg.(tea.KeyMsg); ok && m.Textarea.Focused() {
		cmds = append(cmds, util.RenderMD(m.Textarea.Value(), m.Textarea.Width()))
	}

	return m, tea.Batch(cmds...)
}

func (m Source) View() string { return m.Textarea.View() }

func New(text string) Source {
	input := textarea.New()
	input.CharLimit = 0
	input.SetValue(text)
	return Source{
		Textarea: input,
	}
}
