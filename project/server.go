//https://quii.gitbook.io/learn-go-with-tests/build-an-application/http-server
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) (int, error)
	RecordWin(name string) error
	GetLeague() []Player
}

type PlayerServer struct {
	store        PlayerStore
	http.Handler // embedding an http.Handler interface.
}

type Player struct {
	Name string
	Wins int
}

var (
	ERRPLAYERNOTFOUND = errors.New("Player not found")
)

func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := new(PlayerServer)
	p.store = store

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))

	p.Handler = router
	return p
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {

	league := p.getLeagueTable()
	if len(league) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(league)

}

func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	player := getPlayerName(r)
	switch r.Method {

	case "GET":
		p.showScore(w, player)

	case "POST":
		p.processWin(w, player)
	}
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	err := p.store.RecordWin(player)
	switch err {
	case nil:
		w.WriteHeader(http.StatusAccepted)
	default:
		w.WriteHeader(http.StatusNotFound)
	}

}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score, err := p.store.GetPlayerScore(player)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprintf(w, "%d", score)
}

func getPlayerName(r *http.Request) string {

	tokens := strings.SplitN(r.URL.Path, "/", -1)

	if len(tokens) > 2 { // example: "/players/Mo" >> ["", "players" "Mo"]
		return tokens[2]
	}

	return ""
}

func (p *PlayerServer) getLeagueTable() []Player {
	//return []Player{
	//{"Mo", 10},
	//}
	return p.store.GetLeague()
}
