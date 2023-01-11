package component

import tea "github.com/charmbracelet/bubbletea"

type Markdown string

func (s Markdown) Init() tea.Cmd                          { return nil }
func (s Markdown) Update(msg tea.Msg) (Markdown, tea.Cmd) { return s, nil }
func (s Markdown) View() string                           { return string(s) }
