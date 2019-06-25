//https://quii.gitbook.io/learn-go-with-tests/build-an-application/http-server
package server

import (
	"fmt"
	"net/http"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		return
	}

	player := r.URL.Path[len("/players/"):] // extract player name from the GET Path

	score := p.store.GetPlayerScore(player)
	fmt.Fprintf(w, string(score))
}

func (p *PlayerServer) GetPlayerScore(name string) string {
	switch name {
	case "Mo":
		return "20"
	case "Ziggy":
		return "10"
	default:
		return ""
	}
}
