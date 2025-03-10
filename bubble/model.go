package bubble

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	choices  []string
	queries  []string
	cursor   int
	selected map[int]struct{}
	spinner  spinner.Model
}

func InitModel() model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	return model{
		choices:  []string{"All Tables", "All Users", "All Payments"},
		queries:  []string{"tables", "users", "payments"},
		selected: make(map[int]struct{}),
		spinner:  s,
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(m.spinner.Tick)
}
