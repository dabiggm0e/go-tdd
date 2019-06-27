//https://martinfowler.com/articles/practical-test-pyramid.html

package main

import (
	"flag"
	"log"
	"net/http"
	//"github.com/dabiggm0e/go-tdd/project/server"
)

//flags
var (
	verbose bool
)

var (
	ADDR = ":2222"
)

func init() {
	flag.BoolVar(&verbose, "verbose", false, "Show verbose messages")
	flag.Parse()
}

func IsVerbose() bool {
	return verbose
}

type InMemoryPlayerStore struct {
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) (int, error) {
	switch name {
	case "Mo":
		return 20, nil
	case "Ziggy":
		return 10, nil
	default:
		return 0, ERRPLAYERNOTFOUND
	}
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	return //21, nil //TODO implement a store
}

func main() {
	//handler := http.HandlerFunc(server.PlayerServer)
	//store := &InMemoryPlayerStore{}

	server := &PlayerServer{&InMemoryPlayerStore{}}

	err := http.ListenAndServe(ADDR, server)
	if err != nil {
		log.Fatalf("Couldn't listen to port %v. Err: %v", ADDR, err)
	}

}
