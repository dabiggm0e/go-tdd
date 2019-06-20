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
	slowServer := makeDelayedServer(20 * time.Milliecond)
	defer slowServer.Close()
	fastServer := makeDelayedServer(0 * time.Milliecond)
	defer fastServer.Close()

	slowUrl := slowServer.URL
	fastUrl := fastServer.URL

	want := fastUrl
	got, err := Racer(slowUrl, fastUrl)

	assertError(t, err)
	assertResults(t, got, want)

}

func assertResults(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func assertError(t *testing.T, err error) {
	t.Helper()

	switch err {
	case ERRTIMEOUT:
		t.Fatalf("Error: %v", err.Error())
	}
}
