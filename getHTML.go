package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	resp, err := http.Get(rawURL)
	if err != nil {
		return "", fmt.Errorf("can't fetch HTTP response: %v", err)
	}
	defer resp.Body.Close()

	// Handle HTTP errors (status codes 400+)
	if resp.StatusCode >= 400 {
		return "", fmt.Errorf("received error response: %d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	// Ensure Content-Type is text/html
	contentType := resp.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "text/html") {
		return "", fmt.Errorf("invalid content type: %s", contentType)
	}

	// Read the HTML response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	return string(body), nil
}

// Ensure only unique pages are crawled
