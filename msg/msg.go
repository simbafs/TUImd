package msg

import tea "github.com/charmbracelet/bubbletea"

type ModeChange string
type FilenameChange string
type BodyChange string

type ShowMsg string
type Rendered string

func ChangeMode(mode string) func() tea.Msg {
	return func() tea.Msg {
		return ModeChange(mode)
	}
}

func ChangeFilename(filename string) func() tea.Msg {
	return func() tea.Msg {
		return FilenameChange(filename)
	}
}

func ChangeBody(body string) func() tea.Msg {
	return func() tea.Msg {
		return BodyChange(body)
	}
}
