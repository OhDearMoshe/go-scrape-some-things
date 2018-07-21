package hostnames

import (
	"net/url"
	"log"
	"strings"
)

/**
	Extracts the hostname from a given URL
	aka https://www.facebook.com/some/path
	will return just facebook.com.

	Returns a blank string if it cannot
	determine the hostname

 */
func ExtractHostname(uri string) (string){
	u, err := url.Parse(uri)
	if err != nil {
		// Appears to just return nil if a hostname is invalid
		log.Fatal(err)
	}
	return u.Hostname()
}

func FilterUrls(urls []string, host string) []string {
	var results []string
	for _, u := range urls {
		if len(u) > 1 {
			if u[0] == '/' {
				results = append(results, u)
			}
			// Do nothing, is just a redirect to the current page
		  if strings.HasPrefix(u, "http://") || strings.HasPrefix(u, "https://") {
			  hostname := ExtractHostname(u)
			  if strings.Contains(hostname, host) {
			  	results = append(results, u)
			  }
		  }
		}
	}
	return results
}