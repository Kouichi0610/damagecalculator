package damage

import (
	"damagecalculator/infra/local"
	"testing"
)

func Test_DamgeRateService(t *testing.T) {
	sp := local.Species()
	mv := local.Move()
	ab := local.Ability()
	it := local.Item()
	sv := NewService(sp, mv, ab, it)

	fields := &FieldCondition{
		Weather:      "なし",
		Field:        "なし",
		HasReflector: false,
	}
	attacker := &PokeParams{
		Name:        "ピカチュウ",
		Individuals: "Max",
		BasePoints:  []int{6, 252, 0, 0, 0, 252},
		Ranks:       []int{0, 0, 0, 0, 0},
		Ability:     "せいでんき",
		Item:        "None",
		Nature:      "のんき",
		Condition:   "なし",
	}
	defender := &PokeParams{
		Name:        "シェルダー",
		Individuals: "Max",
		BasePoints:  []int{252, 0, 252, 0, 0, 0},
		Ranks:       []int{0, 0, 0, 0, 0},
		Ability:     "げきりゅう",
		Item:        "None",
		Condition:   "なし",
	}
	move := "ボルテッカー"

	_, rate, err := sv.Calculate(50, attacker, defender, move, fields)
	if err != nil {
		t.Error()
	}
	if rate == nil {
		t.Error()
		return
	}
	if rate.RateMin() != 72.3 {
		t.Errorf("RateMin:%f", rate.RateMin())
	}
	if rate.RateMax() != 85.4 {
		t.Errorf("RateMax:%f", rate.RateMax())
	}

	// TODO:ちょっとずれてるのでダメージ計算式確認

	//t.Errorf("Damage:%s", damages.String())
	//t.Errorf("%s", rate.String())

}
