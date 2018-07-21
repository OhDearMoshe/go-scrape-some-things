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