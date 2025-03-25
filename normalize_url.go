package main

import (
	"net/url"
	"strings"
)

// normalizeURL takes a raw URL and returns a normalized version without the scheme and trailing slash
func normalizeURL(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	// Remove scheme (http/https)
	parsedURL.Scheme = ""

	// Remove "www." prefix if present
	parsedURL.Host = strings.TrimPrefix(parsedURL.Host, "www.")

	// Combine host and path
	normalized := parsedURL.Host + parsedURL.Path

	// Remove trailing slash
	normalized = strings.TrimSuffix(normalized, "/")

	return normalized, nil
}
