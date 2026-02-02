package client

import (
	"fmt"
	"log"
	"nz-cli/internal/api"
	"nz-cli/internal/models"
	"nz-cli/internal/utils"
	"nz-cli/internal/visuals"
	"os"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/list"
	"github.com/charmbracelet/lipgloss/table"
	"golang.org/x/term"
)

type Client struct {
	client *api.NZAPIClient
}

func (c *Client) IsAuthorized() bool {
	return c.IsAuthorized()
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
		subjectRow = append(subjectRow, fmt.Sprintf("%s [ID: %s]", subject.SubjectName, subject.SubjectID))

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

			// log.Println(mark)

			fmt.Fprintf(&marksRow, "%s[%s]%s ", visuals.MarkStyle(mark.Value).Render(mark.Value), mark.Type, separator)
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
		Wrap(true).
		Border(lipgloss.NormalBorder()).
		BorderStyle(visuals.ThirdStyleBold).
		Headers(headers...).
		Rows(grades...)

	fmt.Println(table.Render())

	return nil
}

// Update refresh token
func (c *Client) RefreshToken() error {
	accessToken, err := c.client.RefreshToken(models.RefreshTokenPayload{
		RefreshToken: c.client.Account().RefreshToken,
	})
	if err != nil {
		return fmt.Errorf("failed to refresh token: %v", err)
	}

	c.client.SetNewAccessToken(accessToken)

	return nil
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

	l := list.New().ItemStyle(
		lipgloss.NewStyle().
			Align(lipgloss.Center).
			Bold(true).
			Background(visuals.MainStyle.GetBackground()),
	)

	// adding every mark with lesson date to the list
	for _, grade := range grades.Lessons {
		l.Item(
			fmt.Sprintf(
				"[%s] %s\t(%s)",
				visuals.ThirdStyle.Render(grade.LessonDate),
				visuals.MarkStyle(grade.Mark).Render(grade.Mark),
				grade.LessonType,
			),
		)
	}

	fmt.Printf("\tMissed Lessons: %s\n", visuals.SecondStyleBold.Underline(true).Render(strconv.Itoa(grades.NumberMissedLessons)))
	fmt.Println(l)

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

	// getting current terminal size
	fd := int(os.Stdout.Fd())
	width, _, err := term.GetSize(fd)
	if err != nil {
		return fmt.Errorf("failed to get terminal size: %v", err)
	}

	table := table.New().
		Headers(datesList...).
		Rows(hometasksRow).
		Wrap(true).
		Width(width).
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
