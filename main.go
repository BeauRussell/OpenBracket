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

var FuncMap template.FuncMap = template.FuncMap{
	"len": func(slice []Match) int {
		return len(slice)
	},
	"sub": func(a, b int) int {
		return a - b
	},
	"add": func(a, b int) int {
		return a + b
	},
	"mod": func(a, b int) int {
		return a % b
	},
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
	// Create a new template, add functions, and parse files
	tmpl := template.New("base").Funcs(FuncMap)
	tmpl = template.Must(tmpl.ParseFiles(
		"internal/templates/layouts/base.html",
		"internal/templates/bracket.html",
	))

	// Execute the "base" template
	err := tmpl.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func generateBracket(w http.ResponseWriter, r *http.Request) {
	numMatchesStr := r.URL.Query().Get("num_matches")
	numMatches, err := strconv.Atoi(numMatchesStr)
	if err != nil || numMatches < 1 {
		http.Error(w, "Invalid number of matches", http.StatusBadRequest)
		return
	}

	// Generate matches and rounds
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

	// Add custom functions to the template
	tmpl := template.Must(template.New("bracket.html").Funcs(FuncMap).ParseFiles("internal/templates/bracket.html"))

	// Render the template
	err = tmpl.ExecuteTemplate(w, "bracket", rounds)
	if err != nil {
		log.Printf("Failed to execute bracket template %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
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
