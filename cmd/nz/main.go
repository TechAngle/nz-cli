package main

import (
	"flag"
	"fmt"
	"log"
	"nz-cli/internal/cli"
	"nz-cli/internal/commons"
	"nz-cli/internal/visuals"
)

// mode flags
var (
	// whether give only data
	data = flag.Bool(
		"data",
		true,
		visuals.SecondStyle.Render(commons.DataFlagUsage))

	// whether use TUI version (experimental)
	tui = flag.Bool(
		"tui",
		false,
		visuals.SecondStyle.Render(commons.TUIFlagUsage))
)

func main() {
	flag.Parse()

	switch {
	case *tui:
		// setting -data flag to false
		*data = false
		tui, err := visuals.NewTUI()
		if err != nil {
			log.Panicln(err)
		}

		tui.Run()
		fmt.Println("TUI feature in development at the moment")
	case *data:
		c, err := cli.NewClient()
		if err != nil {
			log.Panicln("failed to create client:", err)
		}

		err = c.ProcessCLIFlags()
		if err != nil {
			log.Panicln("failed to process CLI flags:", err)
		}
	}
}
