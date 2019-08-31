package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dabiggm0e/go-tdd/project/poker"
)

const dbFileName = "game.json.db"

func main() {

	//store, closeFunc, err := poker.NewFilesystemPlayerStoreFromFile(dbFileName)
	store, err := poker.NewMongoPlayerStore("")
	if err != nil {
		log.Fatal(err)
	}
	defer store.Teardown()
	//defer closeFunc()

	cli := poker.NewCLI(store, os.Stdin)

	fmt.Println("Let's player poker")
	fmt.Println("Type '{PlayerName} wins'")

	cli.PlayPoker()
}
