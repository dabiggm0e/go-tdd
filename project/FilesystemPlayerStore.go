package main

import (
	"encoding/json"
	"log"
	"os"
)

type FilesystemPlayerStore struct {
	database *json.Encoder // *tape //*os.File //io.Writer //io.ReadWriteSeeker
	league   League
}

/////////////////////
//File store

func NewFilesystemPlayerStore(database *os.File) *FilesystemPlayerStore {
	database.Seek(0, 0)
	league, _ := NewLeague(database)

	return &FilesystemPlayerStore{
		database: json.NewEncoder(&tape{database}), //&tape{database},
		league:   league,
	}
}

func (f *FilesystemPlayerStore) GetPlayerScore(name string) (int, error) {

	player := f.league.Find(name)

	if player != nil {
		return player.Wins, nil
	}

	return 0, ERRPLAYERNOTFOUND

}

func (f *FilesystemPlayerStore) RecordWin(name string) error {

	player := f.league.Find(name)

	if player != nil {
		player.Wins++

	} else {
		f.league = append(f.league, Player{Name: name, Wins: 1})
	}

	//f.database.Seek(0, 0)
	//err := json.NewEncoder(f.database).Encode(f.league)
	err := f.database.Encode(f.league)

	if err != nil {
		log.Printf("Couldn't encode to json, %v", err)
	}
	return err
}

func (f *FilesystemPlayerStore) GetLeague() League {
	return f.league
}

/////////////////////
