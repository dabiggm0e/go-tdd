package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("Successful pinging of URLs", func(t *testing.T) {
		slowServer, fastServer := initMockServers(20*time.Millisecond, 0*time.Millisecond)
		defer slowServer.Close()
		defer fastServer.Close()
		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		want := fastUrl
		got, err := Racer(slowUrl, fastUrl)

		assertError(t, err, nil)
		assertResults(t, got, want)
	})

	t.Run("Timeout in pinging URLs", func(t *testing.T) {
		slowServer, fastServer := initMockServers(10*time.Second, 11*time.Second)
		defer slowServer.Close()
		defer fastServer.Close()
		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		_, err := ConfigurableRacer(slowUrl, fastUrl, 1*time.Second)

		assertError(t, err, ERRTIMEOUT)
	})

}

func assertResults(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %#v want %#v", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %#v want %#v", got, want)
	}
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

func initMockServers(serverADelay, serverBDelay time.Duration) (serverA, serverB *httptest.Server) {
	serverA = makeDelayedServer(serverADelay)
	serverB = makeDelayedServer(serverBDelay)
	return
}
