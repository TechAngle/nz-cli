package cli

import (
	"nz-cli/internal/api"
	"strings"
)

// structure which represents pair of headers and rows for table
type tablePair struct {
	headers []string
	rows    []string
}

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

// Divide headers and rows to pairs. returns pairs for both
func diaryDatesToPairs(headersList, rowsList []string) (pairs *[]tablePair) {
	pairsList := []tablePair{}

	var pairEnd int
	for i := 0; i < len(headersList); i += 2 {
		pairEnd += 2
		if pairEnd > len(headersList) {
			pairEnd = len(headersList)
		}

		pairsList = append(pairsList, tablePair{
			headers: headersList[i:pairEnd],
			rows:    rowsList[i:pairEnd],
		})
	}

	return &pairsList
}
