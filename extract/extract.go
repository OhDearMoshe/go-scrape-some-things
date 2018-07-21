package extract

import (
	"net/http"
	"io/ioutil"
	"log"
	"golang.org/x/net/html"
	"bytes"
)

/**
 	Extracts bytes from a http.Response
 */
func ExtractResponse(response *http.Response)  []byte{
	bytes, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	return bytes

}


/**
	Extracts all the href's from A tags of a html page and
	returns them back as a list
 */
func ExtractUrlsFromHtml(page []byte) []string {
	var links []string
	doc, err := html.Parse(bytes.NewReader(page))
	if err != nil {
		log.Fatal(err)
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					links = append(links, a.Val)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return links
}