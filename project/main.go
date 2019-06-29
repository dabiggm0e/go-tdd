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
		return 0, ERRPLAYERNOTFOUND
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

func (p *PostgresPlayerStore) RecordWin(name string) error {
	insertPlayerSql := `INSERT INTO public.players(name) VALUES($1)
	ON CONFLICT DO NOTHING;`

	_, err := p.store.Exec(insertPlayerSql, name)
	if err != nil {
		return err
	}

	id := 0
	getPlayerIdSql := `SELECT id from public.players WHERE name=$1`
	p.store.QueryRow(getPlayerIdSql, name).Scan(&id)

	if err != nil || id == 0 {
		return err
	}

	recordWinSql := `INSERT INTO public.scores(id, score)
											VALUES ($1, 1)
									ON CONFLICT ON CONSTRAINT scores_pkey
									DO
									UPDATE
										SET score=scores.score+1`

	_, err = p.store.Exec(recordWinSql, id)

	return err
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

func (i *InMemoryPlayerStore) RecordWin(name string) error {
	i.store[name]++
	return nil
}

func main() {

	//server := &PlayerServer{NewInMemoryPlayerStore()}
	//// TODO: implement a redis inmemory database
	store := NewPostgresPlayerStore()
	defer store.Teardown()
	pserver := &PlayerServer{store: store}

	err := http.ListenAndServe(ADDR, pserver)
	if err != nil {
		log.Fatalf("Couldn't listen to port %v. Err: %v", ADDR, err)
	}

}
