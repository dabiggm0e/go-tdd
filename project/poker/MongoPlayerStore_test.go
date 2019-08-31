package poker

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMongoPlayerStore(t *testing.T) {

	t.Run("Test connection to MongoDB", func(t *testing.T) {
		store, err := NewMongoPlayerStore("test")
		defer store.Teardown()

		if err != nil {
			t.Fatalf("Error connecting to the DB: %v", err.Error())
		}
	})

	t.Run("GET unknown user return 404", func(t *testing.T) {
		player := "AGAGAGA"
		request := newGetScoreRequest(player)

		store, _ := NewMongoPlayerStore("test")
		defer store.Teardown()
		store.deleteCollection()

		server := NewPlayerServer(store)

		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		assertStatusCode(t, response.Code, http.StatusNotFound)
	})

	t.Run("Record win for new player returns Status 201", func(t *testing.T) {
		player := "Mo"

		store, _ := NewMongoPlayerStore("test")
		defer store.Teardown()
		store.deleteCollection()

		request := newPostScoreRequest(player)
		response := httptest.NewRecorder()

		server := NewPlayerServer(store)
		server.ServeHTTP(response, request)

		got := response.Code
		want := http.StatusAccepted
		assertStatusCode(t, got, want)

		request = newGetScoreRequest(player)
		server.ServeHTTP(response, request)

		log.Printf("Player %s score is %v", player, response.Body)
	})

	t.Run("Record win for existing player returns Status 201", func(t *testing.T) {
		player := "Mo"

		store, _ := NewMongoPlayerStore("test")
		defer store.Teardown()
		store.deleteCollection()

		request := newPostScoreRequest(player)
		response := httptest.NewRecorder()

		server := NewPlayerServer(store)
		server.ServeHTTP(response, request)
		server.ServeHTTP(response, request)
		server.ServeHTTP(response, request)
		server.ServeHTTP(response, request)
		server.ServeHTTP(response, request)

		got := response.Code
		want := http.StatusAccepted
		assertStatusCode(t, got, want)

		request = newGetScoreRequest(player)
		response = httptest.NewRecorder()
		server.ServeHTTP(response, request)

		assertStatusCode(t, response.Code, http.StatusOK)
		assertResponseReply(t, response.Body.String(), "5")

	})

	t.Run("GET /league", func(t *testing.T) {
		store, _ := NewMongoPlayerStore("test")
		defer store.Teardown()
		store.deleteCollection()

		player := "Mo"
		request := newPostScoreRequest(player)
		response := httptest.NewRecorder()

		server := NewPlayerServer(store)
		server.ServeHTTP(response, request)
		server.ServeHTTP(response, request)
		server.ServeHTTP(response, request)

		request = newGetLeagueRequest()
		response = httptest.NewRecorder()
		server.ServeHTTP(response, request)

		assertStatusCode(t, response.Code, http.StatusOK)
		assertResponseContentType(t, response, jsonContentType)

		gotLeague := getLeagueFromResponse(t, response.Body)
		wantedLeague := League{
			{"Mo", 3},
		}

		assertLeague(t, gotLeague, wantedLeague)
	})

	t.Run("GET /league sorted", func(t *testing.T) {
		store, _ := NewMongoPlayerStore("test")
		defer store.Teardown()
		store.deleteCollection()

		player := "Mo"
		request := newPostScoreRequest(player)
		response := httptest.NewRecorder()

		server := NewPlayerServer(store)
		server.ServeHTTP(response, request)
		server.ServeHTTP(response, request)

		player = "Ziggy"
		request = newPostScoreRequest(player)
		server.ServeHTTP(response, request)
		server.ServeHTTP(response, request)
		server.ServeHTTP(response, request)

		request = newGetLeagueRequest()
		response = httptest.NewRecorder()
		server.ServeHTTP(response, request)

		assertStatusCode(t, response.Code, http.StatusOK)
		assertResponseContentType(t, response, jsonContentType)

		gotLeague := getLeagueFromResponse(t, response.Body)
		wantedLeague := League{
			{"Ziggy", 3},
			{"Mo", 2},
		}

		assertLeague(t, gotLeague, wantedLeague)
	})
}
