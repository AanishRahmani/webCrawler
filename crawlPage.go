package main

import (
	"fmt"
	"strings"
)

// Recursively crawl internal pages
func (cfg *config) crawlPage(rawCurrentURL string) {
	defer cfg.wg.Done()

	// Ensure we are within the same domain
	currentDomain, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Println("Skipping invalid URL:", rawCurrentURL)
		<-cfg.concurrencyControl // Release a slot in the concurrency buffer
		return
	}

	if !strings.HasPrefix(currentDomain, cfg.baseURL.Host) {
		fmt.Println("Skipping external URL:", rawCurrentURL)
		<-cfg.concurrencyControl
		return
	}

	// If page already exists, increment count
	if !cfg.addPageVisit(currentDomain) {
		<-cfg.concurrencyControl
		return
	}

	// Mark page as visited
	fmt.Println("Crawling:", rawCurrentURL)

	// Fetch HTML content
	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Println("Error fetching page:", err)
		<-cfg.concurrencyControl
		return
	}

	// Extract links and crawl further
	links, err := getURLsFromHTML(htmlBody, rawCurrentURL)
	if err != nil {
		fmt.Println("Error extracting links:", err)
		<-cfg.concurrencyControl
		return
	}

	for _, link := range links {
		cfg.wg.Add(1)
		cfg.concurrencyControl <- struct{}{}
		go cfg.crawlPage(link)
	}

	<-cfg.concurrencyControl // Release a slot in the concurrency buffer
}
