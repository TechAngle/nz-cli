package cli

import (
	"nz-cli/internal/utils"
	"nz-cli/internal/visuals"
)

// validate start and end dates
func validateDates() {
	if *startDate == "" {
		fail(visuals.ErrorStyle.Render("Start Date is invalid!"))
	}
	if *endDate == "" {
		fail(visuals.ErrorStyle.Render("End Date is invalid!"))
	}
}

// replacing range if one of arguments set.
// i dont think if we put two string to stack it will eat so much memory
func processDateFlags() {
	if *dateFlag != "" {
		*startDate, *endDate = *dateFlag, *dateFlag
	} else if *tomorrow {
		*startDate, *endDate = utils.NextDay(), utils.NextDay()
	} else if *yesterday {
		*startDate, *endDate = utils.PreviousDay(), utils.PreviousDay()
	} else {
		// parsing and replacing shortcuts to their dates
		*startDate = utils.ShortcutToDate(*startDate)
		*endDate = utils.ShortcutToDate(*endDate)
	}
}

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
