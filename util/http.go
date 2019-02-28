package util

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// GetURL returns response body of url
func GetURL(url string) (io.Reader, error) {
	response, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GET status: %s", response.Status)
	}

	return response.Body, nil
}

// IsValidURL indicate validate raw url
func IsValidURL(content string) bool {
	_, err := url.ParseRequestURI(content)

	if err != nil {
		return false
	}

	return true
}
