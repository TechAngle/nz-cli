package cli

import (
	"flag"
	"fmt"
	"nz-cli/internal/api"
	"nz-cli/internal/visuals"
)

func (c *CLIClient) processClientFlag() (err error) {
	switch {
	// -diary flag
	case *diary:
		err = c.Diary(*startDate, *endDate)

	// -grades flag
	case *grades:
		if *subjectId == INVALID_ID {
			err = fmt.Errorf("invalid subject id: %v", err)
			break
		}

		err = c.Grades(*startDate, *endDate, *subjectId)

	// -perf flag
	case *performance:
		err = c.Performance(*startDate, *endDate)

	default:
		err = nil
	}

	return
}

// get only data from client, without TUI
func (c *CLIClient) ProcessCLIFlags() (err error) {
	flag.Parse()
	/*
		I tried to separate different flags (like additional and main) with different styles here.
	*/

	// checking for too many client flags
	// because if we try to process few ones then nz can reject our requests and fuck us with Rate Limit (at least if they have one on their mobile api)
	if !clientFlagsValid(*diary, *grades, *performance) {
		fail(visuals.ErrorStyle.Render("Invalid client flags (select only one)"))
	}

	// checking dates one more time ;)
	// just to be sure if someone stupid would set empty dates
	validateDates()

	// check login
	if *login {
		err := c.Login(*username, *password)
		if err != nil {
			return fmt.Errorf("failed to login: %v", err)
		}

		return nil // hehe, i dont wanna go further, WRITE YOUR NEXT COMMAND AFTER AUTH :3
	}

	// restoring session
	err = c.RestoreSession()
	if err != nil {
		fail(visuals.ErrorStyle.Render("Failed to restore session:"), err)
	}

	processDateFlags()

	if err = c.processClientFlag(); err != nil {
		return err
	}

	return nil
}

// initializate new api client wrapper
func NewClient() (*CLIClient, error) {
	client, err := api.NewAPIClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}

	return &CLIClient{
		client: client,
	}, nil
}
