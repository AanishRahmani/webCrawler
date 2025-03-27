package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

// Extract and normalize URLs from HTML
func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	var urls []string

	// Trim whitespace
	htmlBody = strings.TrimSpace(htmlBody)

	// Parse the base URL
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("invalid base URL: %w", err)
	}

	// Parse the HTML content
	doc, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML: %w", err)
	}

	// Recursive function to extract links
	var extractLinks func(*html.Node)
	extractLinks = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					href := strings.TrimSpace(attr.Val)

					// Ignore empty or missing href attributes
					if href == "" {
						continue
					}

					parsedURL, err := url.Parse(href)
					if err == nil {
						resolvedURL := baseURL.ResolveReference(parsedURL).String()
						urls = append(urls, resolvedURL)
					}
				}
			}
		}
		// Recursively check child nodes
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			extractLinks(c)
		}
	}

	// Start extracting links
	extractLinks(doc)

	return urls, nil
}

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
