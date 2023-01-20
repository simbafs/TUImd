package util

import (
	"io/ioutil"
	Msg "tuimd/msg"

	tea "github.com/charmbracelet/bubbletea"
)

func SaveFile(filename, body string) tea.Cmd {
	return func() tea.Msg {
		err := ioutil.WriteFile(filename, []byte(body), 0644)
		if err != nil {
			return Msg.ShowMsg("fail to save " + filename)
		}

		return Msg.ShowMsg("\"" + filename + "\"" + "written")
	}
}
