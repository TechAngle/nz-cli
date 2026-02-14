package utils

import (
	"fmt"
	"nz-cli/internal/api"
	"time"
)

/*
 NOTE: I moved functions here because in utils package they cause `import cycle` error :\
*/

// Shortcut
const (
	today       = "today"
	weekStart   = "week-start"
	weekEnd     = "week-end"
	startOfYear = "start-of-year"
	endOfYear   = "end-of-year"
)

// replace shortcuts with real dates
func ShortcutToDate(date string) string {
	// parsing start date
	// NOTE: I think about changing this to map[string]string
	switch date {
	case today:
		date = TodayDate()
	case weekStart:
		date = WeekStart()
	case weekEnd:
		date = WeekEnd()
	case startOfYear:
		date = StartOfSchoolYear()
	case endOfYear:
		date = EndOfSchoolYear()
	default:

	}

	return date
}

// Get todays date in correct format
func TodayDate() string {
	return time.Now().Format(api.DateFormat)
}

// Get tomorrow's date
func NextDay() string {
	return time.Now().AddDate(0, 0, 1).Format(api.DateFormat)
}

// Get yesterday's date
func PreviousDay() string {
	return time.Now().AddDate(0, 0, -1).Format(api.DateFormat)
}

// Get start of current week
// took this part from "https://github.com/icza/gox/blob/main/timex/timex.go" but adapted it for today's date
func WeekStart() string {
	t := time.Now()

	// It is rollback to Monday of current week
	if wd := t.Weekday(); wd == time.Sunday {
		t = t.AddDate(0, 0, -6)
	} else {
		/*
			NOTE: How it works: if it is not Sunday (who is iota in time package, so week starts from 0)
				then we should add its negative value, so there: Tuesday (2) = -2+1 = 1 - its Monday
		*/
		t = t.AddDate(0, 0, -int(wd)+1)
	}

	return t.Format(api.DateFormat)
}

// Get end of current week.
// Reversed logic of WeekStart
func WeekEnd() string {
	t := time.Now()

	if wd := t.Weekday(); wd == time.Monday {
		t = t.AddDate(0, 0, 6) // adding 6 more day for Sunday
	} else {
		/*
			 NOTE: How it works: in Ukrainian time Monday is the first day, but in time package it is 1,
				so that we can use simple logic like this: if today is Wednesday (3) = 7 - 3 = 4
				=> 4 is how many days we should add.
		*/
		t = t.AddDate(0, 0, 7-int(wd))
	}

	return t.Format(api.DateFormat)
}

// Get the start of school year
func StartOfSchoolYear() string {
	/*
		We need to check here those things:
		- IF it is start of the year (e.g. 2026-01-01) then we can just subdivide 1 year and set 09-01 as date
		- IF it is more than the half of summer, then we automatically set date to current year but 09-01,
			else we just use start of THIS year (afaik, nz cleans your previous grades,
			so we have no need to watch them, but they can be available through arguments like -start-date and -end-date)
	*/

	t := time.Now()

	// current year
	year := t.Year()

	// checking for more than half of summer, so we can return just start of this year
	if t.Month() >= time.July {
		return fmt.Sprintf("%d-09-01", year)
	} else { // it is a start of current year and its continuing, so we subing what we got and returning it
		year -= 1
		return fmt.Sprintf("%d-09-01", year)
	}
}

// Get the start of school year
func EndOfSchoolYear() string {
	/* This thing has a bit other logic. We need to do same thing, BUT we should add to current year and reverse other things */
	t := time.Now()

	// current year
	year := t.Year()

	// checking for less than half of summer, so we can return just end of this year
	if t.Month() <= time.July {
		return fmt.Sprintf("%d-05-31", year)
	} else { // it is a more than half of summer of current year and its continuing, so we add to current year and returning it
		year += 1
		return fmt.Sprintf("%d-05-31", year)
	}
}
