package damage

import (
	"damagecalculator/domain/situation"
	"damagecalculator/infra/local"
	"testing"
)

func Test_DamgeRateService(t *testing.T) {
	sp := local.Species()
	mv := local.Move()
	ab := local.Ability()
	it := local.Item()
	sv := NewService(sp, mv, ab, it)

	fields := &situation.FieldCondition{
		Weather:      "なし",
		Field:        "なし",
		HasReflector: false,
	}
	attacker := &situation.PokeParams{
		Name:        "ピカチュウ",
		Individuals: "Max",
		BasePoints:  []uint{6, 252, 0, 0, 0, 252},
		Ranks:       []int{0, 0, 0, 0, 0},
		Ability:     "せいでんき",
		Item:        "None",
		Nature:      "のんき",
		Condition:   "なし",
	}
	defender := &situation.PokeParams{
		Name:        "シェルダー",
		Individuals: "Max",
		BasePoints:  []uint{252, 0, 252, 0, 0, 0},
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

	if rate.String() != "72.3% ～ 85.4% 確定数2" {
		t.Errorf("%s", rate.String())
	}

	// TODO:ダメージがちょっとずれてるのでダメージ計算式確認
}
