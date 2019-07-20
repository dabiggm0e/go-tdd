package main

import (
	"encoding/json"
	"fmt"
	"io"
)

type League []Player

func (l League) Find(name string) *Player {
	for i, player := range l {
		if player.Name == name {
			return &l[i]
		}
	}

	return nil
}

func NewLeague(rdr io.Reader) ([]Player, error) {
	var league []Player

	err := json.NewDecoder(rdr).Decode(&league)

	if err != nil {
		err = fmt.Errorf("Error parsing json from league, %v", err)
		return nil, err
	}

	return league, nil
}
