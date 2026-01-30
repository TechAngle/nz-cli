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
	login    = flag.Bool("login", false, visuals.ThirdStyleBold.Render("Login to the system. Should be set with --username and --password flags"))
	username = flag.String("username", "", visuals.ThirdStyleBold.Render("Username. Required if -login argument is set"))
	password = flag.String("password", "", visuals.ThirdStyleBold.Render("Password. Required if -login argument is set"))

	// additional flags
	startDate = flag.String("start-date", commons.TodayDate(), visuals.FourthStyleBold.Render(commons.StartDateArgUsage))
	endDate   = flag.String("end-date", commons.TodayDate(), visuals.FourthStyleBold.Render(commons.EndDateArgUsage))
	subjectId = flag.Int("subject-id", INVALID_ID, visuals.FourthStyleBold.Render(commons.SubjectIdArgUsage))

	// client flags
	diary      = flag.Bool("diary", false, visuals.MainStyleBold.Render("Show diary"))
	grades     = flag.Bool("grades", false, visuals.MainStyleBold.Render("Show grades"))
	perfomance = flag.Bool("perf", false, visuals.MainStyleBold.Render("Show perfomance"))

	// mode flags
	data = flag.Bool("data", true, visuals.SecondStyle.Render("Get only data from API and format it (API mode)"))
	tui  = flag.Bool("tui", false, visuals.SecondStyle.Render("Run TUI (Terminal User Interface) (ONLY IN PROGRESS)"))
)
