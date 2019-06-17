//https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/dependency-injection
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func GreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	Greet(os.Stdout, "Mo")

	http.ListenAndServe(":9999", http.HandlerFunc(GreetHandler))
}
