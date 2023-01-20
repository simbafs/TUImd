package msg

import tea "github.com/charmbracelet/bubbletea"

type Mode string
type Filename string
type Body string

type ShowMsg string
type Rendered string

// see README.md
type SaveFile string

func NewCmd[V ~string](val string) func() tea.Msg {
	return func() tea.Msg {
		return V(val)
	}
}
