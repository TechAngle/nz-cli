package visuals

import (
	"fmt"
	"maps"
	"strings"
)

var (
	// commands list
	commands = map[string]string{
		"help": "Show this menu",
		// TODO: add others
		"quit": "Quit app",
	}
)

func help() string {
	var b strings.Builder
	b.WriteString(MainStyleBold.Render("================== HELP ==================") + "\n")

	// generate commands list
	for command := range maps.Keys(commands) {
		fmt.Fprintf(
			&b,
			"- %s - %s\n",
			SecondStyleReverse.Render(command),
			MainStyle.Render(commands[command]),
		)
	}

	return b.String()
}
