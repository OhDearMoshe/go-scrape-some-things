package scrape

import "testing"

func TestHistory_HaveNotYetVisited(t *testing.T) {
	history := History{[]string{"http://monzo.com"}}

	if history.HaveNotYetVisited("http://monzo.com") == true {
		t.Error("Methods returns true for visiting a location it has marked as visited")
	}

	if history.HaveNotYetVisited("http://notVisitHere.com") == false {
		t.Error("Methods returns false for visiting a location it has marked as not visited")
	}
}

func TestHistory_MarkAsVisited(t *testing.T) {
	history := History{}
	history.MarkAsVisited("http://monzo.com")
	if history.HaveNotYetVisited("http://monzo.com") == true {
		t.Error("Methods returns true for visiting a location it has marked as visited")
	}
}

func TestHistory_GetNonVisitedUrls(t *testing.T) {
	expected := []string{
		"http://blog.monzo.com",
		"/latest_news",
	}
	existing := []string{
		"http://monzo.com",
		"/faq",
	}
	input := []string{
		"http://monzo.com",
		"http://blog.monzo.com",
		"/latest_news",
	}

	history := History{existing}
	results := history.GetNonVisitedUrls(input)

	if len(results) != len(expected) {
		t.Fatalf("Error expecting results of length %d but got %d", len(expected), len(results))
	}

	for index, r := range results {
		if r != expected[index] {
			t.Fatalf("Expected %q but was %q", expected[index], r)
		}
	}
}