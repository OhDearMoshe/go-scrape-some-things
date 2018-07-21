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

/**
 	Goes to the given URL and visits the page, returns a ScrapeResult with
	any child pages it has found and any errors it has encountered
 */
func FetchResult(toVisit string, baseUrl string, hostname string) ScrapeResult {
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
	history := History{}
	toVisit := []string{url}

	visitQueue := VisitQueue{toVisit}

	for visitQueue.HasNext() {
		nextUrl := visitQueue.GetNextToVisit()
		if history.HaveNotYetVisited(nextUrl) {
			result := FetchResult(nextUrl, url, hostname)
			results = append(results, result)
			history.MarkAsVisited(nextUrl)

			visitQueue.AddMoreToVisit(history.GetNonVisitedUrls(result.Paths))
		}
	}

	return results
}
