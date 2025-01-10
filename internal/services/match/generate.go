package match

import (
	"log"
	"sync"

	"github.com/BeauRussell/OpenBracket/internal/db/models"
	"github.com/BeauRussell/OpenBracket/internal/db/repositories"
)

type MatchService struct {
	EntrantRepo    *repositories.EntrantRepository
	TournamentRepo *repositories.TournamentRepository
	MatchRepo      *repositories.MatchRepository
}

func NewMatchService(entrantRepo *repositories.EntrantRepository, tournamentRepo *repositories.TournamentRepository, matchRepo *repositories.MatchRepository) *MatchService {
	return &MatchService{
		EntrantRepo:    entrantRepo,
		TournamentRepo: tournamentRepo,
		MatchRepo:      matchRepo,
	}
}

func (s *MatchService) CreateMatches(numMatches int, tournamentID int) (error, *[]models.Match) {
	matches := make([]models.Match, numMatches)

	var wg sync.WaitGroup
	wg.Add(numMatches)
	for i := 0; i < numMatches; i++ {
		go func() {
			defer wg.Done()
			err, match := s.MatchRepo.CreateMatch(i, tournamentID)
			if err != nil {
				log.Printf("Error creating match: %v\n", err)
			}

			matches[i] = *match
		}()
	}

	wg.Wait()

	return nil, &matches
}
