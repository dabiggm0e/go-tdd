package racer

import (
	"errors"
	"net/http"
	"time"
)

func Racer(url1, url2 string) (winner string, err error) {
	startA := time.Now()
	http.Get(url1)
	aDuration := time.Now().Sub(startA)

	startB := time.Now()
	http.Get(url2)
	bDuration := time.Now().Sub(startB)

	winner, err = "", nil
	if aDuration < bDuration {
		winner = url1
	} else if bDuration < aDuration {
		winner = url2
	} else {
		err = errors.New("No one won")
	}
	return winner, nil
}
