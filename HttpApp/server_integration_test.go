package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	//The components we are trying to integrate with
	store := InMemoryPlayerStore{}
	server := NewPlayerServer(&store)
	player := "Pepper"

	//fire off 3 requests to record 3 wins for 'player'
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetScoreRequest(player))
	assertStatus(t, response.Code, http.StatusOK)

	//assert there are 3 wins for player as recorded
	assertResponseBody(t, response.Body.String(), "3")
}