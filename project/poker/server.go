//https://quii.gitbook.io/learn-go-with-tests/build-an-application/http-server
package poker

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	_ "github.com/lib/pq"
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

type PlayerStore interface {
	GetPlayerScore(name string) (int, error)
	RecordWin(name string) error
	GetLeague() League
}

type PlayerServer struct {
	store        PlayerStore
	http.Handler // embedding an http.Handler interface..
}

type Player struct {
	Name string
	Wins int
}

type PostgresPlayerStore struct {
	store *sql.DB
}

var (
	ERRPLAYERNOTFOUND = errors.New("Player not found")
)

/// Postgres store ////

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

func (p *PostgresPlayerStore) GetLeague() League {
	getLeagueSQL := `SELECT p.name, s.score FROM players p, scores s
		 						WHERE s.id = p.id`

	league := League{}

	rows, err := p.store.Query(getLeagueSQL)
	defer rows.Close()

	if err != nil {
		log.Println(err)
		return nil
	}

	for rows.Next() {
		var name string
		var wins int
		err := rows.Scan(&name, &wins)

		if err != nil {
			log.Printf("%v", err)
			return nil
		}

		league = append(league, Player{name, wins})
	}

	if rows.Err() != nil {
		log.Printf("%v", rows.Err())
		return nil
	}

	return league
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

// PlayerServer

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

func (p *PlayerServer) getLeagueTable() League {
	//return League{
	//{"Mo", 10},
	//}
	return p.store.GetLeague()
}
