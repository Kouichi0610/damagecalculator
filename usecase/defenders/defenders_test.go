package defenders

import (
	"damagecalculator/domain/damage"
	"damagecalculator/domain/situation"
	"damagecalculator/domain/stats"
	"damagecalculator/infra/local"
	"testing"
)

func Test_Defenders(t *testing.T) {
	df := NewService(local.Names(), local.Species(), local.Move(), local.Ability(), local.Item())
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
		IsCritical:   false,
	}

	res := df.Create(stats.Level(50), attacker, move, fields)
	if res == nil {
		t.Error()
	}
	if len(res) == 0 {
		t.Error()
	}
}

func Test_Result(t *testing.T) {
	dmgs := damage.NewDamages([]uint{3, 4, 6})
	rates := damage.NewDamageRate(30, dmgs)
	r := newResult("sample", dmgs, rates)

	if r.Target() != "sample" {
		t.Error()
	}
	if r.DamageMin() != 3 {
		t.Error()
	}
	if r.DamageMax() != 6 {
		t.Error()
	}
	if r.RateMin() != 10.0 {
		t.Errorf("%f", r.RateMin())
	}
	if r.RateMax() != 20.0 {
		t.Errorf("%f", r.RateMax())
	}
	if r.DetermineCount() != 10 {
		t.Errorf("%d", r.DetermineCount())
	}
}
