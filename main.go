package main

import (
	"didactic-sniffle-go/bubble"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	fmt.Println("Go Sniffle!")

	p := tea.NewProgram(bubble.InitModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
