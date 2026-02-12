package visuals

import "github.com/rivo/tview"

type newsState struct {
	// full list of news
	newsList *tview.List
}

// initializate news state with defined structures
func initNewsState() *newsState {
	return &newsState{
		newsList: tview.NewList().SetHighlightFullLine(false),
	}
}

// get news page
func (c *TUI) newsPage() *tview.Flex {
	flex := tview.NewFlex()

	flex.
		AddItem(tview.NewTextView().
			SetText("Press R to update news").SetTextAlign(tview.AlignCenter), 2, 1, false).
		AddItem(c.newsState.newsList, 0, 1, true).
		SetDirection(tview.FlexRow).
		SetBorder(true).
		SetTitle("News").
		SetTitleAlign(tview.AlignBottom)

	return flex
}
