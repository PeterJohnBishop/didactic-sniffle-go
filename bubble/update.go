package bubble

import (
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

// update receives a tea message and returns a new version of the model and a command.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

			// The "enter" key and the spacebar (a literal space) toggle
			// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
				m.request = ""
				m.loading = false
			} else {
				m.selected[m.cursor] = struct{}{}
				m.loading = true

				var queryCmd tea.Cmd
				switch m.queries[m.cursor] {
				case "tables":
					m.request = "-> " + m.spinner.View() + " Finding all tables!"
					queryCmd = queryTables()
				case "users":
					m.request = "-> " + m.spinner.View() + " Finding all users!"
					queryCmd = queryUsers()
				case "messages":
					m.request = "-> " + m.spinner.View() + " Finding all messages!"
					queryCmd = queryMessages()
				}

				return m, tea.Batch(
					queryCmd,
					m.spinner.Tick,
				)
			}
		}

	case spinner.TickMsg:
		m.spinner, cmd = m.spinner.Update(msg)
		if m.loading {
			cmds = append(cmds, cmd) // keep ticking
		}

	case respMsg:
		m.selected = map[int]struct{}{} // clear selection so it won't retrigger in view
		m.loading = false
		switch msg {
		case "Tables!":
			m.request = "-> Found all tables!"
		case "Users!":
			m.request = "-> Found all users!"
		case "Messages!":
			m.request = "-> Found all messages!"
		}
	}

	return m, tea.Batch(cmds...)

}

func queryTables() tea.Cmd {
	return func() tea.Msg {
		time.Sleep(5 * time.Second)
		return respMsg("Tables!")
	}
}

func queryUsers() tea.Cmd {
	return func() tea.Msg {
		time.Sleep(8 * time.Second)
		return respMsg("Users!")
	}
}

func queryMessages() tea.Cmd {
	return func() tea.Msg {
		time.Sleep(3 * time.Second)
		return respMsg("Messages!")
	}
}
