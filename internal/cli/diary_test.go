package cli

import (
	"nz-cli/internal/api"
	"testing"
)

func TestDiaryDates(t *testing.T) {
	// TODO: Add test calls
	dates := &api.DiaryResponse{
		Dates: []api.Date{
			{
				Date: "2026-01-01",
			},
			{
				Date: "2026-01-01",
			},
			{
				Date: "2026-01-01",
			},
		},
	}

	datesList, _ := diaryDates(&dates.Dates)
	if len(datesList) != 3 {
		t.Fatalf("dates list must be 3!")
	}
}
