package server

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response  string
	cancelled bool
}

func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func TestHandler(t *testing.T) {
	t.Run("Successful fetch", func(t *testing.T) {
		data := "Hello, world"
		spyStore := SpyStore{response: data}
		server := Server(&spyStore)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf("got '%s' want '%s'", response.Body.String(), data)
		}

		if spyStore.cancelled {
			t.Error("Store was told to cancel")
		}
	})

	t.Run("Cancel the fetch before 100ms", func(t *testing.T) {
		spyStore := SpyStore{"Hello, world", false}
		server := Server(&spyStore)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingctx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(time.Millisecond*5, cancel)
		request = request.WithContext(cancellingctx)

		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		if !spyStore.cancelled {
			t.Error("Store was not told to cancel")
		}
	})

}
