package main

import (
	"fmt"
	"net/url"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <website-url>")
		os.Exit(1)
	}

	rawBaseURL := os.Args[1]
	parsedBaseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Println("Invalid URL:", rawBaseURL)
		os.Exit(1)
	}

	cfg := &config{
		pages:              make(map[string]int),
		baseURL:            parsedBaseURL,
		concurrencyControl: make(chan struct{}, 5),
	}

	cfg.wg.Add(1)
	cfg.concurrencyControl <- struct{}{}
	go cfg.crawlPage(rawBaseURL)

	cfg.wg.Wait()

	fmt.Println("\nCrawled Pages:")
	for url, count := range cfg.pages {
		fmt.Printf("%s - %d times\n", url, count)
	}
}
