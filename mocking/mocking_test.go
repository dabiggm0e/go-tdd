package main

import (
	"bytes"
	"testing"
)

func TestCountDown(t *testing.T) {
	buffer := &bytes.Buffer{}
	CountDown(buffer)
	got := buffer.String()
	//want := "3\n2\n1\nGo"
	want := `3
2
1
Go`

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
