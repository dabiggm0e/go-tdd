//https://martinfowler.com/articles/practical-test-pyramid.html

package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/dabiggm0e/go-tdd/project/poker"
	_ "github.com/lib/pq"
)

//flags
var (
	verbose bool
)

const (
	ADDR       = ":1111"
	DBHOST     = "localhost"
	DBPORT     = 5432
	DBUSER     = "postgres"
	DBPASS     = "admin"
	DBNAME     = "go-tdd"
	dbFilename = "game.db.json"
)

func init() {
	flag.BoolVar(&verbose, "verbose", false, "Show verbose messages")
	flag.Parse()
}

func IsVerbose() bool {
	return verbose
}

func main() {

	//server := &PlayerServer{NewInMemoryPlayerStore()}
	//// TODO: implement a redis inmemory database

	//store := NewPostgresPlayerStore()
	//defer store.Teardown()
	store, closeFunc, err := poker.NewFilesystemPlayerStoreFromFile(dbFilename)
	if err != nil {
		log.Fatal(err)
	}

	defer closeFunc()

	pserver := poker.NewPlayerServer(store)

	err = http.ListenAndServe(ADDR, pserver)
	if err != nil {
		log.Fatalf("Couldn't listen to port %v. Err: %v", ADDR, err)
	}

}
