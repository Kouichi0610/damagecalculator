package attackers

import (
	"damagecalculator/domain/situation"
	"damagecalculator/domain/stats"
	"damagecalculator/infra/local"
	"testing"
)

func Test_Attackers(t *testing.T) {
	at := NewService(local.Names(), local.Species(), local.Move(), local.Ability(), local.Item())
	defender := &situation.PokeParams{
		Name:        "ツボツボ",
		Individuals: "Max",
		BasePoints:  []uint{252, 0, 0, 0, 252, 0},
		Ranks:       []int{0, 0, 0, 0, 0},
		Ability:     "せいでんき",
		Item:        "None",
		Nature:      "いじっぱり",
		Condition:   "なし",
	}
	fields := &situation.FieldCondition{
		Weather:      "なし",
		Field:        "なし",
		HasReflector: false,
	}

	res := at.Create(stats.Level(50), defender, fields)
	if res == nil {
		t.Error()
	}
	if len(res) == 0 {
		t.Error()
	}

	for _, r := range res {
		t.Errorf("%s", r.String())
	}

}
