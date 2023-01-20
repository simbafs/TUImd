package util

import (
	"io/ioutil"

	tea "github.com/charmbracelet/bubbletea"

	Msg "tuimd/msg"
)

func ReadFile(filename string) tea.Cmd {
	return tea.Batch(func() tea.Msg {
		b, err := ioutil.ReadFile(filename)
		if err != nil {
			return nil
		}

		return Msg.Body(string(b))
	}, Msg.NewCmd[Msg.Filename](filename))
}
