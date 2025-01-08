package tournament

import (
	"log"

	"github.com/BeauRussell/OpenBracket/internal/db/models"
	"github.com/BeauRussell/OpenBracket/internal/db/repositories"
)

type TournamentService struct {
	EntrantRepo    *repositories.EntrantRepository
	TournamentRepo *repositories.TournamentRepository
}

func NewBracketService(entrantRepo *repositories.EntrantRepository, tournamentRepo *repositories.TournamentRepository) *TournamentService {
	return &TournamentService{
		EntrantRepo:    entrantRepo,
		TournamentRepo: tournamentRepo,
	}
}

func (s *TournamentService) GenerateTournament(tournamentName string) error {
	tournamentStruct := models.Tournament{
		Name: tournamentName,
	}

	err := s.TournamentRepo.CreateTournament(&tournamentStruct)
	if err != nil {
		log.Printf("Failed to create tournament: %v\n", err)
		return err
	}

	return nil
}
