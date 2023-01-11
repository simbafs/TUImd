package util

import (
	"io/ioutil"

	tea "github.com/charmbracelet/bubbletea"
)

type FileMsg struct {
	Filename string
	Content  string
}

func ReadFile(filename string) tea.Cmd {
	return func() tea.Msg {
		b, err := ioutil.ReadFile(filename)
		if err != nil {
			return nil
		}

		return FileMsg{
			Filename: filename,
			Content:  string(b),
		}
	}
}
