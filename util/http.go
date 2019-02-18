package util

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func GetUrl(url string) (io.Reader, error) {
	response, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GET status: %s", response.Status)
	}

	return response.Body, nil
}

func IsValidUrl(content string) bool {
	_, err := url.ParseRequestURI(content)

	if err != nil {
		return false
	}

	return true
}
