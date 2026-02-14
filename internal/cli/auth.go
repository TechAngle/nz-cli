package cli

import (
	"fmt"
	"log"
	"nz-cli/internal/api"
)

// Is client authorized
func (c *CLIClient) IsAuthorized() bool {
	return c.client.Authorized()
}

// Login to system
func (c *CLIClient) Login(username string, password string) error {
	if c.client.Authorized() {
		log.Println("You're already logged to system!")
		return nil
	}

	if username == "" || password == "" {
		return fmt.Errorf("invalid credentials")
	}

	err := c.client.Login(api.LoginPayload{
		Username: username,
		Password: password,
	})
	if err != nil {
		return fmt.Errorf("failed to login: %v", err)
	}

	// saving immediately
	err = c.client.SaveSession()
	if err != nil {
		fmt.Println("Failed to save session:", err)
		return nil
	}

	return nil
}

// restore session
func (c *CLIClient) RestoreSession() error {
	err := c.client.LoadAccount()
	if err != nil {
		return fmt.Errorf("failed to load account: %v", err)
	}

	return nil
}

// Update refresh token
func (c *CLIClient) RefreshToken() error {
	accessToken, err := c.client.RefreshToken(api.RefreshTokenPayload{
		RefreshToken: c.client.Account().RefreshToken,
	})
	if err != nil {
		return fmt.Errorf("failed to refresh token: %v", err)
	}

	c.client.SetNewAccessToken(accessToken)

	return nil
}
