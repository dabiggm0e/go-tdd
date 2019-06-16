package maps

import "testing"

func TestSearch(t *testing.T) {
	key := "test key"
	want := "This is a test value"
	dict := Dictionary{key: want}

	t.Run("unknown word", func(t *testing.T) {
		key := "Unknown"
		_, got := dict.Search(key)
		if got == nil {
			t.Fatal("Expected to get an error")
		}
		assertErrors(t, got, ErrNotFound)
	})

	t.Run("Known word", func(t *testing.T) {
		got, _ := dict.Search(key)
		assertStrings(t, key, got, want)
	})

}

func assertErrors(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error '%s' want '%s'", got, want)
	}
}
func assertStrings(t *testing.T, key, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got '%s' want '%s'. Key '%s'", got, want, key)
	}
}
