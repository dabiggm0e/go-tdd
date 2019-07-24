package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dabiggm0e/go-tdd/project/poker"
)

const dbFileName = "game.json.db"

func main() {

	store, closeFunc, err := poker.NewFilesystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer closeFunc()

	cli := poker.NewCLI(store, os.Stdin)

	fmt.Println("Let's player poker")
	fmt.Println("Type '{PlayerName} wins'")

	cli.PlayPoker()
}
