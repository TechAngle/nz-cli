package api

import (
	"fmt"
)

// Logins and overwrites current account settings
// Method: POST
func (c *NZAPIClient) Login(payload LoginPayload) error {
	var response LoginResponse
	err := c.SendRequest(PostMethod, apiEndpoint+loginEndpoint, payload, &response)
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

// Refresh current access token.
// Returns new access token
// Method: POST
func (c *NZAPIClient) RefreshToken(payload RefreshTokenPayload) (string, error) {
	var response RefreshTokenResponse
	err := c.SendRequest(PostMethod, apiEndpoint+refreshTokenEndpoint, payload, &response)
	if err != nil {
		return "", fmt.Errorf("failed to refresh token request: %v", err)
	}

	// if some shit occurred on their side
	if response.ErrorMessage != "" {
		return "", fmt.Errorf("failed to refresh token (nz side): %s", response.ErrorMessage)
	}

	return response.NewAccessToken, nil
}
