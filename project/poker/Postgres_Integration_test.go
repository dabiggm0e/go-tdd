package poker

/*
func TestPostgresGetPlayer(t *testing.T) {
	t.Run("PostgresPlayerStore: Return 404 on not found player", func(t *testing.T) {
		request := newGetScoreRequest("JOHNDOE")
		response := httptest.NewRecorder()

		store := NewPostgresPlayerStore()
		defer store.Teardown()
		clearPostgresStore(t, store)

		playerServer := NewPlayerServer(store)
		playerServer.ServeHTTP(response, request)

		want := http.StatusNotFound
		got := response.Code

		assertStatusCode(t, got, want)
	})

	t.Run("Getting Mo's score", func(t *testing.T) {
		player := "Mo"
		response := httptest.NewRecorder()

		store := NewPostgresPlayerStore()
		defer store.Teardown()
		clearPostgresStore(t, store)

		playerServer := NewPlayerServer(store)
		playerServer.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(player))
		playerServer.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(player))
		playerServer.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(player))

		playerServer.ServeHTTP(response, newGetScoreRequest(player))

		assertStatusCode(t, response.Code, http.StatusOK)
		assertResponseReply(t, response.Body.String(), "3")
	})

}

func TestPostgresStoreWin(t *testing.T) {
	player := "Ziggy"

	store := NewPostgresPlayerStore()
	defer store.Teardown()
	clearPostgresStore(t, store)

	playerServer := NewPlayerServer(store)

	response := httptest.NewRecorder()
	playerServer.ServeHTTP(response, newPostScoreRequest(player))

	assertStatusCode(t, response.Code, http.StatusAccepted)
}


func TestPostgresStoreRecordWinsAndRetrieveLeagueInJson(t *testing.T) {

	store := NewPostgresPlayerStore()
	defer store.Teardown()
	clearPostgresStore(t, store)

	server := NewPlayerServer(store)

	wantedLeague := League{
		{"Mo", 9},
		{"Ziggy", 17},
		{"Su", 12},
	}

	for _, p := range wantedLeague {
		for i := 0; i < p.Wins; i++ {
			server.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(p.Name))
		}
	}

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetLeagueRequest())

	gotLeague := getLeagueFromResponse(t, response.Body)
	assertStatusCode(t, response.Code, http.StatusOK)
	assertResponseContentType(t, response, jsonContentType)
	assertLeague(t, gotLeague, wantedLeague) //TODO: Test whether the order of the league affects the DeepEqual
}

func clearPostgresStore(t *testing.T, p *PostgresPlayerStore) {
	truncateSql := "TRUNCATE scores; TRUNCATE players"
	_, err := p.store.Exec(truncateSql)
	if err != nil {
		t.Fatalf("Unable to truncate the store. Err: %v", err)
	}
}

func TestPostgresStoreRecordWinsAndRetrieveScore(t *testing.T) {
	database, cleanDatabase := createTempFile(t, "[]")
	defer cleanDatabase()
	store, err := NewFilesystemPlayerStore(database) //NewPostgresPlayerStore()
	assertNoError(t, err)
	//defer store.Teardown()
	//clearPostgresStore(t, store)
	server := NewPlayerServer(store)
	player := "Mo"

	server.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetScoreRequest(player))

	assertStatusCode(t, response.Code, http.StatusOK)
	assertResponseReply(t, response.Body.String(), "3")
}

*/
