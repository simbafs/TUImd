package component

import (
	tea "github.com/charmbracelet/bubbletea"
)

type SourceMsg string

type Source struct {
	text string
}

func (s Source) Init() tea.Cmd { return nil }
func (s Source) Update(msg tea.Msg) (Source, tea.Cmd) {
	switch msg := msg.(type) {
	case SourceMsg:
		s.text = string(msg)
	}
	return s, nil
}
func (s Source) View() string { return s.text }

func (s Source) SetText(str string) {
	// ioutil.WriteFile("debug.log", []byte(str), 0644)
	s.text = str
}

func NewSouce(text string) Source {
	return Source{
		text: text,
	}
}
