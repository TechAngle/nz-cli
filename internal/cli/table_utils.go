package cli

import (
	"nz-cli/internal/api"
	"strings"
)

// Removes any " ", "." from subject name
func normalizeName(subjectName string) string {
	dotFree := strings.ReplaceAll(subjectName, ".", "")
	spaceFree := strings.TrimSpace(dotFree)

	return spaceFree
}

// Organise, join subjects and returns it
func normalizeSubjects(subjects []api.Subject) []api.Subject {
	subjectsList := []api.Subject{}

	previousSubjectName := ""

	for _, cSubject := range subjects {
		subjectName := normalizeName(cSubject.SubjectName)

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
