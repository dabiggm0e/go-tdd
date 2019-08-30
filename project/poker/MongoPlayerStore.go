package poker



type player struct {
    id  int
    name string
    score int
}

type MongoPlayerStore struct {

}

func newMongoPlayerStore() (*MongoPlayerStore, error) {
  return &MongoPlayerStore{}, nil
}

func (m *MongoPlayerStore) GetLeague() League{
  return nil
}

func (m *MongoPlayerStore) GetPlayerScore(player string) (int, error) {
  return 0, ERRPLAYERNOTFOUND
}

func (m *MongoPlayerStore) RecordWin(player string) error {
  return nil
}
