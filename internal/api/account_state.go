package api

import (
	"encoding/json"
	"fmt"
	"os"
)

// Account state
type AccountState struct {
	// Refresh token
	RefreshToken string `json:"refresh_token"`

	// Access token
	AccessToken string `json:"access_token"`

	// Student ID
	StudentID int `json:"student_id"`

	// User FIO
	FIO string `json:"FIO"`
}

// Save account state to file
func (s *AccountState) Save(filename string) error {
	encoded, _ := json.Marshal(s)

	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer f.Close()

	_, err = f.Write(encoded)
	if err != nil {
		return fmt.Errorf("failed to write file: %v", err)
	}

	return nil
}

// Load account state from filename
func (s *AccountState) Load(filename string) error {
	content, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			// creating empty one
			_, err = os.Create(filename)
			if err != nil {
				return fmt.Errorf("failed to create account file: %v", err)
			}
			content = make([]byte, 0)
		}

		if os.IsPermission(err) {
			return fmt.Errorf("cannot read file due to permission error: %v", err)
		}
	}

	if len(content) == 0 {
		return fmt.Errorf("no account found")
	}

	var state AccountState
	err = json.Unmarshal(content, &state)
	if err != nil {
		return fmt.Errorf("failed to unmarshal state: %v", err)
	}

	// log.Println("account_state.go", state)

	s.FIO = state.FIO
	s.AccessToken = state.AccessToken
	s.RefreshToken = state.RefreshToken
	s.StudentID = state.StudentID

	return nil
}
