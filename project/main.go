package main

import (
	"log"
	"net/http"

	"github.com/dabiggm0e/go-tdd/project/server"
)

var (
	ADDR = ":9999"
)

func main() {
	handler := http.HandlerFunc(server.PlayerServer)
	err := http.ListenAndServe(ADDR, handler)
	if err != nil {
		log.Fatalf("Couldn't listen to port %v. Err: %v", ADDR, err)
	}

}
