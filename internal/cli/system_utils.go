package cli

import (
	"fmt"
	"nz-cli/internal/visuals"
	"os"

	"golang.org/x/term"
)

// Print error and exit with code 1
func fail(message string, v ...any) {
	fmt.Println(visuals.ErrorStyle.Render(message), v)
	os.Exit(1)
}

// Get terminal size
func terminalSize() (width, height int, err error) {
	fd := int(os.Stdout.Fd())
	width, height, err = term.GetSize(fd)
	if err != nil {
		return -1, -1, fmt.Errorf("cannot get terminal size: %v", err)
	}

	return
}
