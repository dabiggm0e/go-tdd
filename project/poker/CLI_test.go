package poker_test

import (
	"strings"
	"testing"

	"github.com/dabiggm0e/go-tdd/project/poker"
)

func TestCLI(t *testing.T) {

	t.Run("Record win for Ziggy", func(t *testing.T) {
		in := strings.NewReader("Ziggy wins")
		playerstore := &poker.StubPlayerStore{}
		cli := poker.NewCLI(playerstore, in)

		cli.PlayPoker()

		want := "Ziggy"
		poker.AssertPlayerWins(t, playerstore, want)
	})

	t.Run("Record win for Mo", func(t *testing.T) {
		in := strings.NewReader("Mo wins")
		playerstore := &poker.StubPlayerStore{}
		cli := poker.NewCLI(playerstore, in)

		cli.PlayPoker()

		want := "Mo"
		poker.AssertPlayerWins(t, playerstore, want)
	})
}

/*func assertPlayerWins(t *testing.T, store *poker.StubPlayerStore, winner string) {
	t.Helper()

	if len(store.winCalls) != 1 {
		t.Fatalf("expected a win call but didn't get any")
	}

	got := store.winCalls[0]
	if got != winner {
		t.Errorf("didn't record correct winner. Got %q want %q", got, winner)
	}
}*/
