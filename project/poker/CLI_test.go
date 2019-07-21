package poker

import "testing"
import "strings"

func TestCLI(t *testing.T) {

	t.Run("Record win for Ziggy", func(t *testing.T) {
		in := strings.NewReader("Ziggy wins")
		playerstore := &StubPlayerStore{}
		cli := &CLI{playerstore, in}

		cli.PlayPoker()

		want := "Ziggy"
		assertPlayerWins(t, playerstore, want)
	})

	t.Run("Record win for Mo", func(t *testing.T) {
		in := strings.NewReader("Mo wins")
		playerstore := &StubPlayerStore{}
		cli := &CLI{playerstore, in}

		cli.PlayPoker()

		want := "Mo"
		assertPlayerWins(t, playerstore, want)
	})
}

func assertPlayerWins(t *testing.T, store *StubPlayerStore, winner string) {
	t.Helper()

	if len(store.winCalls) != 1 {
		t.Fatalf("expected a win call but didn't get any")
	}

	got := store.winCalls[0]
	if got != winner {
		t.Errorf("didn't record correct winner. Got %q want %q", got, winner)
	}
}
