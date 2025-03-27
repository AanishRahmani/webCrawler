package main

import (
	"net/url"
	"strings"
)

func normalizeURL(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	// Keep the scheme (http/https) for making requests
	scheme := parsedURL.Scheme
	if scheme == "" {
		scheme = "https" // Default to HTTPS if missing
	}

	// Remove "www." prefix for consistency
	parsedURL.Host = strings.TrimPrefix(parsedURL.Host, "www.")

	// Combine host and path
	normalized := parsedURL.Host + parsedURL.Path

	// Remove trailing slash
	normalized = strings.TrimSuffix(normalized, "/")

	// Return normalized URL for tracking and full URL for requests
	return normalized, nil
}
