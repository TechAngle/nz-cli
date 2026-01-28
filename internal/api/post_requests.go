package api

import (
	"fmt"
	"nz-cli/internal/commons"
	"nz-cli/internal/models"
	"nz-cli/internal/utils"
)

// logins and overwrites current account settings
// Method: POST
func (c *NZAPIClient) Login(payload models.LoginPayload) error {
	var response models.LoginResponse
	err := c.SendRequest(PostMethod, commons.ApiEndpoint+commons.LoginEndpoint, payload, &response)
	if err != nil {
		return fmt.Errorf("failed to login: %v", err)
	}

	account := &AccountState{
		FIO:          response.Fio,
		AccessToken:  response.AccessToken,
		RefreshToken: response.RefreshToken,
		StudentID:    response.StudentID,
	}

	c.account = account

	return nil
}

// Get perfomance stats
// Method: POST
func (c *NZAPIClient) Perfomance(payload models.DefaultPayload) (*models.PerfomanceResponse, error) {
	startDate, endDate, _ := utils.ValidatePayloadDates(payload.StartDate, payload.EndDate)
	payload.StartDate = startDate
	payload.EndDate = endDate

	var response models.PerfomanceResponse
	err := c.SendRequest(PostMethod, commons.ApiEndpoint+commons.PerfomanceEndpoint, payload, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to send perfomance request: %v", err)
	}

	return &response, nil
}

// returns diary structure
// Method: POST
func (c *NZAPIClient) Diary(payload models.DefaultPayload) (*models.DiaryResponse, error) {
	startDate, endDate, _ := utils.ValidatePayloadDates(payload.StartDate, payload.EndDate)
	payload.StartDate = startDate
	payload.EndDate = endDate

	var response models.DiaryResponse
	err := c.SendRequest(PostMethod, commons.ApiEndpoint+commons.DiaryEndpoint, payload, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to send diary request: %v", err)
	}

	return &response, nil
}

// returns grades for specific subject id
// Method: POST
func (c *NZAPIClient) Grades(payload models.GradesPayload) (*models.GradesResponse, error) {
	startDate, endDate, _ := utils.ValidatePayloadDates(payload.StartDate, payload.EndDate)
	payload.StartDate = startDate
	payload.EndDate = endDate

	// getting response for this shit
	var response models.GradesResponse
	err := c.SendRequest(PostMethod, commons.ApiEndpoint+commons.GradesEndpoint, payload, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to send grades request: %v", err)
	}

	return &response, nil
}
