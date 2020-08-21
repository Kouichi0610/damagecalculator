package pokeapi

import (
	"damagecalculator/domain/move"
	"damagecalculator/domain/repository"
	"damagecalculator/domain/species"
)

const JpName = "ja-Hrkt"

type repositoryFactory struct {
}

func NewRepositoryFactory() repository.RepositoryFactory {
	return new(repositoryFactory)
}

func (repositoryFactory) Species() species.Repository {
	return new(speciesRepository)
}

func (repositoryFactory) Moves() move.Repository {
	return new(movesRepository)
}
