package api

type APIPayload any

// /login payload
type LoginPayload struct {
	Username          string `json:"username"`
	Password          string `json:"password"`
	ExponentPushToken string `json:"exponentPushToken"`
}

// refresh-token payload
type RefreshTokenPayload struct {
	RefreshToken string `json:"refresh_token"`
}

// /grades payload
type GradesPayload struct {
	EndDate   string `json:"end_date"`
	StartDate string `json:"start_date"`
	StudentID int    `json:"student_id"`
	SubjectID int    `json:"subject_id"`
}

// Can be used for diary and student perfomance
type DefaultPayload struct {
	EndDate   string `json:"end_date"`
	StartDate string `json:"start_date"`
	StudentID int    `json:"student_id"`
}
