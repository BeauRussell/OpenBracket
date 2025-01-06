package bracket

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"reflect"
	"strconv"
	"text/template"
)

type Entrant struct {
	Name string
}

type Match struct {
	ID       int
	Name     string
	Entrants [2]Entrant
}

type Round struct {
	Matches []Match
}

var FuncMap template.FuncMap = template.FuncMap{
	"len": GenericLen,
	"sub": func(a, b int) int {
		return a - b
	},
	"add": func(a, b int) int {
		return a + b
	},
	"mod": func(a, b int) int {
		return a % b
	},
	"div": func(a, b int) int {
		return a / b
	},
	"mul": func(a, b int) int {
		return a * b
	},
}

func GenerateBracket(w http.ResponseWriter, r *http.Request) {
	numEntrantsStr := r.URL.Query().Get("num_entrants")
	numEntrants, err := strconv.Atoi(numEntrantsStr)
	if err != nil || numEntrants < 2 {
		http.Error(w, "Invalid number of matches", http.StatusBadRequest)
		return
	}

	// Generate matches and rounds
	entrants := []Entrant{}
	rounds := []Round{}
	matches := []Match{}

	for i := 1; i <= numEntrants; i++ {
		entrants = append(entrants, Entrant{
			Name: fmt.Sprintf("Entrant %d", i),
		})
	}

	numMatches := numMatchesRound1(numEntrants - 1)

	for i := 1; i <= numMatches; i++ {
		matches = append(matches, Match{
			ID:   i,
			Name: fmt.Sprintf("Match %d", i),
			Entrants: [2]Entrant{
				{Name: "test 1"},
				{Name: "test 2"},
			},
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

func RenderBracketForm(w http.ResponseWriter, r *http.Request) {
	// Create a new template, add functions, and parse files
	tmpl := template.New("base").Funcs(FuncMap)
	tmpl = template.Must(tmpl.ParseFiles(
		"internal/templates/layouts/base.html",
		"internal/templates/bracket.html",
		"internal/templates/partials/navbar.html",
	))

	// Execute the "base" template
	err := tmpl.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Printf("Error executing template: %v", err)
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

func numMatchesRound1(numEntrants int) int {
	log2 := math.Log2(float64(numEntrants))
	return int(math.Pow(2, math.Floor(log2)))
}

func GenericLen(slice interface{}) int {
	v := reflect.ValueOf(slice)

	// Check if the input is a slice
	if v.Kind() == reflect.Slice {
		return v.Len()
	}
	// If it's not a slice, return 0 or handle error
	fmt.Printf("Invalid type: expected slice, got %s\n", v.Kind())
	return 0
}
