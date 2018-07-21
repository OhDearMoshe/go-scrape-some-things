package scrape

import (
	"net/http"
	"log"
	"go-scrape-some-things/extract"
	"errors"
	"fmt"
	"go-scrape-some-things/hostnames"
)

type ScrapeResult struct {
	hostname string
	paths []string
	err error
}


func getHtmlFromUrl(url string) ([]byte, error) {
	log.Printf("Attempting to retrieve content from %q", url)
	result, err := http.Get(url)
	if result.StatusCode != 200 {
		message := fmt.Sprintf("Error: Unable to reach %s non 200 response code recieved: %d", url, result.StatusCode)
		log.Printf(message)
		return nil, errors.New(message)
	}
	if err != nil {
		return nil, err
	}
	return extract.FromHttpResponse(result), err

}

func contains(list []string, s string) bool {
	for _,item := range list {
		if item == s {
			return true
		}
	}
	return false
}

func AppendNoneVisit(visited []string, urls []string) []string {
	for _, u := range urls {
		if !contains(visited, u) {
			visited = append(visited, u)
		}
	}
	return visited
}

func Scrape(url string, hostname string) []ScrapeResult {
	var results []ScrapeResult
	var visited []string
	var toVisit = []string {url,}
	for len(toVisit) > 0 {
		nextUrl := toVisit[0]
		toVisit = toVisit[1:]

		html, err := getHtmlFromUrl(nextUrl)
		visited = append(visited, nextUrl)
		var urls []string
		if err == nil {
			urls = extract.UrlsFromHtml(html)
			log.Printf("Found %d adjacent pages from %q", len(urls), nextUrl)
			// Filter Urls
			urls = hostnames.FilterUrls(urls, hostname)
			// Append to next
		}

	}

	return results
}