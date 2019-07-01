//https://quii.gitbook.io/learn-go-with-tests/build-an-application/http-server
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

const jsonContentType = "application/json"

type StubPlayerStore struct {
	score    map[string]int
	winCalls []string
	league   []Player
}

///////////////////////////
//// Unit Tests
///////////////////////////
func TestGETPlayers(t *testing.T) {
	t.Run("Getting Mo's score", func(t *testing.T) {
		request := newGetScoreRequest("Mo")
		response := httptest.NewRecorder()
		store := initPlayersStore()
		playerServer := NewPlayerServer(store)
		playerServer.ServeHTTP(response, request)

		assertStatusCode(t, response.Code, http.StatusOK)
		assertResponseReply(t, response.Body.String(), "20")
	})

	t.Run("Return Ziggy's score", func(t *testing.T) {
		request := newGetScoreRequest("Ziggy")
		response := httptest.NewRecorder()

		store := initPlayersStore()
		playerServer := NewPlayerServer(store)
		playerServer.ServeHTTP(response, request)

		assertStatusCode(t, response.Code, http.StatusOK)
		assertResponseReply(t, response.Body.String(), "10")
	})

	t.Run("Return 404 on not found player", func(t *testing.T) {
		request := newGetScoreRequest("JOHNDOE")
		response := httptest.NewRecorder()

		store := initPlayersStore()
		playerServer := NewPlayerServer(store)

		playerServer.ServeHTTP(response, request)

		want := http.StatusNotFound
		got := response.Code

		assertStatusCode(t, got, want)
	})

}

func TestPostgresGetPlayer(t *testing.T) {
	t.Run("PostgresPlayerStore: Return 404 on not found player", func(t *testing.T) {
		request := newGetScoreRequest("JOHNDOE")
		response := httptest.NewRecorder()

		store := NewPostgresPlayerStore()
		defer store.Teardown()
		clearPostgresStore(t, store)

		playerServer := NewPlayerServer(store)
		playerServer.ServeHTTP(response, request)

		want := http.StatusNotFound
		got := response.Code

		assertStatusCode(t, got, want)
	})

	t.Run("Getting Mo's score", func(t *testing.T) {
		player := "Mo"
		response := httptest.NewRecorder()

		store := NewPostgresPlayerStore()
		defer store.Teardown()
		clearPostgresStore(t, store)

		playerServer := NewPlayerServer(store)
		playerServer.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(player))
		playerServer.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(player))
		playerServer.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(player))

		playerServer.ServeHTTP(response, newGetScoreRequest(player))

		assertStatusCode(t, response.Code, http.StatusOK)
		assertResponseReply(t, response.Body.String(), "3")
	})

}

func TestStoreWins(t *testing.T) {

	t.Run("It records a win when POST", func(t *testing.T) {
		store := initPlayersStore()
		player := "JOHNDOE"
		request := newPostScoreRequest(player)
		response := httptest.NewRecorder()

		playerServer := NewPlayerServer(store)
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

func TestLeague(t *testing.T) {

	t.Run("It returns 200 on GET /league", func(t *testing.T) {
		store := initPlayersStore()
		server := NewPlayerServer(store)

		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetLeagueRequest())

		assertStatusCode(t, response.Code, http.StatusOK)
	})

	/*	t.Run("Return JSON scores on successful GET /league", func(t *testing.T) {
		store := StubPlayerStore{}
		server := NewPlayerServer(&store)

		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetLeagueRequest())

		var got []Player

		err := json.NewDecoder(response.Body).Decode(&got)

		assertStatusCode(t, response.Code, http.StatusOK)

		if err != nil {
			t.Fatalf("Unable to parse response from server '%s' into slice of Players, '%v'", response.Body, err)
		}

	})*/

	t.Run("StubPlayerStore: Return 404 for empty /league response json parsing", func(t *testing.T) {
		store := &StubPlayerStore{}
		server := NewPlayerServer(store)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, newGetLeagueRequest())

		assertStatusCode(t, response.Code, http.StatusNotFound)
	})

	t.Run("PostgresPlayerStore: Return 404 for empty /league response json parsing", func(t *testing.T) {
		store := NewPostgresPlayerStore()
		server := NewPlayerServer(store)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, newGetLeagueRequest())

		assertStatusCode(t, response.Code, http.StatusNotFound)
	})

	t.Run("Test league table returning correct data in json", func(t *testing.T) {
		wantedLeague := []Player{
			{"Mo", 10},
			{"Ziggy", 7},
			{"Moon", 3},
		}

		store := StubPlayerStore{nil, nil, wantedLeague}
		server := NewPlayerServer(&store)

		request := newGetLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		var gotLeague []Player
		gotLeague = getLeagueFromResponse(t, response.Body)
		//err := json.NewDecoder(response.Body).Decode(&gotLeague)

		assertResponseContentType(t, response, jsonContentType)
		assertLeague(t, gotLeague, wantedLeague)

	})
}

func TestPostgresStoreWin(t *testing.T) {
	player := "Ziggy"

	store := NewPostgresPlayerStore()
	defer store.Teardown()
	clearPostgresStore(t, store)

	playerServer := NewPlayerServer(store)

	response := httptest.NewRecorder()
	playerServer.ServeHTTP(response, newPostScoreRequest(player))

	assertStatusCode(t, response.Code, http.StatusAccepted)
}

///////////////////////////
////// Integration Tests
///////////////////////////

func TestInMemoryStoreRecordWinsAndRetrieveScore(t *testing.T) {
	store := NewInMemoryPlayerStore()

	server := NewPlayerServer(store)
	player := "Luffy"

	server.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetScoreRequest(player))

	assertStatusCode(t, response.Code, http.StatusOK)
	assertResponseReply(t, response.Body.String(), "3")

}

func TestPostgresStoreRecordWinsAndRetrieveScore(t *testing.T) {
	store := NewPostgresPlayerStore()
	defer store.Teardown()
	clearPostgresStore(t, store)
	server := NewPlayerServer(store)
	player := "Mo"

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

func newGetLeagueRequest() *http.Request {
	path := fmt.Sprintf("/league")
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

func (s *StubPlayerStore) RecordWin(name string) error {
	//s.score[name]++
	//return s.score[name], nil
	s.winCalls = append(s.winCalls, name)
	return nil
}

func initPlayersStore() *StubPlayerStore {
	store := StubPlayerStore{
		score: map[string]int{
			"Mo":    20,
			"Ziggy": 10,
		},
		league: []Player{
			{"Mo", 20},
			{"Ziggy", 10},
		},
	}

	return &store
}

func clearPostgresStore(t *testing.T, p *PostgresPlayerStore) {
	truncateSql := "TRUNCATE scores; TRUNCATE players"
	_, err := p.store.Exec(truncateSql)
	if err != nil {
		t.Fatalf("Unable to truncate the store. Err: %v", err)
	}
}

func (s *StubPlayerStore) GetLeague() []Player {
	return s.league
}

func getLeagueFromResponse(t *testing.T, body io.Reader) []Player {
	t.Helper()
	var league []Player
	err := json.NewDecoder(body).Decode(&league)

	if err != nil {
		t.Fatalf("Couldn't parse json response '%s' from server. '%v'", body, err)
		return nil
	}

	return league
}

func assertLeague(t *testing.T, gotLeague, wantedLeague []Player) {
	if !reflect.DeepEqual(gotLeague, wantedLeague) {
		t.Errorf("Got %v want %v", gotLeague, wantedLeague)
	}
}

func assertResponseContentType(t *testing.T, w *httptest.ResponseRecorder, want string) {
	t.Helper()
	got := w.Result().Header.Get("content-type")
	if got != want {
		t.Errorf("response did not have content-type of application/json. Got %v", got)
	}
}
