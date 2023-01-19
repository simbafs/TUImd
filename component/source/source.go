package source

import (
	"fmt"
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
		// m.Textarea.SetWidth(msg.Width/2 - 1)
		util.Log(fmt.Sprintf("height: %d, width: %d\n", msg.Height, msg.Width))
		util.Log(fmt.Sprintf("\th: %d, w: %d\n", m.Textarea.Height(), m.Textarea.Width()))
	case Msg.BodyChange:
		m.Textarea.SetValue(string(msg))
		m.Textarea.Focus()
		cmds = append(cmds, util.RenderMD(string(msg)))
	case Msg.ModeChange:
		if msg == "insert" {
			m.Textarea.Focus()
		} else if msg == "normal" {
			m.Textarea.Blur()
		}
	}

	m.Textarea, cmd = m.Textarea.Update(msg)
	cmds = append(cmds, cmd)

	if _, ok := msg.(tea.KeyMsg); ok {
		if m.Textarea.Focused() {
			cmds = append(cmds, util.RenderMD(m.Textarea.Value()))
		}
	}

	return m, tea.Batch(cmds...)
}

func (m Source) View() string { return m.Textarea.View() }

func New(text string) Source {
	input := textarea.New()
	input.CharLimit = 0
	return Source{
		Textarea: input,
	}
}
