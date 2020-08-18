package repository

import (
	"damagecalculator/domain/gender"
	"damagecalculator/domain/stats"
	"damagecalculator/domain/types"
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
		Species() SpeciesRepository
		Moves() MovesRepository
	}

	Species struct {
		Name      string
		Stats     *stats.SpeciesStats
		Types     []types.Type
		Weight    float64
		Gender    gender.Gender
		Abilities []string // 選択可能なとくせい一覧(id intでいいか)
		Moves     []string // 選択可能なわざ一覧
	}

	SpeciesRepository interface {
		Get(name string) (*Species, error)
	}

	DamageClass int
	MoveTarget  int
	// TODO:Skillを統一したい
	Move struct {
		Name     string
		Damage   DamageClass
		Power    int
		Type     types.Type
		Accuracy int
		MinHits  int
		MaxHits  int
		Target   MoveTarget
	}

	MovesRepository interface {
		Get(name string) (*Move, error)
	}
)

const (
	Physical DamageClass = iota
	Special
	Status
)
const (
	Select       MoveTarget = iota // 対象1体
	User                           // 自身
	AllOpponents                   // 全ての相手
	AllOther                       // 自分以外全て
)

/*
gender
必ず異なる
異なる場合もある
必ず同じ
関係なし

*/
