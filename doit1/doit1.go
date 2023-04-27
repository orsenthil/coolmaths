package main

import tea "github.com/charmbracelet/bubbletea"

func main() {
	m := menu{
		options: []menuItem{
			{
				text: "new check-in",
				onPress: func() tea.Msg {
					return toggleCasingMsg{}
				},
			},
			{
				text: "view check-ins",
				onPress: func() tea.Msg {
					return toggleCasingMsg{}
				},
			},
		},
	}

	p := tea.NewProgram(m)

	err := p.Start()
	if err != nil {
		panic(err)
	}
}
