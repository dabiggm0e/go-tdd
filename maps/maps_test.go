package maps

import "testing"

func TestSearch(t *testing.T) {

  dictionary := make(map[string]string)
  key := "test key"
  want := "This is a test value"
  dictionary[key] = want

  got := Search(dictionary, key)

  if got != want {
    t.Errorf("got '%s' want '%s'. Key '%s'", got, want, key)
  }
}
