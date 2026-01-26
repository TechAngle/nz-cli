package visuals

import (
	"fmt"
	"log"
	"nz-cli/internal/api"

	tea "github.com/charmbracelet/bubbletea"
)

// TODO: Finish CLI interface

var (
	client *api.NZAPIClient
)

type CLI struct {
	quit bool
}

// save client session
func (c *CLI) saveSession() {
	err := client.SaveSession()
	if err != nil {
		log.Println("Failed to save session:", err)
	}
}

// Run cli ui
func (c *CLI) Run() {
	log.Println("Starting...")
	p := tea.NewProgram(initialModel(), tea.WithMouseAllMotion())
	if _, err := p.Run(); err != nil {
		c.saveSession()
		panic(err)
	}
}

func NewCLI() (cli *CLI, err error) {
	client, err = api.NewApiClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create api client: %v", err)
	}

	return &CLI{
		quit: false,
	}, nil
}
