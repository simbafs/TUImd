package component

import tea "github.com/charmbracelet/bubbletea"

type Tab string

func (s Tab) Init() tea.Cmd                     { return nil }
func (s Tab) Update(msg tea.Msg) (Tab, tea.Cmd) { return s, nil }
func (s Tab) View() string                      { return string(s) }
