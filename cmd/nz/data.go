package main

import (
	"fmt"
	"nz-cli/internal/client"
	"nz-cli/internal/commons"
	"nz-cli/internal/visuals"
)

// replace shortcuts with real dates
func parseDate(date *string) {
	// parsing start date
	// NOTE: I think about changing this to map[string]string
	switch *date {
	case Today:
		*date = commons.TodayDate()
	case WeekStart:
		*date = commons.WeekStart()
	case WeekEnd:
		*date = commons.WeekEnd()
	case StartOfYear:
		*date = commons.StartOfSchoolYear()
	case EndOfYear:
		*date = commons.EndOfSchoolYear()
	}
}

// validate start and end dates
func validateDates() {
	if *startDate == "" {
		fail(visuals.ErrorStyle.Render("Start Date is invalid!"))
	}
	if *endDate == "" {
		fail(visuals.ErrorStyle.Render("End Date is invalid!"))
	}
}

// get only data from client, without TUI
func processDataFlags() {
	/*
		I tried to separate different flags (like additional and main) with different styles here.
		But I'll need to create more 'distinctive' color palette (or find another one).
	*/

	// checking for too many client flags
	// because if we try to process few ones then nz can reject our requests and fuck us with Rate Limit (at least if they have one on their mobile api)
	if !clientFlagsValid(*diary, *grades, *perfomance) {
		fail(visuals.ErrorStyle.Render("Invalid client flags (select only one)"))
	}

	// checking dates one more time ;)
	// just to be sure if someone stupid would set empty dates
	validateDates()

	// Initializating client
	client, err := client.NewClient()
	if err != nil {
		fail(visuals.ErrorStyle.Render("Failed to initialize client:"), err)
	}

	if *login {
		err := client.Login(*username, *password)
		if err != nil {
			fail("failed to login:", err)
		}

		return // hehe, i dont wanna go further, WRITE YOUR NEXT COMMAND AFTER AUTH :3
	}

	// restoring session
	err = client.RestoreSession()
	if err != nil {
		fail(visuals.ErrorStyle.Render("Failed to restore session:"), err)
	}

	// parsing and changing shortcuts
	parseDate(startDate)
	parseDate(endDate)

	switch true {
	// -diary flag
	case *diary:
		err = client.Diary(*startDate, *endDate)

	// -grades flag
	case *grades: // TODO: Test it when i'll get any grades
		if *subjectId == INVALID_ID {
			err = fmt.Errorf("invalid subject id: %v", err)
			break
		}

		err = client.Grades(*startDate, *endDate, *subjectId)

	// -perfomance flag
	case *perfomance:
		err = client.Perfomance(*startDate, *endDate)
	}

	// if any error occurred we'll just close program with status code 1, why not.
	if err != nil {
		fail(err.Error())
	}
}
