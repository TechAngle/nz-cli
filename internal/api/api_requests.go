package api

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"nz-cli/internal/commons"
	"nz-cli/internal/models"
	"nz-cli/internal/utils"

	"github.com/andybalholm/brotli"
)

// sends request to nz api
/* NOTE: Requires endpoint as concantenation of original API endpoint and needed */
func (c *NZAPIClient) SendRequest(endpoint string, payload models.Payload, responsePtr models.ApiResponse) error {
	// encoding payload
	bodyBytes, _ := json.Marshal(payload)
	body := bytes.NewBuffer(bodyBytes)

	req, err := http.NewRequest(http.MethodPost, endpoint, body)
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	// headers they use mainly
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Connection", "Keep-Alive")
	req.Header.Add("Accept-Charset", "utf-8 *;q=0.8")
	req.Header.Add("Accept-Encoding", "application/json")

	// adding access token if available
	if c.Authorized() {
		req.Header.Add("Authorization", "Bearer "+c.account.AccessToken)
	}

	res, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to do request: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid response: [%d] %s", res.StatusCode, res.Status)
	}

	var bodyReader io.Reader
	// log.Println(res.Header.Get("Content-Encoding"))
	switch res.Header.Get("Content-Encoding") {
	case "gzip":
		bodyGzip, err := gzip.NewReader(res.Body)
		if err != nil {
			return fmt.Errorf("failed to read response body: %v", err)
		}
		defer bodyGzip.Close()
		bodyReader = bodyGzip
	case "br":
		bodyReader = brotli.NewReader(res.Body)
	default:
		bodyReader = res.Body
	}

	bodyContent, err := io.ReadAll(bodyReader)
	if err != nil {
		return fmt.Errorf("failed to read body content: %v", err)
	}

	// unmarshalling body json content
	err = json.Unmarshal(bodyContent, responsePtr)
	if err != nil {

		return fmt.Errorf("failed to unmarshal body content: %v", err)
	}

	return nil
}

// logins and overwrites current account settings
func (c *NZAPIClient) Login(payload models.LoginPayload) error {
	var response models.LoginResponse
	err := c.SendRequest(commons.ApiEndpoint+commons.LoginEndpoint, payload, &response)
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
func (c *NZAPIClient) Perfomance(payload models.DefaultPayload) (*models.PerfomanceResponse, error) {
	startDate, endDate, _ := utils.ValidatePayloadDates(payload.StartDate, payload.EndDate)
	payload.StartDate = startDate
	payload.EndDate = endDate

	var response models.PerfomanceResponse
	err := c.SendRequest(commons.ApiEndpoint+commons.PerfomanceEndpoint, payload, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to send perfomance request: %v", err)
	}

	return &response, nil
}

// returns diary structure
func (c *NZAPIClient) Diary(payload models.DefaultPayload) (*models.DiaryResponse, error) {
	startDate, endDate, _ := utils.ValidatePayloadDates(payload.StartDate, payload.EndDate)
	payload.StartDate = startDate
	payload.EndDate = endDate

	var response models.DiaryResponse
	err := c.SendRequest(commons.ApiEndpoint+commons.DiaryEndpoint, payload, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to send diary request: %v", err)
	}

	return &response, nil
}

// returns grades for specific subject id
func (c *NZAPIClient) Grades(payload models.GradesPayload) (*models.GradesResponse, error) {
	startDate, endDate, _ := utils.ValidatePayloadDates(payload.StartDate, payload.EndDate)
	payload.StartDate = startDate
	payload.EndDate = endDate

	// getting response for this shit
	var response models.GradesResponse
	err := c.SendRequest(commons.ApiEndpoint+commons.GradesEndpoint, payload, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to send grades request: %v", err)
	}

	return &response, nil
}
