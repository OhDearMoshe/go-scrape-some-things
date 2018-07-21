package scrape

import "testing"

func TestVisitQueue_HasNext(t *testing.T) {

	queue := VisitQueue{}

	if queue.HasNext() != false {
		t.Fatal("Empty list return true to having an element")
	}
	input := []string{"http://monzo.com"}
	queue.AddMoreToVisit(input)

	if queue.HasNext() != true {
		t.Fatal("List with an element returns flase to having elements")
	}
}

func TestVisitQueue_GetNextToVisit(t *testing.T) {
	queue := VisitQueue{}
	input := []string{"http://monzo.com"}
	queue.AddMoreToVisit(input)

	result := queue.GetNextToVisit()
	if result != input[0] {
		t.Errorf("Error, was expecting queue to return %q but was %q", input[0], result)
	}

	if queue.HasNext() != false {
		t.Fatal("Empty list return true to having an element")
	}
}