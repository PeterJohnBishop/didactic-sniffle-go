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
	request  string
	loading  bool
}

func InitModel() model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	return model{
		choices:  []string{"All Tables", "All Users", "All Messages"},
		queries:  []string{"tables", "users", "messages"},
		selected: make(map[int]struct{}),
		spinner:  s,
		request:  "",
		loading:  false,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}
