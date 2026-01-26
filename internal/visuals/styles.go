package visuals

import "github.com/charmbracelet/lipgloss"

/*
 *     "pastel_white": "#FCF8F8",
 "pastel_pink": "#FBEFEF",
 "cream": "#F9DFDF",
 "peach": "#F5AFAF",
 "error": "#e11d62",
*/

// HEX colour codes
const (
	mainCode   string = "#FCF8F8" // pastel white
	secondCode string = "#FBEFEF" // pastel pink
	thirdCode  string = "#F9DFDF" // cream code
	fourthCode string = "#F5AFAF" // peach code
)

const (
	gradeBadCode  string = "#ff6b6b" // Pastel red (1-3)
	gradeLowCode  string = "#f9ad6a" // Dusty orange (4-6)
	gradeMidCode  string = "#d4e157" // Sandy lime (7-9)
	gradeHighCode string = "#8ce99a" // Mint green (10-12)
)

// pastel white style
var (
	MainStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color(mainCode))
	MainStyleBold    = lipgloss.NewStyle().Foreground(lipgloss.Color(mainCode)).Bold(true)
	MainStyleReverse = lipgloss.NewStyle().Foreground(lipgloss.Color("#000000")).Background(lipgloss.Color(mainCode))
)

// pastel pink style
var (
	SecondStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color(secondCode))
	SecondStyleBold    = lipgloss.NewStyle().Foreground(lipgloss.Color(secondCode)).Bold(true)
	SecondStyleReverse = lipgloss.NewStyle().Foreground(lipgloss.Color("#000000")).Background(lipgloss.Color(secondCode))
)

// cream style
var (
	ThirdStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color(thirdCode))
	ThirdStyleBold    = lipgloss.NewStyle().Foreground(lipgloss.Color(thirdCode)).Bold(true)
	ThirdStyleReverse = lipgloss.NewStyle().Foreground(lipgloss.Color("#000000")).Background(lipgloss.Color(thirdCode))
)

// peach style
var (
	FourthStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color(fourthCode))
	FourthStyleBold    = lipgloss.NewStyle().Foreground(lipgloss.Color(fourthCode)).Bold(true)
	FourthStyleReverse = lipgloss.NewStyle().Foreground(lipgloss.Color("#000000")).Background(lipgloss.Color(fourthCode))
)

// error style
var (
	ErrorStyle = lipgloss.NewStyle().Background(lipgloss.Color("#e11d62"))
)

var (
	GradeBadStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color(gradeBadCode)).Bold(true)
	GradeLowStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color(gradeLowCode))
	GradeMidStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color(gradeMidCode))
	GradeHighStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(gradeHighCode)).Bold(true)
)

// Marks styles
func MarkStyle(mark string) lipgloss.Style {
	switch mark {
	// if it is mark - return its style
	case "1", "2", "3":
		return GradeBadStyle
	case "4", "5", "6":
		return GradeLowStyle
	case "7", "8", "9":
		return GradeMidStyle
	case "10", "11", "12":
		return GradeHighStyle

	// if not - default bold style
	default:
		return MainStyleBold
	}
}
