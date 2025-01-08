package repositories

import (
	"database/sql"
	"log"
	"strings"

	"github.com/BeauRussell/OpenBracket/internal/db/models"
	_ "github.com/mattn/go-sqlite3"
)

type TournamentRepository struct {
	DB *sql.DB
}

func NewTournamentRepository(db *sql.DB) *TournamentRepository {
	return &TournamentRepository{DB: db}
}

func (r *TournamentRepository) GetTournamentById(id int) (*models.Tournament, error) {
	query := "SELECT id, name, num_entrants FROM tournaments WHERE id = $1"
	row := r.DB.QueryRow(query, id)

	var tournament models.Tournament
	if err := row.Scan(&tournament.ID, &tournament.Name); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Printf("Error fetching tournaments by ID: %v\n", err)
		return nil, err
	}

	return &tournament, nil
}

func (r *TournamentRepository) CreateTournament(tournament *models.Tournament) error {
	translateNameToSlug(tournament)
	query := "INSERT INTO tournaments (name, slug, num_entrants) VALUES ($1, $2, $3) RETURNING id"
	log.Println(query)
	err := r.DB.QueryRow(query, tournament.Name, tournament.Slug, tournament.NumEntrants).Scan(&tournament.ID)
	if err != nil {
		log.Printf("Error creating tournament: %v\n", err)
		return err
	}
	return nil
}

func translateNameToSlug(tournament *models.Tournament) {
	tournament.Slug = strings.ToLower(strings.ReplaceAll(tournament.Name, " ", "-"))
}
