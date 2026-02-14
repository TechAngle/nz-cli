package api

import (
	"fmt"
	"time"
)

// Validates dates.
// If end date was before start date - it switches them.
func ValidatePayloadDates(startDate string, endDate string) (start string, end string, err error) {
	// parsing periods
	startTime, err := time.Parse(DateFormat, startDate)
	if err != nil {
		return "", "", fmt.Errorf("failed to parse start date: %v", err)
	}
	endTime, err := time.Parse(DateFormat, endDate)
	if err != nil {
		return "", "", fmt.Errorf("failed to parse end date: %v", err)
	}

	// if we got invalid sequence of dates we just switch them, why not
	if startTime.After(endTime) {
		startDate, endDate = endDate, startDate
	}

	return startDate, endDate, nil
}

// Check if their error message is not empty
func IsNZError(errorMessage string) bool {
	return errorMessage != ""
}
