package main

import (
	"flag"
	"fmt"
	"nz-cli/internal/visuals"
	"os"
)

const (
	/*
		fr, idk which i should use here, but i guess it should be negative.
		If it positive, then such ID will be invalid, so let it be -100 at least.

		P.S. I'll hate nz devs (i am already are, whatever) if they would add negatives as IDs.
	*/
	INVALID_ID = -100
)

func clientFlagsValid(flags ...bool) bool {
	// checking how many flags are true
	amount := 0

	for _, f := range flags {
		if f {
			amount++
		}
	}

	// if more than one flag are true we failing attempt
	return amount == 1
}

// Print error and exit with code 1
func fail(message string, v ...any) {
	fmt.Println(visuals.ErrorStyle.Render(message), v)
	os.Exit(1)
}

func main() {
	flag.Parse()

	switch {
	case *tui:
		// setting -data flag to false
		*data = false
		// TODO: Run TUI
		fmt.Println("TUI feature in development at the moment")
	case *data:
		processDataFlags()
	}
}
