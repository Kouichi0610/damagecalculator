package pokeapi

import (
	"damagecalculator/domain/repository"
	"damagecalculator/domain/species"
)

const JpName = "ja-Hrkt"

type repositoryFactory struct {
}

func NewRepositoryFactory() repository.RepositoryFactory {
	return new(repositoryFactory)
}

func (repositoryFactory) Species() species.SpeciesRepository {
	return new(speciesRepository)
}

func (repositoryFactory) Moves() repository.MovesRepository {
	return new(movesRepository)
}
