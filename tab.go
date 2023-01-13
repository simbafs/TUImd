package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/simbafs/TUImd/util"
)

type Tab struct {
	Text string
}

func (m Tab) Init() tea.Cmd { return nil }
func (m Tab) Update(msg tea.Msg) (Tab, tea.Cmd) {
	// var cmd tea.Cmd
	// var cmds []tea.Cmd = make([]tea.Cmd, 4)

	switch msg := msg.(type) {
	case util.FileMsg:
		m.Text = msg.Filename
	}

	return m, nil
	// return m, tea.Batch(cmds...)
}

// func (m Tab) View() string {
// 	if mode == InsertMode {
// 		return "InsertMode"
// 	} else {
// 		return "NormalMode"
// 	}
// }

func (m Tab) View() string { return m.Text }

func NewTab() Tab {
	return Tab{
		Text: "",
	}
}
