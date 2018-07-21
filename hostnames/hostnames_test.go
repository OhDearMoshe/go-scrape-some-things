package hostnames

import "testing"

func TestExtractHostname(t *testing.T) {
	cases := []struct {
		in, want string
	} {
		{"http://monzo.com", "monzo.com"},
		{"https://monzo.com", "monzo.com"},
		{"http://monzo.co.uk", "monzo.co.uk"},
		{"http://blog.monzo.com", "blog.monzo.com"},
		{"https://facebook.com/some/path", "facebook.com"},
		{"i2901uousao", ""},
	}

	for _, c := range cases {
		got := ExtractHostname(c.in)
		if got != c.want {
			t.Errorf("ExtractHostname(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestFilterUrls(t *testing.T) {
	input := []string {
		"/",
		"/about",
		"/community",
		"/faq",
		"https://itunes.apple.com/gb/app/mondo/id1052238659",
		"/-play-store-redirect", // This is a tricky one, Technically this is a valid subdomain
		"tel:08008021281",
		"mailto:help@monzo.com",
		"https://facebook.com/monzo",
		"https://blog.monzo.com",
	}
	expected := []string {
		"/about",
		"/community",
		"/faq",
		"/-play-store-redirect",
		"https://blog.monzo.com",
	}
	results := FilterUrls(input, "monzo.com")
	if len(results) != len(expected) {
		t.Errorf("Invalid length was %d expected %d",len(results), len(expected) )
	}
	for index, e := range expected {
		if e != results[index] {
			t.Errorf("Unexpected result. Expected: %q was %q", e, results[index])
		}
	}
}