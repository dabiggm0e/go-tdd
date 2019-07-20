package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type FilesystemPlayerStore struct {
	database *json.Encoder // *tape //*os.File //io.Writer //io.ReadWriteSeeker
	league   League
}

/////////////////////
//File store

func NewFilesystemPlayerStore(database *os.File) (*FilesystemPlayerStore, error) {

	if err := initializePlayerDBFile(database); err != nil {
		log.Fatalf("problem initializing the player store file %s, %v", database.Name(), err)
	}

	league, err := NewLeague(database)

	if err != nil {
		return nil, fmt.Errorf("problem loading player store from file %s, %v", database.Name(), err)
	}
	return &FilesystemPlayerStore{
		database: json.NewEncoder(&tape{database}), //&tape{database},
		league:   league,
	}, nil
}

func initializePlayerDBFile(file *os.File) error {
	file.Seek(0, 0)

	info, err := file.Stat()

	if err != nil {
		return fmt.Errorf("problem getting file info file file %s, %v", file.Name(), err)
	}

	if info.Size() == 0 {
		file.Write([]byte("[]"))
		file.Seek(0, 0)
	}

	return nil
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
