package hostnames

import (
	"net/url"
	"log"
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