//https://quii.gitbook.io/learn-go-with-tests/build-an-application/http-server
package gameserver

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func PlayerServer(w *httptest.ResponseRecorder, r *http.Request) {
	fmt.Fprintf(w, "20")
}
