package util

import "os"

func Getenv(key string, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}

	return value
}

func IsValidDir(path string) bool {
	fi, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}

	return fi.Mode().IsDir()
}
