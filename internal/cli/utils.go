package cli

import (
	"nz-cli/internal/models"
	"strings"
)

// Removes any " ", "." from subject name
func NormalizeName(subjectName string) string {
	dotFree := strings.ReplaceAll(subjectName, ".", "")
	spaceFree := strings.TrimSpace(dotFree)

	return spaceFree
}

// Organise, join subjects and returns it
func NormalizeSubjects(subjects []models.Subject) []models.Subject {
	subjectsList := []models.Subject{}

	previousSubjectName := ""

	for _, cSubject := range subjects {
		subjectName := NormalizeName(cSubject.SubjectName)

		if subjectName == previousSubjectName {
			// using ptr to the previous one
			subject := &subjectsList[len(subjectsList)-1]

			// adding marks
			subject.Marks = append(subject.Marks, cSubject.Marks...)
			subject.SubjectID = strings.Join([]string{subject.SubjectID, cSubject.SubjectID}, ", ")

			continue
		}

		subjectsList = append(subjectsList, cSubject)
		previousSubjectName = subjectName
	}

	return subjectsList
}
