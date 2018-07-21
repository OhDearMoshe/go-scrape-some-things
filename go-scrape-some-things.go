package main

import (
	"os"
	"log"
	"go-scrape-some-things/scrape"
	"encoding/json"
	"io/ioutil"
)

func getArgs() string {
	if len(os.Args) < 2 {
		log.Fatal("Error no target provided")
	}
	return os.Args[1:][0]
}


func toFile(results []scrape.ScrapeResult) {
	bytes, err := json.Marshal(results)
	if err != nil {
		log.Fatal("Error writing results")
	}
	ioutil.WriteFile("/tmp/results.json", bytes, 0644)
}
func main() {
	target := getArgs()
	log.Printf("Creating site map of %q", target)
	results := scrape.Scrape(target)
	toFile(results)


}
