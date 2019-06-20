package racer

import (
	"errors"
	"net/http"
	"time"
)

const (
	TIMEOUT = 3
)

var (
	ERRTIMEOUT = errors.New("Fetching URL operation timed out")
)

func timeout(d time.Duration, timeoutChan *chan bool) {
	time.Sleep(d)
	*timeoutChan <- true
}

func ping(url string, ch *chan bool) {
	http.Get(url)
	*ch <- true
}

func Racer(a, b string) (winner string, err error) {
	pingA := make(chan bool)
	pingB := make(chan bool)
	timedOut := make(chan bool)
	go ping(a, &pingA)
	go ping(b, &pingB)
	go timeout(TIMEOUT*time.Second, &timedOut)

	select {
	case <-pingA:
		{
			return a, nil
		}
	case <-pingB:
		{
			return b, nil
		}
	case <-timedOut:
		{
			return "", ERRTIMEOUT
		}
	}
}
