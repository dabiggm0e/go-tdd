//https://martinfowler.com/articles/practical-test-pyramid.html

package main

import (
	"database/sql"
	"flag"
	"fmt"
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

type PostgresPlayerStore struct {
	store *sql.DB
}

func NewPostgresPlayerStore() *PostgresPlayerStore {
	pSqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		DBHOST, DBPORT, DBUSER, DBPASS, DBNAME)

	db, err := sql.Open("postgres", pSqlInfo)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return &PostgresPlayerStore{store: db}
}

func (p *PostgresPlayerStore) Teardown() {
	p.store.Close()
}

func (p *PostgresPlayerStore) GetPlayerScore(name string) (int, error) {
	id, err := p.getPlayerIdSql(name)

	if err != nil {
		return 0, nil
	}

	var score int
	getPlayerScoreSql := `SELECT "score" FROM scores WHERE id=$1;`
	row := p.store.QueryRow(getPlayerScoreSql, id)

	switch err := row.Scan(&score); err {
	case sql.ErrNoRows:
		return 0, ERRPLAYERNOTFOUND
	case nil:
		return score, nil
	default:
		return 0, err
	}

}

func (p *PostgresPlayerStore) getPlayerIdSql(name string) (int, error) {
	var id int
	getUserIdSql := `SELECT "id" FROM players WHERE name=$1;`
	row := p.store.QueryRow(getUserIdSql, name)

	switch err := row.Scan(&id); err {
	case sql.ErrNoRows:
		return 0, ERRPLAYERNOTFOUND
	case nil:
		return id, nil
	default:
		return 0, err
	}

}

func (p *PostgresPlayerStore) RecordWin(name string) {}

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

	//server := &PlayerServer{NewInMemoryPlayerStore()}
	store := NewPostgresPlayerStore()
	defer store.Teardown()
	pserver := &PlayerServer{store: store}

	err := http.ListenAndServe(ADDR, pserver)
	if err != nil {
		log.Fatalf("Couldn't listen to port %v. Err: %v", ADDR, err)
	}

}
