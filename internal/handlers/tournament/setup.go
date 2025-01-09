package tournament

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/BeauRussell/OpenBracket/internal/db"
	"github.com/BeauRussell/OpenBracket/internal/db/repositories"
	"github.com/BeauRussell/OpenBracket/internal/services/tournament"
)

func RenderTournamentForm(w http.ResponseWriter, r *http.Request) {
	tmpl := template.New("base")
	tmpl = template.Must(tmpl.ParseFiles(
		"internal/templates/layouts/base.html",
		"internal/templates/createTournament.html",
	))

	err := tmpl.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func CreateTournament(w http.ResponseWriter, r *http.Request) {
	dbConn := db.InitDB()

	err := r.ParseForm()
	if err != nil {
		log.Printf("Failed read request body: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	tournamentName := r.FormValue("tournament_name")

	tournamentService := tournament.NewBracketService(repositories.NewEntrantRepository(dbConn), repositories.NewTournamentRepository(dbConn))
	err, tournamentId := tournamentService.GenerateTournament(tournamentName)
	if err != nil {
		log.Printf("Failed to create tournament: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	tournamentService.CreateEntrant("Test Entrant1")
	tournamentService.CreateEntrant("Test Entrant2")

	http.Redirect(w, r, "/tournament/"+strconv.Itoa(tournamentId), http.StatusSeeOther)
}
