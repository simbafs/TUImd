package util

import (
	Msg "tuimd/msg"

	MDrender "github.com/MichaelMure/go-term-markdown"
	tea "github.com/charmbracelet/bubbletea"
)

func RenderMD(md string, width int) tea.Cmd {
	return func() tea.Msg {
		rendered := MDrender.Render(md, width, 0)

		return Msg.Rendered(string(rendered))
	}
}
