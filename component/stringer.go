package component

import tea "github.com/charmbracelet/bubbletea"

type Stringer string

func (s Stringer) Init() tea.Cmd                           { return nil }
func (s Stringer) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return s, nil }
func (s Stringer) View() string                            { return string(s) }
