package bubble

import (
	"fmt"
)

type respMsg string

func (m model) View() string {
	s := "Show:\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	if m.loading {
		s += fmt.Sprintf("\n-> %s Querying database...\n", m.spinner.View())
	} else if m.request != "" {
		s += fmt.Sprintf("\n%s\n", m.request)
	}

	s += "\nPress q to quit.\n"
	return s
}
