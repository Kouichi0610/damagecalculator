package repository

import (
	"damagecalculator/domain/species"
)

type (
	RepositoryFactory interface {
		Species() species.SpeciesRepository
		Moves() MovesRepository
	}
)
