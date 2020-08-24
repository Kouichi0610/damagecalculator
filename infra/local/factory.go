package local

import (
	"damagecalculator/domain/ability"
	"damagecalculator/domain/item"
	"damagecalculator/domain/move"
	"damagecalculator/domain/species"
)

func Ability() ability.Repository {
	return new(abilityRepository)
}

func Item() item.Repository {
	return new(itemRepository)
}

func Move() move.Repository {
	return new(movesRepository)
}

func Species() species.Repository {
	return new(speciesRepository)
}
