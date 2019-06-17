package main

import (
	"bytes"
	"testing"
)

func TestCountDown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	CountDown(buffer, spySleeper)
	got := buffer.String()
	want := `3
2
1
Go`

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}

	if spySleeper.Calls != countDownStart+1 {
		t.Errorf("called '%d' should be called '%d'", spySleeper.Calls, 4)
	}
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type SpySleeper struct {
	Calls int
}
