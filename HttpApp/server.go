package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)


type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() []Player
}

type PlayerServer struct {
	store PlayerStore
	http.Handler //embedding -> borrowing implementation
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := new(PlayerServer)
	
	p.store = store

	router := http.NewServeMux()

	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playerHandler))

	p.Handler = router

	return p
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(p.store.GetLeague())
}

func (p *PlayerServer) playerHandler(w http.ResponseWriter, r *http.Request){
	player := strings.TrimPrefix(r.URL.Path, "/players/");
	
	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

type Player struct{
	Name string
	Wins int
}
// -----------------------------------------------------------------------------------------------------------------------------------------------------------

// //A ResponseWriter interface is used by an HTTP handler to construct an HTTP response.
// func PlayerServer(w http.ResponseWriter, r *http.Request) {
// 	player := strings.TrimPrefix(r.URL.Path, "/players/")

// 	//ResponseWriter also implements io's Writer, so we can use fmt.Fprint to send strings as HTTP responses.
// 	fmt.Fprint(w, GetPlayerScore(player))
// }

// func GetPlayerScore(name string) string {
// 	if name == "Pepper" {
// 		return "20"
// 	}

// 	if name == "Floyd" {
// 		return "10"
// 	}

// 	return ""
// }
