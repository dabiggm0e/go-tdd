package poker

import "testing"
import "strings"

func TestCLI(t *testing.T) {
	in := strings.NewReader("Ziggy wins\n")
	playerstore := &StubPlayerStore{}
	cli := &CLI{playerstore, in}

	cli.PlayPoker()

	if len(playerstore.winCalls) < 1 {
		t.Fatalf("expected a win call but didn't get any")
	}

	got := playerstore.winCalls[0]
	want := "Ziggy"

	if got != want {
		t.Errorf("didn't record correct winner. Got %q want %q", got, want)
	}
}
