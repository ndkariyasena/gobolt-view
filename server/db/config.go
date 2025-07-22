package db

import (
	"fmt"
	"os"
)

var currentPath string

func LoadDatabase(path string) error {
	if _, err := os.Stat(path); err != nil {
		return fmt.Errorf("cannot access file: %v", err)
	}

	if db != nil {
		db.Close()
	}

	currentPath = path
	return InitDB(path)
}

func GetCurrentPath() string {
	return currentPath
}

func IsConnected() bool {
	if db != nil && currentPath != "" {
		return true
	}
	return false
}
