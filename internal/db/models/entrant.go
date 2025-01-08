package models

type Entrant struct {
	ID         int
	Seed       int
	Name       string
	Standing   int
	Tournament Tournament
}
