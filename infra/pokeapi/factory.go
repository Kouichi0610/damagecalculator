package pokeapi

import (
	"damagecalculator/domain/repository"
	"damagecalculator/domain/skill"
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

func (repositoryFactory) Moves() skill.Repository {
	return new(movesRepository)
}
