package visuals

import "github.com/rivo/tview"

// login page primitive
func (c *CLI) loginPage() *tview.Form {
	f := tview.NewForm()
	f.
		AddInputField("Username", "", 16, nil, func(text string) {
			c.userData.username = text
		}).
		AddInputField("Password", "", 16, nil, func(text string) {
			c.userData.username = text
		}).
		AddButton("Log in", func() {
			err := c.Login()
			if err != nil {
				c.modalState.message.SetText(err.Error())
				c.pages.SwitchToPage(ModalPage)
				return
			}

			c.pages.SwitchToPage("main")
		}).
		SetBorder(true).
		SetBorderPadding(5, 5, 5, 5).
		SetTitleAlign(tview.AlignCenter).
		SetTitle("Login")

	return f
}
