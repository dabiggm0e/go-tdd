package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Mo")
	got := buffer.String()
	want := "Hello, Mo"
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
