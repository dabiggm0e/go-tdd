//https://quii.gitbook.io/learn-go-with-tests/build-an-application/http-server
package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	t.Run("Getting Mo's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Mo", nil)
		response := httptest.NewRecorder()

		PlayerServer(response, request)
		got := response.Body.String()

		want := "20"

		assertScore(t, got, want)
	})

	t.Run("Return Ziggy's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/Players/Ziggy", nil)
		response := httptest.NewRecorder()
		want := "10"

		PlayerServer(response, request)

		got := response.Body.String()
		assertScore(t, got, want)
	})

	t.Run("Call index / returns empty response", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		want := ""
		PlayerServer(response, request)
		got := response.Body.String()
		assertScore(t, got, want)
	})
}

func assertScore(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got '%s' want '%s'", got, want)
	}
}
