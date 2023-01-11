package keymap

import "github.com/charmbracelet/bubbles/key"

var Matches = key.Matches

var Quit = key.NewBinding(
	key.WithKeys("ctrl+c"),
	key.WithHelp("C-c", "quit"),
)

var CommandPrefix = key.NewBinding(key.WithKeys(":"))

var EnterCommand = key.NewBinding(key.WithKeys("enter"))

// need more detail
var BeginInsertMode = key.NewBinding(key.WithKeys("a", "i"))

var BeginNormalMode = key.NewBinding(key.WithKeys("esc"))
