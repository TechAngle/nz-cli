package commons

import (
	"fmt"
	"os"
)

func GetConfigPath() string {
	path, err := os.UserConfigDir()
	if err != nil {
		fmt.Println("Failed to get config dir:", err)
		return ""
	}

	return path
}
