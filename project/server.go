//https://quii.gitbook.io/learn-go-with-tests/build-an-application/http-server
package main

import (
	"errors"
	"fmt"
	"net/http"
)

type PlayerStore interface {
	GetPlayerScore(name string) (int, error)
	RecordWin(name string)
}

type PlayerServer struct {
	store PlayerStore
}

var (
	ERRPLAYERNOTFOUND = errors.New("Player not found")
)

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := getPlayerName(r)
	switch r.Method {
	case "GET":
		p.showScore(w, player)

	case "POST":
		p.processWin(w, player)
	}

}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score, err := p.store.GetPlayerScore(player)
	if err == ERRPLAYERNOTFOUND {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprintf(w, "%d", score)
}

func getPlayerName(r *http.Request) string {
	return r.URL.Path[len("/players/"):] // extract player name from the GET Path
}