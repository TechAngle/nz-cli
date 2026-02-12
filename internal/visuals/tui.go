package visuals

import (
	"fmt"
	"log"
	"nz-cli/internal/api"

	"github.com/rivo/tview"
)

// TODO: Finish CLI interface

// CLI state manager structure
type TUI struct {
	app        *tview.Application
	pages      *tview.Pages
	userData   *userData
	mainState  *mainState
	modalState *modalState
	newsState  *newsState

	client *api.NZAPIClient
}

// Restore account session
func (c *TUI) RestoreSession() {
	err := c.client.LoadAccount()
	if err != nil {
		log.Println("failed to restore session:", err)
	}
}

// Save account session
func (c *TUI) SaveSession() {
	err := c.client.SaveSession()
	if err != nil {
		log.Println("Failed to save session:", err)
	}
}

// Run tui
func (c *TUI) Run() {
	// log.Println("Starting...")
	// Loading it here, because in NewCLI it will cause troubles if we don't need account before running
	c.RestoreSession()
	c.renderPages()

	// running program
	if err := c.app.Run(); err != nil {
		c.SaveSession()
		log.Panicln(err)
	}
}

func NewCLI() (cli *TUI, err error) {
	client, err := api.NewApiClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create api client: %v", err)
	}

	app := tview.NewApplication()
	pages := tview.NewPages()

	return &TUI{
		// states
		userData:   &userData{},
		mainState:  initMainState(),
		modalState: initModalState(),
		newsState:  initNewsState(),

		// others
		app:    app,
		pages:  pages,
		client: client,
	}, nil
}
