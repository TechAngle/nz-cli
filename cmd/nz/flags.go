package main

import (
	"flag"
	"nz-cli/internal/commons"
	"nz-cli/internal/visuals"
)

// Shortcut
const (
	Today       = "today"
	WeekStart   = "week-start"
	WeekEnd     = "week-end"
	StartOfYear = "start-of-year"
	EndOfYear   = "end-of-year"
)

// flags list
var (
	// login flags

	// whether user wants to login (requires set -username and -password)
	login = flag.Bool(
		"login",
		false,
		visuals.ThirdStyleBold.Render(commons.LoginFlagUsage),
	)

	// username for login
	username = flag.String(
		"username",
		"",
		visuals.ThirdStyleBold.Render(commons.UsernameFlagUsage),
	)

	// password for login
	password = flag.String(
		"password",
		"",
		visuals.ThirdStyleBold.Render(commons.PasswordFlagUsage),
	)

	// additional flags

	// use some specific day
	dateFlag = flag.String(
		"date",
		"",
		visuals.FourthStyle.Render(),
	)

	// range start
	startDate = flag.String(
		"start-date",
		commons.TodayDate(),
		visuals.FourthStyleBold.Render(commons.StartDateFlagUsage),
	)

	// range end
	endDate = flag.String(
		"end-date",
		commons.TodayDate(),
		visuals.FourthStyleBold.Render(commons.EndDateFlagUsage),
	)

	// specific subject id
	subjectId = flag.Int(
		"subject-id",
		INVALID_ID,
		visuals.FourthStyleBold.Render(commons.SubjectIdFlagUsage),
	)

	// date-for-use flags

	// whether use tomorrow's date
	tomorrow = flag.Bool(
		"tomorrow",
		false,
		visuals.SecondStyle.Render(commons.TomorrowFlagUsage),
	)

	// whether use yesterday's date
	yesterday = flag.Bool(
		"yesterday",
		false,
		visuals.SecondStyle.Render(commons.TomorrowFlagUsage),
	)

	// client flags (requires start-date and end-date)

	// show diary
	diary = flag.Bool(
		"diary",
		false,
		visuals.MainStyleBold.Render(commons.DiaryFlagUsage),
	)

	// show grades of specific subject using its ID
	grades = flag.Bool(
		"grades",
		false,
		visuals.MainStyleBold.Render(commons.GradesFlagUsage),
	)

	// show perfomance
	perfomance = flag.Bool(
		"perf",
		false,
		visuals.MainStyleBold.Render(commons.PerfomanceFlagUsage),
	)

	// mode flags

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
