package main

import (
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(PlayerServer)
	if err := http.ListenAndServe(":9999", handler); err != nil {
		log.Fatalf("Couldn't listen to port 9999. Err: %v", err)
	}

}
