package bracket

import (
	"log"
	"math"
	"net/http"
	"strconv"
	"text/template"

	"github.com/BeauRussell/OpenBracket/internal/db"
	"github.com/BeauRussell/OpenBracket/internal/db/models"
	"github.com/BeauRussell/OpenBracket/internal/db/repositories"
	"github.com/BeauRussell/OpenBracket/internal/services/match"
	"github.com/BeauRussell/OpenBracket/pkg/templateFunctions"
)

type Round struct {
	Matches []models.Match
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

func GenerateBracket(w http.ResponseWriter, r *http.Request) {
	dbConn := db.InitDB()

	err := r.ParseForm()
	if err != nil {
		log.Printf("Failed read request body: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	numEntrantsStr := r.FormValue("num_entrants")
	tournamentIDStr := r.FormValue("tournament_id")
	tournamentID, err := strconv.Atoi(tournamentIDStr)
	if err != nil {
		log.Printf("tournamentID invalid type: %v", err)
		http.Error(w, "tournamentID is an invalid type", http.StatusBadRequest)
	}
	numEntrants, err := strconv.Atoi(numEntrantsStr)

	matchService := match.NewMatchService(repositories.NewEntrantRepository(dbConn), repositories.NewTournamentRepository(dbConn), repositories.NewMatchRepository(dbConn))

	err, matches := matchService.CreateMatches(numEntrants, tournamentID)
	if err != nil {
		log.Printf("Failed to create matches: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	log.Println(matches)
}

func numMatchesRound1(numEntrants int) int {
	log2 := math.Log2(float64(numEntrants))
	return int(math.Pow(2, math.Floor(log2)))
}
