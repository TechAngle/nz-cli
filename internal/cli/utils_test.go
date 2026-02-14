package cli

import (
	"nz-cli/internal/models"
	"testing"
)

func TestNormalizeName(t *testing.T) {
	extraStops := "Історія України.."
	extraStopsResult := "Історія України"

	extraWhitespaces := "Українська мова "
	extraWhitespacesResult := "Українська мова"

	stopsSanitized := NormalizeName(extraStops)
	spacesSanitized := NormalizeName(extraWhitespaces)

	if stopsSanitized != extraStopsResult {
		t.Fatalf("invalid sanitized stops: %s", stopsSanitized)
	}

	t.Logf("Start: %s -> Final: %s (=%s)", extraStops, stopsSanitized, extraStopsResult)

	if spacesSanitized != extraWhitespacesResult {
		t.Fatalf("invalid sanitized whitespaces: %s", spacesSanitized)
	}

	t.Logf("Start: %s -> Final: %s (=%s)", extraWhitespaces, spacesSanitized, extraWhitespacesResult)
}

func TestNormalizeSubjects(t *testing.T) {
	testSubjects := []models.Subject{
		{
			SubjectName: "test1",
			SubjectID:   "12345",
			Marks: []models.Mark{
				{
					Value: "1",
				},
				{
					Value: "2",
				},
				{
					Value: "3",
				},
			},
		},
		{
			SubjectName: "test1",
			SubjectID:   "12346", // different number
			Marks: []models.Mark{
				{
					Value: "1",
				},
				{
					Value: "2",
				},
				{
					Value: "3",
				},
			},
		},
	}

	subjectsList := NormalizeSubjects(testSubjects)

	// must be one element
	if len(subjectsList) != 1 {
		t.Fatalf("Subject list must contain only one element(%d)!", len(subjectsList))
	}

	t.Logf("Subject IDs: %s", subjectsList[0].SubjectID)
}
