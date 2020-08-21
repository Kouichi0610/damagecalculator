package repository

import (
	"damagecalculator/domain/move"
	"damagecalculator/domain/species"
)

type (
	RepositoryFactory interface {
		Species() species.Repository
		Moves() move.Repository
	}
)
