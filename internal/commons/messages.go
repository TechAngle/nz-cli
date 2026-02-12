package commons

// login flags usages
const (
	LoginFlagUsage    = "Login to the system. Should be set with --username and --password flags"
	UsernameFlagUsage = "Username. Required if -login argument is set"
	PasswordFlagUsage = "Password. Required if -login argument is set"
)

// additional flags usages
const (
	StartDateFlagUsage = "Start date. Required always. Set it if you wanna get any information starting from specific date."
	EndDateFlagUsage   = "End date.  Required always. Set it if you wanna get any information up to specific date."
	SubjectIdFlagUsage = "Subject ID. Required if you wanna get grades."
)

// date flags usages
const (
	TomorrowFlagUsage  = "Use tomorrow's date (start-date and end-date will be overwritten)"
	YesterdayFlagUsage = "Use yesterday's date (start-date and end-date will be overwritten)"
)

// client flags usages
const (
	DiaryFlagUsage      = "Show diary"
	GradesFlagUsage     = "Show grades (requires `-subject-id` flag)"
	PerfomanceFlagUsage = "Show perfomance"
)

// mode flags usages
const (
	DataFlagUsage = "Get only data from API and format it (API mode)"
	TUIFlagUsage  = "Run TUI (Terminal User Interface) (Experimental)"
)
