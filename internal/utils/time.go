package utils

import (
	"fmt"
	"nz-cli/internal/commons"
	"time"
)

func ValidatePayloadDates(startDate string, endDate string) (start string, end string, err error) {
	// parsing periods
	startTime, err := time.Parse(commons.DateFormat, startDate)
	if err != nil {
		return "", "", fmt.Errorf("failed to parse start date: %v", err)
	}
	endTime, err := time.Parse(commons.DateFormat, endDate)
	if err != nil {
		return "", "", fmt.Errorf("failed to parse end date: %v", err)
	}

	// if we got invalid sequence of dates we just switch them, why not
	if startTime.After(endTime) {
		startDate, endDate = endDate, startDate
	}

	return startDate, endDate, nil
}
