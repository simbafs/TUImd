package util

import (
	"io/fs"
	"io/ioutil"
	"os/exec"
	Msg "tuimd/msg"

	tea "github.com/charmbracelet/bubbletea"
)

func RenderMD(md string) tea.Cmd {
	return func() tea.Msg {
		err := ioutil.WriteFile("/tmp/tuimd", []byte(md), fs.ModePerm)
		if err != nil {
			return Msg.ShowMsg("fail to save to temp file")
		}

		cmd := exec.Command("glow", "/tmp/tuimd")
		rendered, err := cmd.CombinedOutput()
		if err != nil {
			return Msg.ShowMsg("fail to render md")
		}

		return Msg.Rendered(string(rendered))
	}
}
