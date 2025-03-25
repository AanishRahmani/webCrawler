package main

import (
	"reflect"
	"testing"
)

func TestExtractAndNormalizeLinks(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `<html>
		<body>
			<a href="/path/one">Boot.dev</a>
			<a href="https://other.com/path/one">Boot.dev</a>
		</body>
	</html>`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		{
			name:     "multiple links",
			inputURL: "https://example.com",
			inputBody: `<html>
		<body>
			<a href="/home">Home</a>
			<a href="/about">About</a>
			<a href="https://other.com/contact">Contact</a>
		</body>
	</html>`,
			expected: []string{"https://example.com/home", "https://example.com/about", "https://other.com/contact"},
		},
		{
			name:     "handle missing or invalid hrefs",
			inputURL: "https://example.com",
			inputBody: `<html>
		<body>
			<a>No href</a>
			<a href="">Empty</a>
			<a href="https://valid.com/page">Valid</a>
		</body>
	</html>`,
			expected: []string{"https://valid.com/page"},
		},
	}
	t.Log("Running TestExtractAndNormalizeLinks...") // Extra log line
	for i, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := getURLsFromHTML(test.inputBody, test.inputURL)
			if err != nil {
				t.Errorf("Test %d - '%s' FAIL: unexpected error: %v", i, test.name, err)
			}
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Test %d - '%s' FAIL: expected %v, got %v", i, test.name, test.expected, result)
			}
		})
	}
	t.Log("TestExtractAndNormalizeLinks done...")
}
