package cli

import (
	"nz-cli/internal/api"
	"testing"
)

// Test name normalization
func TestNormalizeName(t *testing.T) {
	extraStops := "Історія України.."
	extraStopsResult := "Історія України"

	extraWhitespaces := "Українська мова "
	extraWhitespacesResult := "Українська мова"

	stopsSanitized := normalizeName(extraStops)
	spacesSanitized := normalizeName(extraWhitespaces)

	if stopsSanitized != extraStopsResult {
		t.Fatalf("invalid sanitized stops: %s", stopsSanitized)
	}

	t.Logf("Start: %s -> Final: %s (=%s)", extraStops, stopsSanitized, extraStopsResult)

	if spacesSanitized != extraWhitespacesResult {
		t.Fatalf("invalid sanitized whitespaces: %s", spacesSanitized)
	}

	t.Logf("Start: %s -> Final: %s (=%s)", extraWhitespaces, spacesSanitized, extraWhitespacesResult)
}

// Test subjects normalization
func TestNormalizeSubjects(t *testing.T) {
	testSubjects := []api.Subject{
		{
			SubjectName: "test1",
			SubjectID:   "12345",
			Marks: []api.Mark{
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
			Marks: []api.Mark{
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

	subjectsList := normalizeSubjects(testSubjects)

	// must be one element
	if len(subjectsList) != 1 {
		t.Fatalf("Subject list must contain only one element(%d)!", len(subjectsList))
	}

	t.Logf("Subject IDs: %s", subjectsList[0].SubjectID)
}
