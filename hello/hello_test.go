package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Mo", "")
		want := "Hello, Mo"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Say 'Hello, world' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Mo", spanish)
		want := "Hola, Mo"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string in Spanish", func(t *testing.T) {
		got := Hello("", spanish)
		want := "Hola, world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Mo", "French")
		want := "Bonjour, Mo"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string in French", func(t *testing.T) {
		got := Hello("", french)
		want := "Bonjour, world"
		assertCorrectMessage(t, got, want)
	})
}
