//https://quii.gitbook.io/learn-go-with-tests/build-an-application/http-server
package main

import (
	"errors"
	"fmt"
	"net/http"
)

type PlayerStore interface {
	GetPlayerScore(name string) (int, error)
}

type PlayerServer struct {
	store PlayerStore
}

var (
	ERRPLAYERNOTFOUND = errors.New("Player not found")
)

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		return
	}

	player := r.URL.Path[len("/players/"):] // extract player name from the GET Path

	score, err := p.store.GetPlayerScore(player)
	if IsVerbose() {
		fmt.Printf("player: %s. score: %d. path: %s. err: %v\n", player, score, r.URL.Path, err)
	}

	if err == ERRPLAYERNOTFOUND {
		http.NotFound(w, r)
		return
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
