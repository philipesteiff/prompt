package input

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	Confirm key.Binding
	Quit    key.Binding
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Confirm, k.Quit}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Confirm, k.Quit},
	}
}

func keys() KeyMap {
	keys := KeyMap{
		Confirm: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "confirm"),
		),
		Quit: key.NewBinding(
			key.WithKeys("esc", "ctrl+c"),
			key.WithHelp("esc", "quit"),
		),
	}

	return keys
}

func WithHelp(show bool) Option {
	return func(m *Model) {
		m.showHelp = show
	}
}
