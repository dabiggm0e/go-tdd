package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Mo")
	want := "Hello, Mo"

	if got != want {
		t.Errorf("Got '%s', want '%s'", got, want)
	}
}
