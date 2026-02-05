package visuals

import "github.com/rivo/tview"

type modalState struct {
	message *tview.TextView
}

// initialize empty modal state with widgets
func initModalState() *modalState {
	return &modalState{
		message: tview.NewTextView(),
	}
}

// get modal with message
func (c *CLI) modalPage() *tview.Modal {
	modal := tview.NewModal()
	modal.
		SetText(c.modalState.message.GetText(true)).
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			c.pages.SwitchToPage("main")
		})

	return modal
}
