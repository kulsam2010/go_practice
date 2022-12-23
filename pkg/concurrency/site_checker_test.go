package concurrency

import (
	"reflect"
	"testing"
)

func TestCheckWebsites(t *testing.T) {
	websites := []string{"http://www.this_website_returns_true.com",
		"www.google.com",
		"wwww.facebook.com",
		"http://www.this_website_returns_false.com"}

	want := map[string]bool{
		"http://www.this_website_returns_true.com": true,
		"www.google.com":                            true,
		"wwww.facebook.com":                         true,
		"http://www.this_website_returns_false.com": false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func mockWebsiteChecker(url string) bool {
	return url != "http://www.this_website_returns_false.com"
}
