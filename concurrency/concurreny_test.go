package concurrency

import (
	"reflect"
	"testing"
	"time"
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

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(time.Millisecond * 40)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)

	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
