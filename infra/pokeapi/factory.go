package pokeapi

import (
	"damagecalculator/domain/move"
	"damagecalculator/domain/pokenames"
	"damagecalculator/domain/species"
)

const JpName = "ja-Hrkt"

func Names() pokenames.Repository {
	return new(namesRepository)
}

func Species() species.Repository {
	return new(speciesRepository)
}
func Moves() move.Repository {
	return new(movesRepository)
}
