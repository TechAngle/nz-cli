package api

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"nz-cli/internal/models"

	"github.com/andybalholm/brotli"
)

type Method string

// Method
const (
	GetMethod  = "GET"
	PostMethod = "POST"
)

// sends request to nz api
/* NOTE: Requires endpoint as concantenation of original API endpoint and needed */
func (c *NZAPIClient) SendRequest(method Method, endpoint string, payload models.Payload, responsePtr models.ApiResponse) error {
	// encoding payload
	bodyBytes, _ := json.Marshal(payload)
	body := bytes.NewBuffer(bodyBytes)

	req, err := http.NewRequest(string(method), endpoint, body)
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
