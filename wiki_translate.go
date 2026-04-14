package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

// WikiAPIResponse represents the structure of the Wikipedia API response
type WikiAPIResponse struct {
	BatchComplete string `json:"batchcomplete"`
	Query         struct {
		Pages map[string]struct {
			LangLinks []struct {
				Lang  string `json:"lang"`
				Title string `json:"*"`
			} `json:"langlinks"`
		} `json:"pages"`
	} `json:"query"`
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: wiki_translate <article_title>")
	}

	articleTitle := os.Args[1]

	// Validate and encode the article title for URL
	encodedTitle := url.PathEscape(articleTitle)
	apiURL := fmt.Sprintf("https://ru.wikipedia.org/w/api.php?format=json&action=query&prop=langlinks&titles=%s&lllimit=500", encodedTitle)

	// Create HTTP client with timeout and User-Agent header
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Make request with error handling
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("User-Agent", "WikiTranslateBot/1.0 (contact: example@example.com)")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	// Check HTTP status code
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Request failed with status: %d", resp.StatusCode)
	}

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	// Parse JSON response
	var apiResp WikiAPIResponse
	if err := json.Unmarshal(body, &apiResp); err != nil {
		log.Fatalf("Failed to parse JSON response: %v", err)
	}

	// Output language links
	found := false
	for _, page := range apiResp.Query.Pages {
		if len(page.LangLinks) == 0 {
			fmt.Println("No language links found for this article.")
			return
		}
		for _, link := range page.LangLinks {
			fmt.Printf("%s (%s)\n", link.Title, link.Lang)
			found = true
		}
		break // Only process first page
	}

	if !found {
		fmt.Println("No language links found for this article.")
	}
}
