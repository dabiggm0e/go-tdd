//https://quii.gitbook.io/learn-go-with-tests/build-an-application/http-server
package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	score map[string]int
}

func TestGETPlayers(t *testing.T) {
	t.Run("Getting Mo's score", func(t *testing.T) {
		request, _ := newGetScoreRequest("Mo")
		response := httptest.NewRecorder()
		store := &StubPlayerStore{}
		playerServer := &PlayerServer{store: store}
		playerServer.ServeHTTP(response, request)

		assertResponseReply(t, response.Body.String(), "20")
	})

	t.Run("Return Ziggy's score", func(t *testing.T) {
		request, _ := newGetScoreRequest("Ziggy")
		response := httptest.NewRecorder()

		store := &StubPlayerStore{}
		playerServer := &PlayerServer{store: store}
		playerServer.ServeHTTP(response, request)

		assertResponseReply(t, response.Body.String(), "10")
	})

	t.Run("Call index / returns empty response", func(t *testing.T) {
		request, _ := newGetScoreRequest("")
		response := httptest.NewRecorder()

		store := &StubPlayerStore{}
		playerServer := &PlayerServer{store: store}
		playerServer.ServeHTTP(response, request)

		assertResponseReply(t, response.Body.String(), "")
	})
}

func assertResponseReply(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Response body is wrong. Got '%s' want '%s'", got, want)
	}
}

func newGetScoreRequest(player string) (*http.Request, error) {
	path := fmt.Sprintf("/players/%s", player)
	return http.NewRequest(http.MethodGet, path, nil)
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	return 0
}
