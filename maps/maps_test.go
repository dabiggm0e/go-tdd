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

func TestInsert(t *testing.T) {

	t.Run("New word", func(t *testing.T) {
		dict := Dictionary{}
		word := "word"
		want := "description"
		err := dict.Insert(word, want)
		assertErrors(t, err, nil)
		assertDefinition(t, dict, word, want)
	})

	t.Run("Existing word", func(t *testing.T) {
		word := "word"
		want := "description"
		dict := Dictionary{word: want}
		err := dict.Insert(word, want)
		assertErrors(t, err, ErrWordAlreadyExists)
		assertDefinition(t, dict, word, want)
	})

}

func assertErrors(t *testing.T, got, want error) {
	t.Helper()

	if got == want {
		return
	}

	if got == ErrWordAlreadyExists {
		t.Fatal(got.Error())
	}

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
func assertDefinition(t *testing.T, dict Dictionary, word, want string) {
	t.Helper()
	got, err := dict.Search(word)

	if err != nil {
		t.Fatalf("Should find added word: %s", err)
	}
	if got != want {
		t.Errorf("got '%s' want '%s'. %s", got, want, word)
	}
}
