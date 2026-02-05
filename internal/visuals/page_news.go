package visuals

import "github.com/rivo/tview"

type newsState struct {
	// full list of news
	newsList *tview.List
}

func initNewsState() *newsState {
	return &newsState{
		newsList: tview.NewList(),
	}
}

func (*CLI) newsPage() *tview.Flex {
	flex := tview.NewFlex()

	return flex
}
