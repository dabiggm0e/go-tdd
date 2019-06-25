//https://quii.gitbook.io/learn-go-with-tests/build-an-application/http-server
package main

import (
	"fmt"
	"net/http"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {

	player := r.URL.Path[len("/players/"):] // extract player name from the GET Path

	switch player {
	case "Mo":
		fmt.Fprintf(w, "20")
	case "Ziggy":
		fmt.Fprintf(w, "10")
	}

}
