package visuals

import (
	"github.com/rivo/tview"
)

// main page state
type mainState struct {
	// Current logged account.
	loggedAccountLabel *tview.TextView

	// Current clock time
	clockLabel *tview.TextView

	// Quantity of notifications
	unreadNewsQty *tview.TextView

	// Notifications list
	// NOTE: Better to use only 5 latest news as it was planned initially in prototype
	shortNewsList *tview.List
}

// initializate main state with widgets
func initMainState() *mainState {
	return &mainState{
		loggedAccountLabel: tview.NewTextView().
			SetTextAlign(tview.AlignRight).
			SetTextColor(fourthXTermCode),
		clockLabel: tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetTextColor(mainXTermCode),
		unreadNewsQty: tview.NewTextView().
			SetTextAlign(tview.AlignLeft).
			SetDynamicColors(true),
		shortNewsList: tview.NewList().
			SetMainTextColor(fourthXTermCode).
			SetSecondaryTextColor(secondXTermCode).
			SetHighlightFullLine(true),
	}
}

// main page primitive
func (c *CLI) mainPage() *tview.Flex {
	// grid := tview.NewGrid()
	flex := tview.NewFlex()

	// adding components
	flex.
		// logged account information
		AddItem(c.mainState.loggedAccountLabel, 1, 1, false).
		// time clock
		AddItem(c.mainState.clockLabel, 1, 1, false).

		// notifications quantity
		AddItem(c.mainState.unreadNewsQty, 1, 1, false).

		// Latest news block
		AddItem(tview.NewTextView().SetText("Latest news:"), 1, 1, false).
		// latest notifications
		AddItem(c.mainState.shortNewsList, 0, 1, true).
		// help message
		AddItem(c.helpMessage(), 0, 1, false)

	// flex settings
	flex.SetDirection(tview.FlexRow). // everythin should be in the row
						SetBorder(true). // borders
						SetTitle("Main")

	return flex
}
