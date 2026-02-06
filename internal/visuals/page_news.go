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
func (c *CLI) newsPage() *tview.Flex {
	flex := tview.NewFlex()

	flex.AddItem(c.newsState.newsList, 0, 1, true).
		SetBorder(true).
		SetTitle("News").
		SetTitleAlign(tview.AlignBottom)

	return flex
}
