package cli

import (
	"fmt"
	"nz-cli/internal/api"
	"nz-cli/internal/visuals"
	"strconv"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/list"
)

// Get grades and print a list of them
func (c *CLIClient) Grades(startDate string, endDate string, subjectId int) error {
	if startDate == "" || endDate == "" {
		return fmt.Errorf("invalid dates range: %s - %s", startDate, endDate)
	}

	grades, err := c.client.Grades(api.GradesPayload{
		StartDate: startDate,
		EndDate:   endDate,
		StudentID: c.client.Account().StudentID,
		SubjectID: subjectId,
	})
	if err != nil {
		return fmt.Errorf("failed to get grades: %v", err)
	}

	displayGradesList(grades)

	return nil
}

// display grades list based on response
func displayGradesList(grades *api.GradesResponse) {
	// creating new list
	list := list.New().ItemStyle(
		lipgloss.NewStyle().
			Align(lipgloss.Center).
			Bold(true).
			Background(visuals.MainStyle.GetBackground()),
	)

	// adding every mark with lesson date to the list
	for _, grade := range grades.Lessons {
		list.Item(
			fmt.Sprintf(
				"[%s] %s\t(%s)",
				visuals.ThirdStyle.Render(grade.LessonDate),
				visuals.MarkStyle(grade.Mark).Render(grade.Mark),
				grade.LessonType,
			),
		)
	}

	fmt.Printf("\tMissed Lessons: %s\n", visuals.SecondStyleBold.Underline(true).Render(strconv.Itoa(grades.NumberMissedLessons)))
	fmt.Println(list)
}
