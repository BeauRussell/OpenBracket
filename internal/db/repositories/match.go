package repositories

import (
	"database/sql"
	"log"

	"github.com/BeauRussell/OpenBracket/internal/db/models"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

type MatchRepository struct {
	DB *sql.DB
}

func NewMatchRepository(db *sql.DB) *MatchRepository {
	return &MatchRepository{DB: db}
}

func (r *MatchRepository) CreateMatch(matchID int, tournamentID int) (error, *models.Match) {
	var match models.Match
	query := `
	INSERT INTO matches (match_id, tournament_id) 
	VALUES (?, ?) 
	RETURNING id, match_id, tournament_id`

	err := r.DB.QueryRow(query, matchID, tournamentID).Scan(
		&match.ID,
		&match.MatchID,
		&match.Tournament.ID,
	)
	if err != nil {
		log.Printf("Error creating match: %v\n", err)
		return err, nil
	}

	return nil, &match
}
