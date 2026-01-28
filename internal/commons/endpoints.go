package commons

// More endpoints you could find at https://github.com/Artemka1806/api-mobile.nz.ua
const (
	ApiEndpoint = "https://api-mobile.nz.ua/v2"

	LoginEndpoint        = "/user/login"
	TestEndpoint         = "/user/test" // Useless
	PerfomanceEndpoint   = "/schedule/student-performance"
	DiaryEndpoint        = "/schedule/diary"
	GradesEndpoint       = "/schedule/subject-grades"
	RefreshTokenEndpoint = "/user/refresh-token"

	// TODO: Add notifications endpoints

	TimetableEndpoint = "/schedule/timetable"          // useless, but i'll leave it here
	MarksListEndpoint = "/personnel-journal/mark-list" // also useless, but ok
)
