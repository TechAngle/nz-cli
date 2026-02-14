package api

import (
	"fmt"
	"os"
	"path/filepath"
)

var (
	SessionCookiesBase = filepath.Join(getConfigPath(), "nzCookies.json")
	AccountStateBase   = filepath.Join(getConfigPath(), "nzAccountState.json")
)

const (
	// User-Agent Header
	UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36"
)

const (
	// Date format which nz.ua uses in payloads
	DateFormat string = "2006-01-02"
)

// Get config path for files
func getConfigPath() string {
	path, err := os.UserConfigDir()
	if err != nil {
		fmt.Println("Failed to get config dir:", err)
		return ""
	}

	return path
}
