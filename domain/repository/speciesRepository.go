package repository

import (
	"damagecalculator/domain/species"
)

// 実体はinfra層に
// 目標：Situation作成

/*
	種族一覧

	種族	(タイプ、種族値)
		技一覧
		特性一覧

	もちもの一覧
*/
type (
	RepositoryFactory interface {
		Species() species.SpeciesRepository
		Moves() MovesRepository
	}
)
