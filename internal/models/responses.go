package models

type ApiResponse any

type LoginResponse struct {
	AccessToken     string `json:"access_token"`
	RefreshToken    string `json:"refresh_token"`
	ExpiresToken    int    `json:"expires_token"`
	EmailHash       string `json:"email_hash"`
	StudentID       int    `json:"student_id"`
	ClassName       string `json:"class_name"`
	ClassManagerFio string `json:"class_manager_fio"`
	Fio             string `json:"FIO"`
	Avatar          struct {
		ImageURL string `json:"image_url"`
		Datetime int    `json:"datetime"`
	} `json:"avatar"`
	Permissions struct {
		IsuoNzportalChildren []string `json:"isuo_nzportal_children"`
	} `json:"permissions"`
	ErrorMessage string `json:"error_message"`
}

type PerfomanceResponse struct {
	Missed struct {
		Days    int `json:"days"`
		Lessons int `json:"lessons"`
	} `json:"missed"`
	Subjects []struct {
		SubjectID        string `json:"subject_id"`
		SubjectName      string `json:"subject_name"`
		SubjectShortname string `json:"subject_shortname"`
		Marks            []struct {
			Value string `json:"value"`
			Type  string `json:"type"`
		} `json:"marks"`
	} `json:"subjects"`
	ErrorMessage string `json:"error_message"`
}

type DiaryResponse struct {
	Dates []struct {
		Date  string `json:"date"`
		Calls []struct {
			CallID        int    `json:"call_id"`
			CallNumber    int    `json:"call_number"`
			CallTimeStart string `json:"call_time_start"`
			CallTimeEnd   string `json:"call_time_end"`
			Subjects      []struct {
				Lesson []struct {
					Type    string `json:"type"`
					Mark    string `json:"mark"`
					Comment string `json:"comment"`
				} `json:"lesson"`
				Hometask                 []string `json:"hometask"`
				DistanceHometaskID       any      `json:"distance_hometask_id"`
				DistanceHometaskIsClosed any      `json:"distance_hometask_is_closed"`
				SubjectName              string   `json:"subject_name"`
				Room                     string   `json:"room"`
				Teacher                  struct {
					ID   int    `json:"id"`
					Name string `json:"name"`
				} `json:"teacher"`
			} `json:"subjects"`
		} `json:"calls"`
	} `json:"dates"`
	ErrorMessage string `json:"error_message"`
}

type GradesResponse struct {
	ErrorMessage string `json:"error_message"`
	Lessons      []struct {
		Comment    string `json:"comment"`
		LessonDate string `json:"lesson_date"`
		LessonID   string `json:"lesson_id"`
		LessonType string `json:"lesson_type"`
		Mark       string `json:"mark"`
		Subject    string `json:"subject"`
	} `json:"lessons"`
	NumberMissedLessons int `json:"number_missed_lessons"`
}

type RefreshTokenResponse struct {
	NewAccessToken string `json:"access_token"`
	ErrorMessage   string `json:"error_message"`
}
