package concurrency

import "testing"

func mockWebsiteChecker(Url string) bool {
	if Url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}

func TestCheckWebistes(t *testing.T) {

	websitesList := []string{"http://www.google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	testset := []struct {
		Name          string
		Url           string
		ShouldBeAlive bool
	}{
		{Name: "Google", Url: "http://www.google.com", ShouldBeAlive: true},
		{Name: "blog", Url: "http://blog.gypsydave5.com", ShouldBeAlive: true},
		{Name: "Invalid Url", Url: "waat://furhurterwe.geds", ShouldBeAlive: false},
	}

	t.Run("Check a list of websites", func(t *testing.T) {

		results := CheckWebsites(mockWebsiteChecker, websitesList)
		for _, test := range testset {
			got, ok := results[test.Url]
			if got != test.ShouldBeAlive && ok {
				t.Errorf("Is %v alive? Got %v, should be %v", test.Url, got, test.ShouldBeAlive)
			}
		}
	})
}
