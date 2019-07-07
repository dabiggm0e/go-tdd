//https://martinfowler.com/articles/practical-test-pyramid.html

package main

import (
	"flag"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	//"github.com/dabiggm0e/go-tdd/project/server"
)

//flags
var (
	verbose bool
)

const (
	ADDR   = ":1111"
	DBHOST = "localhost"
	DBPORT = 5432
	DBUSER = "postgres"
	DBPASS = "admin"
	DBNAME = "go-tdd"
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
	store := NewPostgresPlayerStore()
	defer store.Teardown()
	pserver := NewPlayerServer(store)

	err := http.ListenAndServe(ADDR, pserver)
	if err != nil {
		log.Fatalf("Couldn't listen to port %v. Err: %v", ADDR, err)
	}

}
