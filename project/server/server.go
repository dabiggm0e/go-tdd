//https://quii.gitbook.io/learn-go-with-tests/build-an-application/http-server
package server

import (
	"errors"
	"fmt"
	"net/http"
)

type PlayerStore interface {
	GetPlayerScore(name string) (int, error)
}

type PlayerServer struct {
	Store PlayerStore
}

var (
	ERRPLAYERNOTFOUND = errors.New("Player not found")
)

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		return
	}

	player := r.URL.Path[len("/players/"):] // extract player name from the GET Path

	score, err := p.Store.GetPlayerScore(player)
	fmt.Printf("player: %s. score: %d. path: %s. err: %v\n", player, score, r.URL.Path, err)
	if err == nil {
		fmt.Fprintf(w, "%d", score)
		return
	}

	fmt.Fprintf(w, "")

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
