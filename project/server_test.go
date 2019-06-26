//https://quii.gitbook.io/learn-go-with-tests/build-an-application/http-server
package main

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
		store := initPlayersStore()
		playerServer := &PlayerServer{store: store}
		playerServer.ServeHTTP(response, request)

		assertResponseReply(t, response.Body.String(), "20")
	})

	t.Run("Return Ziggy's score", func(t *testing.T) {
		request, _ := newGetScoreRequest("Ziggy")
		response := httptest.NewRecorder()

		store := initPlayersStore()
		playerServer := &PlayerServer{store: store}
		playerServer.ServeHTTP(response, request)

		assertResponseReply(t, response.Body.String(), "10")
	})

	t.Run("Return 404 on not found player", func(t *testing.T) {
		request, _ := newGetScoreRequest("JOHNDOE")
		response := httptest.NewRecorder()

		store := initPlayersStore()
		playerServer := &PlayerServer{store: store}

		playerServer.ServeHTTP(response, request)

		want := http.StatusNotFound
		got := response.Code

		if got != want {
			t.Errorf("Got HTTP Status %d, want %d ", got, want)
		}
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

func (s *StubPlayerStore) GetPlayerScore(name string) (int, error) {

	if score, ok := s.score[name]; ok {
		return score, nil
	}
	return 0, ERRPLAYERNOTFOUND
}

func initPlayersStore() *StubPlayerStore {
	store := StubPlayerStore{
		map[string]int{
			"Mo":    20,
			"Ziggy": 10,
		},
	}

	return &store
}
