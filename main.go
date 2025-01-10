package main

import (
	"log"
	"net/http"

	"github.com/BeauRussell/OpenBracket/internal/handlers/bracket"
	"github.com/BeauRussell/OpenBracket/internal/handlers/tournament"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", tournament.RenderTournamentForm)
	mux.HandleFunc("/create-tournament", tournament.CreateTournament)
	mux.HandleFunc("/tournament/", tournament.RenderTournamentPage)
	mux.HandleFunc("/generate-bracket", bracket.GenerateBracket)
	mux.HandleFunc("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))).ServeHTTP)

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
