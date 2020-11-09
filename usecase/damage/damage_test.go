package damage

import (
	"damagecalculator/domain/situation"
	"damagecalculator/domain/stats"
	"damagecalculator/infra/local"
	"testing"
)

func Test_Defenders(t *testing.T) {
	df := NewVirtualDefendersService(local.Names(), local.Species(), local.Move(), local.Ability(), local.Item())
	attacker := &situation.PokeParams{
		Name:        "ピカチュウ",
		Individuals: "Max",
		BasePoints:  []uint{6, 252, 0, 0, 0, 252},
		Ranks:       []int{0, 0, 0, 0, 0},
		Ability:     "せいでんき",
		Item:        "None",
		Nature:      "いじっぱり",
		Condition:   "なし",
	}
	move := "かみなりパンチ"
	fields := &situation.FieldCondition{
		Weather:      "なし",
		Field:        "なし",
		HasReflector: false,
	}

	df.Create(stats.Level(50), attacker, move, fields)

}
