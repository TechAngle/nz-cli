package client

import (
	"fmt"
	"log"
	"nz-cli/internal/api"
	"nz-cli/internal/models"
	"nz-cli/internal/utils"
	"nz-cli/internal/visuals"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

type Client struct {
	client *api.NZAPIClient
}

// Print perfomance
func (c *Client) Perfomance(startDate string, endDate string) error {
	if !c.client.Authorized() {
		return fmt.Errorf("not authorized")
	}

	if startDate == "" || endDate == "" {
		return fmt.Errorf("invalid dates range: %s - %s", startDate, endDate)
	}

	perfomance, err := c.client.Perfomance(models.DefaultPayload{
		StartDate: startDate,
		EndDate:   endDate,
	})
	if err != nil {
		return fmt.Errorf("failed to get perfomance: %v", err)
	}

	// headers: Subject Name | Marks | Average
	headers := []string{"Subject", "Marks", "Semestr"}
	grades := [][]string{}
	for _, subject := range perfomance.Subjects {
		subjectRow := []string{}
		average := []int{}
		var marksRow strings.Builder

		// Setting subject name
		subjectRow = append(subjectRow, subject.SubjectName)

		// adding marks
		for i, mark := range subject.Marks {
			separator := ","
			if i == len(subject.Marks)-1 {
				separator = ""
			}

			// Adding year marks(non-zero)
			if strings.HasPrefix(strings.ToLower(mark.Type), "тем") && mark.Value != "Н" {
				// log.Println("Tem:", mark.Value, mark.Type)
				if i, err := strconv.Atoi(mark.Value); err == nil {
					average = append(average, i)
				}
			}

			log.Println(mark)

			fmt.Fprintf(&marksRow, "%s[%s]%s ", mark.Value, mark.Type, separator)
		}

		// adding marks row
		subjectRow = append(subjectRow, marksRow.String())

		// adding average
		averageMark := utils.CalculateAverage[float32, int](average)
		// log.Println("Average:", averageMark)
		subjectRow = append(subjectRow, fmt.Sprintf("%d", int(averageMark)))
		grades = append(grades, subjectRow)
	}

	table := table.New().
		Headers(headers...).
		Rows(grades...).
		Wrap(true).
		Width(150).
		Border(lipgloss.ThickBorder()).BorderStyle(visuals.ThirdStyleBold)

	fmt.Println(table.Render())

	return nil
}

func (c *Client) IsAuthorized() bool {
	return c.IsAuthorized()
}

// Print grades
func (c *Client) Grades(startDate string, endDate string, subjectId int) error {
	if !c.client.Authorized() {
		return fmt.Errorf("not authorized")
	}

	if startDate == "" || endDate == "" {
		return fmt.Errorf("invalid dates range: %s - %s", startDate, endDate)
	}

	grades, err := c.client.Grades(models.GradesPayload{
		StartDate: startDate,
		EndDate:   endDate,
		StudentID: c.client.Account().StudentID,
		SubjectID: subjectId,
	})
	if err != nil {
		return fmt.Errorf("failed to get grades: %v", err)
	}

	var s strings.Builder
	s.WriteString("Grades [%s]:\n\t")

	for i, grade := range grades.Lessons {
		separator := ","

		// changing separator when we in the end
		if i == len(grades.Lessons)-1 {
			separator = "."
		}

		fmt.Fprintf(&s,
			"[%s] %s (%s)%s ",
			grade.LessonDate,
			grade.Mark,
			grade.LessonType,
			separator,
		)
	}

	fmt.Println(s.String())

	return nil
}

// Get diary
func (c *Client) Diary(startDate string, endDate string) error {
	if !c.client.Authorized() {
		return fmt.Errorf("not authorized")
	}

	if startDate == "" || endDate == "" {
		return fmt.Errorf("invalid dates range: %s - %s", startDate, endDate)
	}

	diary, err := c.client.Diary(models.DefaultPayload{
		StartDate: startDate,
		EndDate:   endDate,
		StudentID: c.client.Account().StudentID,
	})
	if err != nil {
		return fmt.Errorf("failed to get diary: %v", err)
	}

	datesList := []string{}
	hometasksRow := make([]string, len(diary.Dates))

	// going through dates
	for i, date := range diary.Dates {
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

	table := table.New().
		Headers(datesList...).
		Rows(hometasksRow).
		Wrap(true).
		Width(100).
		Border(lipgloss.ThickBorder()).BorderStyle(visuals.ThirdStyleBold)

	fmt.Println(table)

	return nil
}

// Login to system
func (c *Client) Login(username string, password string) error {
	if c.client.Authorized() {
		log.Println("You're already logged to system!")
		return nil
	}

	if username == "" || password == "" {
		return fmt.Errorf("invalid credentials")
	}

	err := c.client.Login(models.LoginPayload{
		Username: username,
		Password: password,
	})
	if err != nil {
		return fmt.Errorf("failed to login: %v", err)
	}

	// saving immediately
	err = c.client.SaveSession()
	if err != nil {
		fmt.Println("Failed to save session:", err)
		return nil
	}

	return nil
}

func (c *Client) RestoreSession() error {
	err := c.client.LoadAccount()
	if err != nil {
		return fmt.Errorf("failed to load account: %v", err)
	}

	return nil
}

func NewClient() (*Client, error) {
	client, err := api.NewApiClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}

	return &Client{
		client: client,
	}, nil
}
