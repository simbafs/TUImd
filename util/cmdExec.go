package util

import (
	"strings"
	Msg "tuimd/msg"

	tea "github.com/charmbracelet/bubbletea"
)

func CmdExec(cmd string) tea.Cmd {
	// file, _ := os.OpenFile("debug.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// file.WriteString(cmd + "\n")

	if len(cmd) == 0 {
		return nil
	}
	args := strings.Split(cmd[1:], " ")

	if len(args) < 1 {
		return nil
	}

	switch args[0] {
	case "e", "edit":
		if len(args) < 2 {
			return nil
		}
		return ReadFile(args[1])
	case "q", "quit":
		return tea.Quit
	case "w", "write":
		return func() tea.Msg {
			return Msg.ShowMsg("The write command has not yet been implemented")
		}
		// if len(args) >= 2 {
		// 	filename = args[1]
		// }
		//
		// return SaveFile()
	}

	return nil
}
