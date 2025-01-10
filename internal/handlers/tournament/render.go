package tournament

import (
	"html/template"
	"log"
	"net/http"

	"github.com/BeauRussell/OpenBracket/pkg/templateFunctions"
)

type TournamentPageData struct {
	TournamentID string
}

func RenderTournamentPage(w http.ResponseWriter, r *http.Request) {
	tournamentIDStr := r.URL.Path[len("/tournament/"):]

	tmpl := template.New("base").Funcs(templateFunctions.MathOps)
	tmpl = template.Must(tmpl.ParseFiles(
		"internal/templates/layouts/base.html",
		"internal/templates/bracket.html",
	))

	data := TournamentPageData{
		TournamentID: tournamentIDStr,
	}

	err := tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
