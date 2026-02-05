package visuals

import (
	"fmt"
	"maps"
	"strings"

	"github.com/rivo/tview"
)

var (
	keysHelpMap = map[string]string{
		"Ctrl + L": "Login page",
		"Ctrl + H": "Main page",
		"Ctrl + N": "News page",
	}
)

// get help message element
func (c *CLI) helpMessage() *tview.TextView {
	var helpStr strings.Builder

	helpStr.WriteString("========= HELP =========\n")

	// parsing keys and values
	for key, page := range maps.All(keysHelpMap) {
		fmt.Fprintf(
			&helpStr,
			"- %s - %s\n",
			fmt.Sprintf("[%s::b]%s[%s::b]", thirdCode, key, thirdCode),
			fmt.Sprintf("[%s]%s[%s]", fourthCode, page, thirdCode),
		)
	}

	helpStr.WriteString("\tOthers soon....")

	return tview.NewTextView().
		SetText(helpStr.String()).
		SetTextAlign(tview.AlignBottom).
		SetDynamicColors(true)
}
