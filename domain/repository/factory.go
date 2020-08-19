package repository

import (
	"damagecalculator/domain/skill"
	"damagecalculator/domain/species"
)

type (
	RepositoryFactory interface {
		Species() species.Repository
		Moves() skill.Repository
	}
)
