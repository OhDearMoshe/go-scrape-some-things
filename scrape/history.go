package scrape


/**
	Abstraction of all the places we have visited so far
 */
type History struct {
	visitedUrls []string
}

/**
	Returns true if the url has been visited yet. Else false
 */
func (h *History) HaveNotYetVisited(url string) bool {
	return !contains(h.visitedUrls, url)
}


/**
	Marks a url as having being visited
 */
func (h *History) MarkAsVisited(url string) {
	h.visitedUrls = append(h.visitedUrls, url)
}

/**
	Retrieves and array of all the urls not visited from the
	input array
 */
func (h *History) GetNonVisitedUrls(urls[] string) []string {
	var toVisit []string
	for _, u := range urls {
		if !contains(h.visitedUrls, u) {
			toVisit = append(toVisit, u)
		}
	}
	return toVisit
}

func contains(list []string, s string) bool {
	for _, item := range list {
		if item == s {
			return true
		}
	}
	return false
}