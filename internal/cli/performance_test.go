package cli

import (
	"nz-cli/internal/api"
	"testing"
)

func TestMarksFromSubject(t *testing.T) {
	subject := &api.Subject{
		Marks: []api.Mark{
			{
				Value: "8",
				Type:  "К/р",
			},
			{
				Value: "8",
				Type:  "К/р",
			},
			{
				Value: "7",
				Type:  "К/р",
			},
			{
				Value: "Н", // must be ignored
				Type:  "К/р",
			},
		},
	}

	marks, _ := marksFromSubject(subject)
	if len(marks) != 3 {
		t.Fatalf("Marks slice len must be 3 (%d)", len(marks))
	}
}
