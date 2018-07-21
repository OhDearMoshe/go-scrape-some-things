package scrape

import "testing"

func TestAppendNoneVisit(t *testing.T) {
	expected := []string{
		"http://monzo.com",
		"/faq",
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

	results := GetNoneVisited(existing, input)

	if len(results) != len(expected) {
		t.Fatalf("Error expecting results of length %d but got %d", len(expected), len(results))
	}

	for index, r := range results {
		if r != expected[index] {
			t.Fatalf("Expected %q but was %q", expected[index], r)
		}
	}
}
