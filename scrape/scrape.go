package scrape

import (
	"net/http"
	"io/ioutil"
	"log"
	"go-scrape-some-things/extract"
)

type ScrapeResult struct {
	hostname string
	paths []string
}


func getHtmlFromUrl(url string) []byte {
	result, err := http.Get(url)
	if result.StatusCode != 200 {
		log.Fatal("Unable to reach %q non 200 response code recieved: %d", url, result.StatusCode)
	}
	if err != nil {
		log.Fatal(err)
	}
	return extract.ExtractResponse(result)

}

func Scrape(url string, hostname string) []ScrapeResult {
	var results []ScrapeResult
	html := getHtmlFromUrl(url)

	return results
}