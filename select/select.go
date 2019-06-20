package racer

import (
	"errors"
	"net/http"
	"time"
)

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Now().Sub(start)
}

func Racer(a, b string) (winner string, err error) {

	aDuration := measureResponseTime(a)
	bDuration := measureResponseTime(b)

	winner, err = "", nil
	if aDuration < bDuration {
		return a, nil
	} else if bDuration < aDuration {
		return b, nil
	} else {
		err = errors.New("No one won")
	}
	return "", err
}
