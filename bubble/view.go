package bubble

import (
	"fmt"
)

func QueryTables() string {
	return "Finding all tables!"
}

func QueryUsers() string {
	return "Finding all users!"
}

func QueryPayments() string {
	return "Finding all payments!"
}

func (m model) View() string {
	// The header
	s := "Show:\n\n"

	// Iterate over our choices
	for i, choice := range m.choices {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?
		checked := " " // not selected
		if _, ok := m.selected[i]; ok {
			checked = "x" // selected!
		}

		getting := ""
		if _, ok := m.selected[i]; ok {
			switch m.queries[i] {
			case "tables":
				result := QueryTables()
				getting = ("-> " + m.spinner.View() + result)
			case "users":
				result := QueryUsers()
				getting = ("-> " + m.spinner.View() + result)
			case "payments":
				result := QueryPayments()
				getting = ("-> " + m.spinner.View() + result)
			}
		} else {
			getting = ""
		}

		s += fmt.Sprintf("%s [%s] %s %s\n", cursor, checked, choice, getting)
	}

	// The footer
	s += "\nPress q to quit.\n"

	return s
}
