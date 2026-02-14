package cli

import (
	"fmt"
	"nz-cli/internal/api"
	"nz-cli/internal/utils"
	"nz-cli/internal/visuals"
	"os"
	"strings"
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

// replacing range if one of arguments set
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
		utils.ShortcutToDate(*startDate)
		utils.ShortcutToDate(*endDate)
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

// Print error and exit with code 1
func fail(message string, v ...any) {
	fmt.Println(visuals.ErrorStyle.Render(message), v)
	os.Exit(1)
}

// Removes any " ", "." from subject name
func normalizeName(subjectName string) string {
	dotFree := strings.ReplaceAll(subjectName, ".", "")
	spaceFree := strings.TrimSpace(dotFree)

	return spaceFree
}

// Organise, join subjects and returns it
func normalizeSubjects(subjects []api.Subject) []api.Subject {
	subjectsList := []api.Subject{}

	previousSubjectName := ""

	for _, cSubject := range subjects {
		subjectName := normalizeName(cSubject.SubjectName)

		if subjectName == previousSubjectName {
			// using ptr to the previous one
			subject := &subjectsList[len(subjectsList)-1]

			// adding marks
			subject.Marks = append(subject.Marks, cSubject.Marks...)
			subject.SubjectID = strings.Join([]string{subject.SubjectID, cSubject.SubjectID}, ", ")

			continue
		}

		subjectsList = append(subjectsList, cSubject)
		previousSubjectName = subjectName
	}

	return subjectsList
}
