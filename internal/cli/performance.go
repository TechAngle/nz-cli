package cli

import (
	"fmt"
	"nz-cli/internal/api"
	"nz-cli/internal/utils"
	"nz-cli/internal/visuals"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

// Print performance
func (c *CLIClient) Performance(startDate string, endDate string) error {
	performance, err := c.client.Perfomance(api.DefaultPayload{
		StartDate: startDate,
		EndDate:   endDate,
	})
	if err != nil {
		return fmt.Errorf("failed to get perfomance: %v", err)
	}

	displayPerfomanceTable(performance)

	return nil
}

// Marks from subject.
// Returns slice of all marks and customized string representation.
func marksFromSubject(subject *api.Subject) ([]int, string) {
	marksList := make([]int, 0, len(subject.Marks))
	var marksRow strings.Builder

	for i, mark := range subject.Marks {
		// separator should be added if it is last element
		separator := ","
		if i == len(subject.Marks)-1 {
			separator = ""
		}

		// Adding year marks(non-zero)
		if strings.HasPrefix(strings.ToLower(mark.Type), "тем") {
			// log.Println("Tem:", mark.Value, mark.Type)
			if i, err := strconv.Atoi(mark.Value); err == nil {
				marksList = append(marksList, i)
			}
		}

		fmt.Fprintf(&marksRow, "%s[%s]%s ", visuals.MarkStyle(mark.Value).Render(mark.Value), mark.Type, separator)
	}

	return marksList, marksRow.String()
}

// Display perfomance table based on response
func displayPerfomanceTable(performance *api.PerfomanceResponse) error {
	// headers: Subject Name | Marks | Average of semester
	headers := []string{"Subject", "Marks", "Semester"}
	grades := [][]string{}

	// Code below contains such logic:
	// Initially, we extracting all subjects from subjects and mixing them with subjects with identical names.
	// Because nz.ua somehow returns different IDs for one subject.
	subjects := normalizeSubjects(performance.Subjects)

	for _, subject := range subjects {
		// gettimg marks from subject
		marks, marksRow := marksFromSubject(&subject)

		// adding average if it not 0, if not - just empty string
		// log.Println("[DEBUG] Marks:", marks)
		averageMark := utils.CalculateAverage[float32, int](marks)
		semesterMark := strconv.Itoa(int(averageMark))
		if averageMark == 0 {
			semesterMark = "" // if no average at the moment
		}

		// creating empty row
		subjectRow := []string{}

		// Setting subject name
		subjectRow = append(subjectRow, fmt.Sprintf("%s [ID: %s]", subject.SubjectName, subject.SubjectID))
		// adding marks row
		subjectRow = append(subjectRow, marksRow)
		// adding semester mark
		subjectRow = append(subjectRow, fmt.Sprintf("%s", semesterMark))

		grades = append(grades, subjectRow)
	}

	// getting width of terminal
	width, _, err := terminalSize()
	if err != nil {
		return fmt.Errorf("[performance] failed to get terminal size:")
	}

	table := table.New().
		Wrap(true).
		Width(width - 1). // removing 1 character from width to fit into terminal window perfectly
		Border(lipgloss.NormalBorder()).
		BorderStyle(visuals.ThirdStyleBold).
		Headers(headers...).
		Rows(grades...)

	fmt.Println(table.Render())

	return nil
}
