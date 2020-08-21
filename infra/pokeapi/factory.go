package pokeapi

import (
	"damagecalculator/domain/move"
	"damagecalculator/domain/species"
)

const JpName = "ja-Hrkt"

func Species() species.Repository {
	return new(speciesRepository)
}
func Moves() move.Repository {
	return new(movesRepository)
}
