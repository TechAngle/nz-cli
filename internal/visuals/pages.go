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
	c.StartClockUpdater()
	c.StartAccountStateUpdater()
	c.StartNotificationsUpdater()

	// pages definitions
	c.pages.
		AddAndSwitchToPage(MainPage, c.mainPage(), true).
		AddPage(LoginPage, c.loginPage(), true, false).
		AddPage(ModalPage, c.modalPage(), true, false).
		AddPage(NewsPage, c.newsPage(), true, false)

	// keys handler
	c.app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		// login page
		case 'l':
			c.pages.SwitchToPage(LoginPage)

		// main page
		case 'm':
			c.pages.SwitchToPage(MainPage)

		// news page
		case 'n':
			c.pages.SwitchToPage(NewsPage)

		// news page update
		case 'r':
			visiblePages := c.pages.GetPageNames(true)
			if len(visiblePages) != 0 {
				if visiblePages[0] == NewsPage {
					// TODO: Update news
				}
			}

		// quit hotkey
		case 'q':
			c.SaveSession()
			c.app.Stop()
		}

		return event
	})

	c.app.SetRoot(c.pages, true).
		EnableMouse(true)
}
