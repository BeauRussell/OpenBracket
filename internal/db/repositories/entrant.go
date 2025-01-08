package repositories

import (
	"database/sql"
	"log"

	"github.com/BeauRussell/OpenBracket/internal/db/models"
	_ "github.com/mattn/go-sqlite3"
)

type EntrantRepository struct {
	DB *sql.DB
}

func NewEntrantRepository(db *sql.DB) *EntrantRepository {
	return &EntrantRepository{DB: db}
}

func (r *EntrantRepository) GetEntrantById(id int) (*models.Entrant, error) {
	query := "SELECT id, name FROM entrants WHERE id = $1"
	row := r.DB.QueryRow(query, id)

	var entrant models.Entrant
	if err := row.Scan(&entrant.ID, &entrant.Name); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Printf("Error fetching entrant by ID: %v\n", err)
		return nil, err
	}

	return &entrant, nil
}

func (r *EntrantRepository) CreateEntrant(entrant *models.Entrant, tournament *models.Tournament) error {
	query := "INSERT INTO entrants (name, tournament_id) VALUES ($1, $2) RETURNING id"
	err := r.DB.QueryRow(query, entrant.Name, tournament.ID).Scan(&entrant.ID)
	if err != nil {
		log.Printf("Error creating entrant: %v\n", err)
		return err
	}
	return nil
}

func (r *EntrantRepository) GetEntrantsByTournament(tournamentId int) ([]*models.Entrant, error) {
	query := "SELECT * FROM entrants WHERE tournament_id = ?"
	rows, err := r.DB.Query(query, tournamentId)
	if err != nil {
		log.Printf("Error fetching entrants by tournament: %v\n", err)
		return nil, err
	}
	defer rows.Close()
	var entrants []*models.Entrant
	for rows.Next() {
		var entrant models.Entrant
		if err := rows.Scan(&entrant.ID, &entrant.Name, &entrant.Seed); err != nil {
			log.Printf("Failed to scan row to entrant variable: %v\n", err)
			return nil, err
		}
		entrants = append(entrants, &entrant)
	}
	if err = rows.Err(); err != nil {
		log.Printf("Error querying entrants by tournament: %v\n", err)
		return nil, err
	}

	return entrants, nil
}
