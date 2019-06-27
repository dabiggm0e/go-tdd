//https://quii.gitbook.io/learn-go-with-tests/build-an-application/http-server
package main

import (
	"errors"
	"fmt"
	"net/http"
)

type PlayerStore interface {
	GetPlayerScore(name string) (int, error)
	RecordPlayerScore(name string) (int, error)
}

type PlayerServer struct {
	store PlayerStore
}

var (
	ERRPLAYERNOTFOUND = errors.New("Player not found")
)

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var score int
	var err error

	player := r.URL.Path[len("/players/"):] // extract player name from the GET Path

	switch r.Method {

	case "GET":
		score, err = p.store.GetPlayerScore(player)
		if err == ERRPLAYERNOTFOUND {
			w.WriteHeader(http.StatusNotFound)
		}

	case "POST":
		score, err = p.store.RecordPlayerScore(player)
		w.WriteHeader(http.StatusAccepted)
	}

	if IsVerbose() {
		fmt.Printf("player: %s. score: %d. path: %s. err: %v\n", player, score, r.URL.Path, err)
	}

	fmt.Fprintf(w, "%d", score)

}

/*func (p *PlayerServer) GetPlayerScore(name string) string {
	switch name {
	case "Mo":
		return "20"
	case "Ziggy":
		return "10"
	default:
		return ""
	}
}*/
