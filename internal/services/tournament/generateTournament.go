package tournament

import (
	"log"

	"github.com/BeauRussell/OpenBracket/internal/db/models"
	"github.com/BeauRussell/OpenBracket/internal/db/repositories"
)

type TournamentService struct {
	EntrantRepo    *repositories.EntrantRepository
	TournamentRepo *repositories.TournamentRepository
	Tournament     *models.Tournament
}

func NewBracketService(entrantRepo *repositories.EntrantRepository, tournamentRepo *repositories.TournamentRepository) *TournamentService {
	return &TournamentService{
		EntrantRepo:    entrantRepo,
		TournamentRepo: tournamentRepo,
	}
}

func (s *TournamentService) GenerateTournament(tournamentName string) (error, int) {
	tournamentStruct := models.Tournament{
		Name: tournamentName,
	}

	err := s.TournamentRepo.CreateTournament(&tournamentStruct)
	if err != nil {
		log.Printf("Failed to create tournament: %v\n", err)
		return err, 0
	}

	s.Tournament = &tournamentStruct

	return nil, tournamentStruct.ID
}

func (s *TournamentService) CreateEntrant(name string) (error, *models.Entrant) {
	entrantStruct := models.Entrant{
		Name:       name,
		Tournament: s.Tournament,
	}
	err := s.EntrantRepo.CreateEntrant(&entrantStruct)
	if err != nil {
		log.Printf("Failed to create entrant: %v", err)
		return err, nil
	}

	return nil, &entrantStruct
}
