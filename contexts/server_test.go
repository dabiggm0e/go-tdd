package server

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response string
	t        *testing.T
	ctx      context.Context
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)
	s.ctx = ctx

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				s.t.Log("Spy store was cancelled")
				return

			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}

		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

func TestHandler(t *testing.T) {
	t.Run("Successful fetch", func(t *testing.T) {
		data := "Hello, world"
		spyStore := SpyStore{response: data, t: t}
		server := Server(&spyStore)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf("got '%s' want '%s'", response.Body.String(), data)
		}

		if spyStore.ctx != request.Context() {
			t.Errorf("store was not passed through a context %v", spyStore.ctx)
		}

	})

	t.Run("Cancel the fetch before 100ms", func(t *testing.T) {
		spyStore := SpyStore{response: "Hello, world", t: t}
		server := Server(&spyStore)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingctx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(time.Millisecond*5, cancel)
		request = request.WithContext(cancellingctx)

		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

	})

}
