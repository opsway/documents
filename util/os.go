package util

import "os"

// Getenv returns the environment variables specified key in this machine
func Getenv(key string, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}

	return value
}

// IsValidDir indicates existence of directory path in this machine
func IsValidDir(path string) bool {
	fi, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}

	return fi.Mode().IsDir()
}
