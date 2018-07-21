package main

import (
	"os"
	"log"
	"go-scrape-some-things/scrape"
)

func getArgs() string {
	if len(os.Args) < 2 {
		log.Fatal("Error no target provided")
	}
	return os.Args[1:][0]
}



func main() {
	target := getArgs()
	log.Printf("Creating site map of %q", target)
	results := scrape.Scrape(target)
	for _, r := range results {
		log.Print(r)
	}
}
