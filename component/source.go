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

type Source struct {
	Mode     Mode
	Textarea textarea.Model
}

func (m Source) Init() tea.Cmd { return nil }
func (m Source) Update(msg tea.Msg) (Source, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd = make([]tea.Cmd, 4)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Textarea.SetHeight(msg.Height - 4)
	case util.FileMsg:
		m.Textarea.SetValue(msg.Content)
		m.Textarea.Focus()
	case tea.KeyMsg:
		// mode change should in main model
		if m.Mode == InsertMode {
			m.Textarea, cmd = m.Textarea.Update(msg)
			cmds = append(cmds, cmd)
			m.Textarea.KeyMap = textarea.DefaultKeyMap
		} else if m.Mode == NormalMode {
			// m.Textarea.KeyMap.CharacterForward.SetKeys("right", "ctrl+f", "l")
			// m.Textarea.KeyMap.CharacterBackward.SetKeys("left", "ctrl+b", "h")
			// m.Textarea.KeyMap.LineNext.SetKeys("down", "ctrl+n", "j")
			// m.Textarea.KeyMap.LinePrevious.SetKeys("up", "ctrl+p", "k")
		}

	case SaveFileMsg:
		cmds = append(cmds, func() tea.Msg {
			filename := string(msg)
			file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				return UpdateMsgMsg("Failed to open " + filename)
			}
			_, err = file.WriteString(m.Textarea.Value())
			if err != nil {
				return UpdateMsgMsg("Failed to write to " + filename)
			}
			return UpdateMsgMsg(fmt.Sprintf("\"%s\" %dL written", filename, m.Textarea.LineCount()))
		})
	}

	return m, tea.Batch(cmds...)
}
func (m Source) View() string { return m.Textarea.View() }

func NewSouce(text string) Source {
	input := textarea.New()
	input.CharLimit = 0
	input.KeyMap.CharacterForward.SetKeys("right", "ctrl+f", "l")
	input.KeyMap.CharacterBackward.SetKeys("left", "ctrl+b", "h")
	input.KeyMap.LineNext.SetKeys("down", "ctrl+n", "j")
	input.KeyMap.LinePrevious.SetKeys("up", "ctrl+p", "k")
	return Source{
		Mode:     NormalMode,
		Textarea: input,
	}
}
