package poker

import "testing"
import "strings"

func TestCLI(t *testing.T) {
	in := strings.NewReader("Ziggy wins\n")
	playerstore := &StubPlayerStore{}
	cli := &CLI{playerstore, in}

	cli.PlayPoker()

	want := "Ziggy"
	assertPlayerWins(t, playerstore, want)

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
