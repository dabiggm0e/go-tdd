package concurrency

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(Url string) bool {
	if Url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}

func TestCheckWebistes(t *testing.T) {

	websites := []string{
		"http://www.google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	want := map[string]bool{
		"http://www.google.com":      true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}

	t.Run("Check a list of websites", func(t *testing.T) {

		got := CheckWebsites(mockWebsiteChecker, websites)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
