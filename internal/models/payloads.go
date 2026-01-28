package models

type Payload any // not any, but maybe i'll change it

// generic struct
type LoginPayload struct {
	Username          string `json:"username"`
	Password          string `json:"password"`
	ExponentPushToken string `json:"exponentPushToken"`
}

// Can be used for diary and student perfomance
type DefaultPayload struct {
	EndDate   string `json:"end_date"`
	StartDate string `json:"start_date"`
	StudentID int    `json:"student_id"`
}

type GradesPayload struct {
	EndDate   string `json:"end_date"`
	StartDate string `json:"start_date"`
	StudentID int    `json:"student_id"`
	SubjectID int    `json:"subject_id"`
}

type RefreshTokenPayload struct {
	RefreshToken string `json:"refresh_token"`
}
