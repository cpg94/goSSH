package main

import (
	"fmt"
	"os"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/cpg94/goSSH/jsonutils"
)

type model struct {
	cursor   int
	choices  jsonutils.Sessions
	selected jsonutils.Session
}

func main() {
	sessions := jsonutils.Read()

	p := tea.NewProgram(model{cursor: 0, choices: *sessions, selected: sessions.Sessions[0]})
	if _, err := p.Run(); err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {
		case "ctrl+c", "q":
			cmd := exec.Command("ssh", "x@x.com", "-v")
			return m, tea.ExecProcess(cmd, func(err error) tea.Msg {

				fmt.Printf("%v", err)
				return nil
			})

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.choices.Sessions)-1 {
				m.cursor++
			}

		case "enter", " ":
			selection := m.choices.Sessions[m.cursor]
			m.selected = selection
		}
	}

	return m, nil
}

func (m model) View() string {
	s := "Choose SSH Session\n\n"

	for i, choice := range m.choices.Sessions {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if m.selected.Id == choice.Id {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice.Name)
	}

	s += "\nPress q to quit.\n"

	return s
}
