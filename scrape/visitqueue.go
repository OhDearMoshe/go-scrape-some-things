package scrape


/**
 	Abstraction of getting the next queue to visit
 */
type VisitQueue struct {
	toVisit []string
}

/**
	Returns true if there are more places to visit
 */
func (v *VisitQueue) HasNext() bool {
	return len(v.toVisit) > 0
}

/**
	Pop's the next element to visit
 */
func (v *VisitQueue) GetNextToVisit() string {
	toReturn := (v.toVisit)[0]
	v.toVisit = (v.toVisit)[1:]
	return toReturn
}

/**
	Adds more to the queue to visit
 */
func (v *VisitQueue) AddMoreToVisit(urls []string) {
	v.toVisit = append(v.toVisit, urls...)
}
