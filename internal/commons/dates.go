package commons

import (
	"fmt"
	"time"
)

var (
	TODAY_DATE           = TodayDate()
	START_OF_SCHOOL_YEAR = StartOfSchoolYear()
	END_OF_SCHOOL_YEAR   = EndOfSchoolYear()
)

/*
 NOTE: I moved functions here because in utils package them cause `import cycle` error :\
*/

// get todays date and format it
func TodayDate() string {
	return time.Now().Format(DateFormat)
}

// get the start of school year
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

// get the start of school year
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
