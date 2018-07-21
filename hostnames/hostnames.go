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


/**
	If url is just a path appends the baseUrl to it otherwise returns as normal
	This is to help us follow links that are relative
 */
func SanatizeUrl(potentialPath string, baseUrl string) string {
	if potentialPath[0] == '/' {
		return baseUrl + potentialPath
	}
	return potentialPath
}

/**
   Filters out any urls that are not owned by the current host
 */
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