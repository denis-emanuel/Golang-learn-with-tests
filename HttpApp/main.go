package main

import (
	"log"
	"net/http"
)

func main() {
	// //By type casting the PlayerServer function with HandlerFunc we have implemented the required Handler
	// handler := http.HandlerFunc(PlayerServer)
	// //ListenAndServe takes a port to listen on a Handler
	// log.Fatal(http.ListenAndServe(":5000", handler))

	server := NewPlayerServer(NewInMemoryPlayerStore())
	
	log.Fatal(http.ListenAndServe(":5000", server))
}