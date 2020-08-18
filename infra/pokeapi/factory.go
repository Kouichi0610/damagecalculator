package pokeapi

import (
	"damagecalculator/domain/repository"
)

const JpName = "ja-Hrkt"

type repositoryFactory struct {
}

func NewRepositoryFactory() repository.RepositoryFactory {
	return new(repositoryFactory)
}

func (repositoryFactory) Species() repository.SpeciesRepository {
	return new(speciesRepository)
}

func (repositoryFactory) Moves() repository.MovesRepository {
	return new(movesRepository)
}
