//https://quii.gitbook.io/learn-go-with-tests/build-an-application/http-server
package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) (int, error)
	RecordWin(name string) error
}

type PlayerServer struct {
	store PlayerStore
}

var (
	ERRPLAYERNOTFOUND = errors.New("Player not found")
)

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	router := http.NewServeMux()

	router.Handle("/league", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	router.Handle("/players/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			p.processGetRequests(w, r)

		case "POST":
			p.processPostRequests(w, r)
		}
	}))

	router.ServeHTTP(w, r)
}

func (p *PlayerServer) processGetRequests(w http.ResponseWriter, r *http.Request) {

	switch getEndpoint(r) {

	case "players":
		player := getPlayerName(r)
		p.showScore(w, player)

	default:
		w.WriteHeader(http.StatusNotFound)
	}

}

func (p *PlayerServer) processPostRequests(w http.ResponseWriter, r *http.Request) {
	//endpoint:= strings.Spli
	switch getEndpoint(r) {
	case "players":
		player := getPlayerName(r)
		p.processWin(w, player)

	default:
		w.WriteHeader(http.StatusNotFound)
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

func getEndpoint(r *http.Request) string {

	tokens := strings.SplitN(r.URL.Path, "/", -1)

	if len(tokens) > 1 { // example: "/players/Mo" >> ["", "players" "Mo"]
		return tokens[1]
	}

	return ""

}
