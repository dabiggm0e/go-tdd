//https://martinfowler.com/articles/practical-test-pyramid.html

package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	//"github.com/dabiggm0e/go-tdd/project/server"
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

	db, err := os.OpenFile(dbFilename, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("problem opening %s %v", dbFilename, err)
	}

	//store := NewPostgresPlayerStore()
	//defer store.Teardown()
	store, err := NewFilesystemPlayerStore(db) //&FilesystemPlayerStore{db}
	if err != nil {
		log.Fatal(err.Error())
	}

	pserver := NewPlayerServer(store)

	err = http.ListenAndServe(ADDR, pserver)
	if err != nil {
		log.Fatalf("Couldn't listen to port %v. Err: %v", ADDR, err)
	}

}
