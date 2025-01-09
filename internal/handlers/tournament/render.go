package tournament

import (
	"html/template"
	"log"
	"net/http"

	"github.com/BeauRussell/OpenBracket/pkg/templateFunctions"
)

func RenderTournamentPage(w http.ResponseWriter, r *http.Request) {
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
