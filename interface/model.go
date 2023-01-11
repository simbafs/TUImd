package interfaces

import tea "github.com/charmbracelet/bubbletea"

type Model interface {
	tea.Model
	Get(string) interface{}
	Set(string, interface{})
}
