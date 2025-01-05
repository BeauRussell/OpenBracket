package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type Match struct {
	ID   int
	Name string
}

type Round struct {
	Matches []Match
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", renderBracketForm)
	mux.HandleFunc("/generate-bracket", generateBracket)
	mux.HandleFunc("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))).ServeHTTP)

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func renderBracketForm(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("internal/templates/layouts/base.html", "internal/templates/bracket.html"))
	err := tmpl.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func generateBracket(w http.ResponseWriter, r *http.Request) {
	numMatchesStr := r.URL.Query().Get("num_matches")
	numMatches, err := strconv.Atoi(numMatchesStr)
	if err != nil || numMatches < 1 {
		http.Error(w, "Invalid number of matches", http.StatusBadRequest)
		return
	}

	// Generate rounds
	rounds := []Round{}
	matches := []Match{}

	for i := 1; i <= numMatches; i++ {
		matches = append(matches, Match{
			ID:   i,
			Name: fmt.Sprintf("Match %d", i),
		})
	}

	// Split matches into rounds
	for len(matches) > 0 {
		rounds = append(rounds, Round{Matches: matches})
		matches = nextRound(matches)
	}

	tmpl := template.Must(template.ParseFiles("internal/templates/bracket.html"))
	tmpl.ExecuteTemplate(w, "bracket", rounds)
}

func nextRound(matches []Match) []Match {
	next := []Match{}
	for i := 0; i < len(matches)/2; i++ {
		next = append(next, Match{
			ID:   i + 1,
			Name: fmt.Sprintf("Winner of Match %d & %d", matches[i*2].ID, matches[i*2+1].ID),
		})
	}
	return next
}
