package api

// More endpoints you could find at https://github.com/Artemka1806/api-mobile.nz.ua
const (
	apiEndpoint = "https://api-mobile.nz.ua/v2"

	loginEndpoint        = "/user/login"
	perfomanceEndpoint   = "/schedule/student-performance"
	diaryEndpoint        = "/schedule/diary"
	gradesEndpoint       = "/schedule/subject-grades"
	refreshTokenEndpoint = "/user/refresh-token"

	notificationsEndpoint       = "/notification/"
	unreadNotificationsEndpoint = "/notification/unread-qty"

	testEndpoint      = "/user/test"                   // Useless
	timetableEndpoint = "/schedule/timetable"          // useless
	marksListEndpoint = "/personnel-journal/mark-list" // useless
)
