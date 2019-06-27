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
	store map[string]int
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) (int, error) { //TODO implement inmemory store
	if score, ok := i.store[name]; ok {
		return score, nil
	}
	return 0, ERRPLAYERNOTFOUND
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

func main() {

	server := &PlayerServer{NewInMemoryPlayerStore()}

	err := http.ListenAndServe(ADDR, server)
	if err != nil {
		log.Fatalf("Couldn't listen to port %v. Err: %v", ADDR, err)
	}

}
