package bracket

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"text/template"

	"github.com/BeauRussell/OpenBracket/pkg/templateFunctions"
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

func RenderBracketForm(w http.ResponseWriter, r *http.Request) {
	tmpl := template.New("base").Funcs(templateFunctions.MathOps)
	tmpl = template.Must(tmpl.ParseFiles(
		"internal/templates/layouts/base.html",
		"internal/templates/bracket.html",
	))

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
