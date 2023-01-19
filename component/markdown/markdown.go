package markdown

import (
	Msg "tuimd/msg"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type Markdown struct {
	viewport viewport.Model
	rendered string
}

func (m Markdown) Init() tea.Cmd { return m.viewport.Init() }

func (m Markdown) Update(msg tea.Msg) (Markdown, tea.Cmd) {
	switch msg := msg.(type) {
	case Msg.Rendered:
		m.rendered = string(msg)
	case tea.WindowSizeMsg:
		m.viewport = viewport.New(msg.Width/2-1, msg.Height-4)
	}

	return m, nil
}

func (m Markdown) View() string { return m.rendered }

func New() Markdown {
	return Markdown{
		viewport: viewport.New(10, 10),
	}
}
