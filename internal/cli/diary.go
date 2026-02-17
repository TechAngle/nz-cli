package cli

import (
	"fmt"
	"nz-cli/internal/api"
	"nz-cli/internal/visuals"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

// API client wrapper
type CLIClient struct {
	// API client
	client *api.NZAPIClient
}

// Get diary
func (c *CLIClient) Diary(startDate string, endDate string) error {
	diary, err := c.client.Diary(api.DefaultPayload{
		StartDate: startDate,
		EndDate:   endDate,
		StudentID: c.client.Account().StudentID,
	})
	if err != nil {
		return fmt.Errorf("failed to get diary: %v", err)
	}

	// if no information just return
	if len(diary.Dates) == 0 {
		fmt.Println(visuals.ThirdStyleBold.Render("No information to show!"))
		return nil
	}

	displayDiaryTable(diary)

	return nil
}

// Returns dates list as headers and hometasks as rows
func diaryDates(dates *[]api.Date) ([]string, []string) {
	datesList := make([]string, 0, len(*dates))
	hometasksRow := make([]string, len(*dates))

	// going through dates
	for i, date := range *dates {
		datesList = append(datesList, date.Date)
		var day strings.Builder

		// checking calls
		for _, call := range date.Calls {
			// parsing subjects
			for _, subject := range call.Subjects {
				task := strings.TrimSpace(strings.Join(subject.Hometask, ";"))
				fmt.Fprintf(
					&day,
					"[%s] (Teacher: %s): \n\t%s\n\n",
					visuals.SecondStyleBold.Render(strings.TrimSpace(subject.SubjectName)),
					subject.Teacher.Name,
					visuals.MainStyle.Render(task),
				)
			}
		}

		hometasksRow[i] = day.String()
	}

	return datesList, hometasksRow
}

// build headers and rows for diary table
func displayDiaryTable(diary *api.DiaryResponse) error {
	headers, rows := diaryDates(&diary.Dates)

	// getting terminal width
	width, _, err := terminalSize()
	if err != nil {
		return fmt.Errorf("failed to get terminal size: %v", err)
	}

	table := table.New().
		Headers(headers...).
		Rows(rows).
		Wrap(true).
		Width(width - 1). // removing 1 character from width to fit into terminal window perfectly
		Border(lipgloss.ThickBorder()).BorderStyle(visuals.ThirdStyleBold)

	fmt.Println(table)

	return nil
}
