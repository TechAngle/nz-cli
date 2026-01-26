package visuals

import (
	"fmt"
	"maps"
	"nz-cli/internal/models"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	// main command prompt
	commandInput textinput.Model

	// username input
	usernameInput textinput.Model
	passwordInput textinput.Model

	// current output of command
	currentOutput string

	// focus
	focusIndex int

	// current panel
	panel models.Panel
}

func initialModel() model {
	// generating commands slice
	var commandsSlice []string
	for command := range maps.Keys(commands) {
		commandsSlice = append(commandsSlice, command)
	}

	// command prompt entry
	commandInputModel := textinput.New()
	commandInputModel.Placeholder = "help"
	commandInputModel.Focus()
	commandInputModel.CharLimit = 16
	commandInputModel.SetSuggestions(commandsSlice)
	commandInputModel.ShowSuggestions = true
	commandInputModel.TextStyle = ThirdStyle
	commandInputModel.PromptStyle = MainStyleBold

	// username input entry
	username := textinput.New()
	username.Placeholder = "Username"
	username.TextStyle = SecondStyleBold
	username.CharLimit = 32

	// password input entry
	password := textinput.New()
	password.Placeholder = "Password"
	password.TextStyle = SecondStyleBold
	password.CharLimit = 32

	m := model{
		commandInput:  commandInputModel,
		usernameInput: username,
		passwordInput: password,
	}

	return m
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) View() string {
	var s strings.Builder // initial value

	// showing banner
	s.WriteString(FourthStyle.Render(Banner()) + "\n")

	// showing info
	s.WriteString(strings.Repeat(" ", 15) + AccountString() + "\n")

	// displaying current output
	s.WriteString(m.currentOutput + "\n")

	switch m.panel {
	case models.MainPanel:
		// showing command input
		s.WriteString(m.commandInput.View())
	case models.LoginPanel:

	}

	return s.String()
}

func (m *model) setError(error string) {
	m.currentOutput = ErrorStyle.Render(error) + "\n"
}

func (m *model) setOutput(output string) {
	m.currentOutput = output
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		// quit
		case "ctrl+c", "esc":
			return m, tea.Quit

		// command enter
		case "enter":
			// m.commandInput.Blur()
			commandInput := m.commandInput.Value()
			// log.Println("command inputed:", commandInput)
			if len(strings.TrimSpace(commandInput)) == 0 {
				return m, cmd
			}

			command := strings.Split(commandInput, " ")[0]
			switch command {
			case "help":
				m.currentOutput = help()
			case "quit":
				err := client.SaveSession()
				if err != nil {
					m.setError(err.Error())
				}

				return m, tea.Quit
			}

			m.commandInput.SetValue("")

			return m, cmd
		}
	}

	m.commandInput, cmd = m.commandInput.Update(msg)

	return m, cmd
}

func AccountString() string {
	if client.Authorized() {
		acc := client.Account()
		return (fmt.Sprintf("Logged in as %s (Student ID: %d)", acc.FIO, acc.StudentID))
	} else {
		return "No account!"
	}
}
