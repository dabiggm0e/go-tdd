//https://quii.gitbook.io/learn-go-with-tests/build-an-application/http-server
package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	score    map[string]int
	winCalls []string
}

///////////////////////////
//// Unit Tests
///////////////////////////
func TestGETPlayers(t *testing.T) {
	t.Run("Getting Mo's score", func(t *testing.T) {
		request := newGetScoreRequest("Mo")
		response := httptest.NewRecorder()
		store := initPlayersStore()
		playerServer := &PlayerServer{store: store}
		playerServer.ServeHTTP(response, request)

		assertStatusCode(t, response.Code, http.StatusOK)
		assertResponseReply(t, response.Body.String(), "20")
	})

	t.Run("Return Ziggy's score", func(t *testing.T) {
		request := newGetScoreRequest("Ziggy")
		response := httptest.NewRecorder()

		store := initPlayersStore()
		playerServer := &PlayerServer{store: store}
		playerServer.ServeHTTP(response, request)

		assertStatusCode(t, response.Code, http.StatusOK)
		assertResponseReply(t, response.Body.String(), "10")
	})

	t.Run("Return 404 on not found player", func(t *testing.T) {
		request := newGetScoreRequest("JOHNDOE")
		response := httptest.NewRecorder()

		store := initPlayersStore()
		playerServer := &PlayerServer{store: store}

		playerServer.ServeHTTP(response, request)

		want := http.StatusNotFound
		got := response.Code

		assertStatusCode(t, got, want)
	})

}

func TestStoreWins(t *testing.T) {

	t.Run("It records a win when POST", func(t *testing.T) {
		store := initPlayersStore()
		player := "JOHNDOE"
		request := newPostScoreRequest(player)
		response := httptest.NewRecorder()

		playerServer := &PlayerServer{store: store}
		playerServer.ServeHTTP(response, request)

		assertStatusCode(t, response.Code, http.StatusAccepted)

		if len(store.winCalls) != 1 {
			t.Fatalf("Expecting %d calls to RecordWin, got %d", 1, len(store.winCalls))
		}

		if store.winCalls[0] != player {
			t.Errorf("did not store the correct winner. Got %s, want %s", store.winCalls[0], player)
		}
	})
}

///////////////////////////
////// Integration Tests
///////////////////////////

func TestRecordWinsAndRetrieveScore(t *testing.T) {
	store := &InMemoryPlayerStore{}

	server := &PlayerServer{store: store}
	player := "Luffy"

	server.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetScoreRequest(player))

	assertStatusCode(t, response.Code, http.StatusOK)
	assertResponseReply(t, response.Body.String(), "3")

}

////////////
/// Assertions helper functions
///////////
func assertResponseReply(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Response body is wrong. Got '%s' want '%s'", got, want)
	}
}

func assertStatusCode(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("Got response code %d, want %d", got, want)
	}
}

//////////////
/// helper functions
///////////////
func newGetScoreRequest(player string) *http.Request {
	path := fmt.Sprintf("/players/%s", player)
	request, _ := http.NewRequest(http.MethodGet, path, nil)
	return request
}

func newPostScoreRequest(player string) *http.Request {
	path := fmt.Sprintf("/players/%s", player)
	request, _ := http.NewRequest(http.MethodPost, path, nil)
	return request
}

//////////////////
//// stub implementation
//////////////////
func (s *StubPlayerStore) GetPlayerScore(name string) (int, error) {

	if score, ok := s.score[name]; ok {
		return score, nil
	}
	return 0, ERRPLAYERNOTFOUND
}

func (s *StubPlayerStore) RecordWin(name string) {
	//s.score[name]++
	//return s.score[name], nil
	s.winCalls = append(s.winCalls, name)
}

func initPlayersStore() *StubPlayerStore {
	store := StubPlayerStore{
		score: map[string]int{
			"Mo":    20,
			"Ziggy": 10,
		},
	}

	return &store
}
