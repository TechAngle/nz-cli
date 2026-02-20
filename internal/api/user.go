package api

import (
	"fmt"
)

// Get perfomance stats
// Method: POST
func (c *NZAPIClient) Perfomance(payload DefaultPayload) (*PerfomanceResponse, error) {
	startDate, endDate, _ := validatePayloadDates(payload.StartDate, payload.EndDate)
	payload.StartDate = startDate
	payload.EndDate = endDate

	response, err := sendRequest[PerfomanceResponse](c, PostMethod, perfomanceEndpoint, payload)
	if err != nil {
		return nil, fmt.Errorf("failed to send perfomance request: %v", err)
	}

	return response, nil
}

// Get diary structure
// Method: POST
func (c *NZAPIClient) Diary(payload DefaultPayload) (*DiaryResponse, error) {
	startDate, endDate, _ := validatePayloadDates(payload.StartDate, payload.EndDate)
	payload.StartDate = startDate
	payload.EndDate = endDate

	response, err := sendRequest[DiaryResponse](c, PostMethod, diaryEndpoint, payload)
	if err != nil {
		return nil, fmt.Errorf("failed to send diary request: %v", err)
	}

	return response, nil
}

// Returns grades for specific subject id
// Method: POST
func (c *NZAPIClient) Grades(payload GradesPayload) (*GradesResponse, error) {
	startDate, endDate, _ := validatePayloadDates(payload.StartDate, payload.EndDate)
	payload.StartDate = startDate
	payload.EndDate = endDate

	// getting response for this shit
	response, err := sendRequest[GradesResponse](c, PostMethod, gradesEndpoint, payload)
	if err != nil {
		return nil, fmt.Errorf("failed to send grades request: %v", err)
	}

	return response, nil
}
