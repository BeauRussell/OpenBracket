package main

import (
	"log"
	"net/http"

	"github.com/BeauRussell/OpenBracket/internal/handlers/bracket"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", bracket.RenderBracketForm)
	mux.HandleFunc("/generate-bracket", bracket.GenerateBracket)
	mux.HandleFunc("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))).ServeHTTP)

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
