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
	Hostname string
	Paths    []string
	Err      error
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
	for _, item := range list {
		if item == s {
			return true
		}
	}
	return false
}

/**
	Returns a list of all the results that are not present in the visited parameter
 */
func GetNoneVisited(visited []string, urls []string) []string {
	var toVisit []string
	for _, u := range urls {
		if !contains(visited, u) {
			toVisit = append(toVisit, u)
		}
	}
	return toVisit
}
/**
 	Goes to the given URL and visits the page, returns a ScrapeResult with
	any child pages it has found and any errors it has encountered
 */
func FetchPage(toVisit string, baseUrl string, hostname string) ScrapeResult {
	url := hostnames.SanatizeUrl(toVisit, baseUrl)
	html, err := getHtmlFromUrl(url)
	var urls []string
	if err == nil {
		urls = extract.UrlsFromHtml(html)
		urls = hostnames.FilterUrls(urls, hostname)
		log.Printf("Found %d adjacent pages from %q", len(urls), toVisit)
	}

	return ScrapeResult{toVisit, urls, err}
}

func Scrape(url string) []ScrapeResult {
	hostname := hostnames.ExtractHostname(url)
	var results []ScrapeResult
	var visited []string
	var toVisit = []string{url}
	for len(toVisit) > 0 {
		// Pop the next result
		nextUrl := toVisit[0]
		toVisit = toVisit[1:]
		if !contains(visited, nextUrl) {
			result := FetchPage(nextUrl, url, hostname)
			results = append(results, result)
			visited = append(visited, nextUrl)

			toVisit = append(toVisit, GetNoneVisited(visited, result.Paths)...)
		}
	}

	return results
}
