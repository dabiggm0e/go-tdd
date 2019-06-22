package reflect

import "testing"

func TestWalk(t *testing.T) {
	expected := "Mo"
	var got []string

	x := struct {
		Name string
	}{expected}

	walk(x, func(input string) {
		got = append(got, input)
	})

	if len(got) != 1 {
		t.Errorf("wrong number of function calls. Want %d got %d", 1, len(got))
	}

	if got[0] != expected {
		t.Errorf("got %s want %s", got[0], expected)
	}
}
