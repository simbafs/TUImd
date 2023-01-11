package component

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/simbafs/TUImd/util"
)

type UpdateContentMsg string
type SaveFileMsg string

type mode int

const (
	insertMode mode = iota
	normalMode
)

type Source struct {
	mode     mode
	textarea textarea.Model
}

func (m Source) Init() tea.Cmd { return nil }
func (m Source) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd = make([]tea.Cmd, 4)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.textarea.SetHeight(msg.Height - 4)
	case util.FileMsg:
		m.textarea.SetValue(msg.Content)
		m.textarea.Focus()
	case tea.KeyMsg:
		// mode change should in main model
		if msg.Type == tea.KeyEsc {
			m.mode = normalMode
		}

		if msg.String() == "a" || msg.String() == "i" {
			m.mode = insertMode
		}

		if m.mode == insertMode {
			m.textarea, cmd = m.textarea.Update(msg)
			cmds = append(cmds, cmd)
		}

		// cmds = append(cmds, func() tea.Msg {
		//     return UpdateContentMsg(m.textarea.Value())
		// })
	case SaveFileMsg:
		cmds = append(cmds, func() tea.Msg {
			filename := string(msg)
			file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				return UpdateMsgMsg("Failed to open " + filename)
			}
			_, err = file.WriteString(m.textarea.Value())
			if err != nil {
				return UpdateMsgMsg("Failed to write to " + filename)
			}
			return UpdateMsgMsg(fmt.Sprintf("\"%s\" %dL written", filename, m.textarea.LineCount()))
		})
	}

	return m, tea.Batch(cmds...)
}
func (m Source) View() string { return m.textarea.View() }

func NewSouce(text string) Source {
	return Source{
		mode:     normalMode,
		textarea: textarea.New(),
	}
}
