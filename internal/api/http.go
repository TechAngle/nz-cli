package api

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"

	http "github.com/bogdanfinn/fhttp"

	"github.com/andybalholm/brotli"
)

type Method string

// Method
const (
	GetMethod  = "GET"
	PostMethod = "POST"
)

/*
Send request to nz api and writes response to modelResponse.

NOTE: Requires endpoint as concatenation of original API endpoint and needed

NOTE: if method is GET, payload can be nil, because it won't be marshaled
*/
func (c *NZAPIClient) SendRequest(method Method, endpoint string, payload Payload, modelResponse ApiResponse) (err error) {
	var body *bytes.Buffer
	if method == PostMethod {
		body, err = marshalPayload(payload)
	}

	req, err := newRequest(method, endpoint, body)
	c.addHeaders(req)

	apiResponse, err := c.sendRequest(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer apiResponse.Body.Close()

	if err := readApiResponse(apiResponse, &modelResponse); err != nil {
		return fmt.Errorf("failed to read api response: %v", err)
	}

	return nil
}

// Adds headers to request. If client is authorized - adds Authorization Bearer
func (c *NZAPIClient) addHeaders(req *http.Request) {
	// headers they use mainly
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Connection", "Keep-Alive")
	req.Header.Add("Accept-Charset", "utf-8 *;q=0.8")
	req.Header.Add("Accept-Encoding", "application/json")
	req.Header.Add("User-Agent", UserAgent)

	// adding access token if available
	if c.Authorized() {
		req.Header.Add("Authorization", "Bearer "+c.account.AccessToken)
	}
}

// Sends request
func (c *NZAPIClient) sendRequest(req *http.Request) (*http.Response, error) {
	res, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to do request: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid response: [%d] %s", res.StatusCode, res.Status)
	}

	return res, nil
}

// Prepare request body
func marshalPayload(payload Payload) (*bytes.Buffer, error) {
	// encoding payload
	body := new(bytes.Buffer)
	bodyBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %v", err)
	}

	body = bytes.NewBuffer(bodyBytes)

	return body, nil
}

// Get body reader from response
func getBodyReader(res *http.Response) (*io.Reader, error) {
	var bodyReader io.Reader

	// log.Println(res.Header.Get("Content-Encoding"))
	switch res.Header.Get("Content-Encoding") {
	case "gzip":
		bodyGzip, err := gzip.NewReader(res.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read response body: %v", err)
		}
		defer bodyGzip.Close()
		bodyReader = bodyGzip
	case "br":
		bodyReader = brotli.NewReader(res.Body)
	default:
		bodyReader = res.Body
	}

	return &bodyReader, nil
}

// Reads body. If error occurred, body will be nil.
func readBody(res *http.Response) ([]byte, error) {
	bodyReader, err := getBodyReader(res)
	if err != nil {
		return nil, fmt.Errorf("failed to get body reader: %v", err)
	}

	bodyContent, err := io.ReadAll(*bodyReader)
	if err != nil {
		return nil, fmt.Errorf("failed to read body content: %v", err)
	}

	return bodyContent, nil
}

// Reads api response and unmarshals it to response pointer
func readApiResponse(res *http.Response, modelResponse ApiResponse) error {
	bodyContent, err := readBody(res)
	if err != nil {
		return fmt.Errorf("failed to read body: %v", err)
	}

	// unmarshalling body json content
	err = json.Unmarshal(bodyContent, &modelResponse)
	if err != nil {
		return fmt.Errorf("failed to unmarshal body content: %v", err)
	}

	return nil
}

// Initializates new request
func newRequest(method Method, endpoint string, body *bytes.Buffer) (*http.Request, error) {
	req, err := http.NewRequest(string(method), endpoint, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	return req, nil
}
