//https://quii.gitbook.io/learn-go-with-tests/build-an-application/http-server
package server

import (
	"fmt"
	"net/http"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		return
	}

	player := r.URL.Path[len("/players/"):] // extract player name from the GET Path
	score := getScore(player)
	fmt.Fprintf(w, score)
}

func getScore(player string) string {
	switch player {
	case "Mo":
		return "20"
	case "Ziggy":
		return "10"
	default:
		return ""
	}
}
