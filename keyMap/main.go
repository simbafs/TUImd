package keymap

import "github.com/charmbracelet/bubbles/key"

var Quit = key.NewBinding(
	key.WithKeys("ctrl+c"),
	key.WithHelp("C-c", "quit"),
)

var CommandPrefix = key.NewBinding(key.WithKeys(":"))
