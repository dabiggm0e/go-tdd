package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dabiggm0e/go-tdd/project/poker"
)

const dbFileName = "game.json.db"

func main() {

	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("Error opening the database file %v. %v", dbFileName, err.Error())
	}

	store, err := poker.NewFilesystemPlayerStore(db)
	if err != nil {
		log.Fatalf("Error in creating the filesystem store. %v", err.Error())
	}

	cli := &poker.CLI{PlayerStore: store, Input: os.Stdin}

	fmt.Println("Let's player poker")
	fmt.Println("Type '{PlayerName} wins'")

	cli.PlayPoker()
}
