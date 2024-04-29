package main

import (
	"testing"
)

func TestIsValidURL(t *testing.T) {
	allowedDomains := []string{"en.wikipedia.org"}

	tests := []struct {
		url   string
		valid bool
	}{
		{"https://en.wikipedia.org/wiki/Robotics", true},
		{"https://google.com", false},
		{"ftp://en.wikipedia.org/wiki/Robot", false},
		{"https://en.wikipedia.org", true},
		{"https://wikipedia.org", false},
	}

	for _, test := range tests {
		if isValidURL(test.url, allowedDomains) != test.valid {
			t.Errorf("isValidURL (%s) = %t, want %t", test.url, !test.valid, test.valid)
		}
	}
}
