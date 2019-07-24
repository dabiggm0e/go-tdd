package poker

import "testing"

type StubPlayerStore struct {
	score    map[string]int
	winCalls []string
	league   League
}

func (s *StubPlayerStore) GetLeague() League {
	return s.league
}

func InitPlayersStore() *StubPlayerStore {
	store := StubPlayerStore{
		score: map[string]int{
			"Mo":    20,
			"Ziggy": 10,
		},
		league: League{
			{"Mo", 20},
			{"Ziggy", 10},
		},
	}

	return &store
}

func (s *StubPlayerStore) GetPlayerScore(name string) (int, error) {

	if score, ok := s.score[name]; ok {
		return score, nil
	}
	return 0, ERRPLAYERNOTFOUND
}

func (s *StubPlayerStore) RecordWin(name string) error {
	//s.score[name]++
	//return s.score[name], nil
	s.winCalls = append(s.winCalls, name)
	return nil
}

func AssertPlayerWins(t *testing.T, store *StubPlayerStore, winner string) {
	t.Helper()

	if len(store.winCalls) != 1 {
		t.Fatalf("expected a win call but didn't get any")
	}

	got := store.winCalls[0]
	if got != winner {
		t.Errorf("didn't record correct winner. Got %q want %q", got, winner)
	}
}
