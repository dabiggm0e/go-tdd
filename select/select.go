package racer

import (
	"errors"
	"net/http"
	"time"
)

var (
	ERRTIMEOUT = errors.New("Fetching URL operation timed out")
)

func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(url)
		close(ch)
	}()

	return ch
}

func Racer(a, b string, timeout time.Duration) (winner string, err error) {

	select {
	case <-ping(a):
		{
			return a, nil
		}
	case <-ping(b):
		{
			return b, nil
		}
	case <-time.After(timeout):
		{
			return "", ERRTIMEOUT
		}
	}
}
