package main

import (
	"log"
	"net/http"

	"github.com/dabiggm0e/go-tdd/project/server"
)

var (
	ADDR = ":2222"
)

type InMemoryPlayerStore struct {
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) (int, error) {
	switch name {
	case "Mo":
		return 20, nil
	case "Ziggy":
		return 10, nil
	default:
		return 0, server.ERRPLAYERNOTFOUND
	}

}

func main() {
	//handler := http.HandlerFunc(server.PlayerServer)
	//store := &InMemoryPlayerStore{}

	server := &server.PlayerServer{&InMemoryPlayerStore{}}

	err := http.ListenAndServe(ADDR, server)
	if err != nil {
		log.Fatalf("Couldn't listen to port %v. Err: %v", ADDR, err)
	}

}
