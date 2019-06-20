package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

func TestRacer(t *testing.T) {
	slowServer := makeDelayedServer(20 * time.Millisecond)
	defer slowServer.Close()
	fastServer := makeDelayedServer(0 * time.Millisecond)
	defer fastServer.Close()

	slowUrl := slowServer.URL
	fastUrl := fastServer.URL

	want := fastUrl
	got, _ := Racer(slowUrl, fastUrl)

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}