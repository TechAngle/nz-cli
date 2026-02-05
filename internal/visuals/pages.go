package visuals

import (
	"log"

	"github.com/gdamore/tcell/v2"
)

// List of pages name variables
const (
	MainPage  = "main"
	LoginPage = "login"
	ModalPage = "modals"
	NewsPage  = "news"
)

// renders pages and sets one as root
func (c *CLI) renderPages() {
	if c.pages == nil {
		log.Panicln("c.pages are nil!")
	}

	// starting updaters
	c.startClockUpdater()
	c.startAccountStateUpdater()
	c.startNotificationsUpdater()

	// pages definitions
	c.pages.
		AddAndSwitchToPage(MainPage, c.mainPage(), true).
		AddPage(LoginPage, c.loginPage(), true, false).
		AddPage(ModalPage, c.modalPage(), true, false).
		AddPage(NewsPage, c.newsPage(), true, false)

	// keys handler
	c.app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		// login page
		case tcell.KeyCtrlL:
			c.pages.SwitchToPage(LoginPage)

		// main page
		case tcell.KeyCtrlH:
			c.pages.SwitchToPage(MainPage)
		}

		return event
	})

	c.app.SetRoot(c.pages, true).EnableMouse(true)
}
