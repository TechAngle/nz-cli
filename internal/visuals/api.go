package visuals

import (
	"fmt"
	"log"
	"nz-cli/internal/api"
)

// Login to system
func (c *TUI) Login() error {
	// checking user data
	if c.userData.username == "" || c.userData.password == "" {
		return fmt.Errorf("invalid username(%s) or password(%s)", c.userData.username, c.userData.password)
	}
	log.Println("User data is valid, logging in....")

	err := c.client.Login(api.LoginPayload{
		Username: c.userData.username,
		Password: c.userData.password,
	})
	if err != nil {
		return fmt.Errorf("failed to log in: %v", err)
	}

	return nil
}
