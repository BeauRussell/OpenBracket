package models

type Match struct {
	ID         int
	MatchID    int
	Tournament Tournament
	Entrants   [2]Entrant
	Winner     Entrant
}
