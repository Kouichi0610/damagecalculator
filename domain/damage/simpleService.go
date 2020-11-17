package damage

import (
	"damagecalculator/domain/ability"
	"damagecalculator/domain/item"
	"damagecalculator/domain/move"
	"damagecalculator/domain/situation"
	"damagecalculator/domain/species"
	"damagecalculator/domain/stats"
)

/*
	ダメージの目安を取得するための簡易サービスクラス
*/
type (
	SimpleService interface {
		Calculate(attacker, defender, aAbility, dAbility, move string) uint
	}
)

func NewSimpleService(sp species.Repository, mv move.Repository, ab ability.Repository, it item.Repository) SimpleService {
	return &simple{
		sp: sp,
		mv: mv,
		ab: ab,
		it: it,
	}
}

func (s *simple) Calculate(attacker, defender, aAbility, dAbility, move string) uint {
	builder := situation.NewBuilder(s.sp, s.ab, s.mv, s.it)
	level := stats.NewLevel(50)

	at := &situation.PokeParams{
		Name:        attacker,
		Nature:      "てれや",
		Individuals: "Max",
		BasePoints:  []uint{0, 0, 0, 0, 0, 0},
		Ranks:       []int{0, 0, 0, 0, 0},
		Ability:     aAbility,
		Item:        "",
		Condition:   "",
	}
	df := &situation.PokeParams{
		Name:        defender,
		Nature:      "てれや",
		Individuals: "Max",
		BasePoints:  []uint{0, 0, 0, 0, 0, 0},
		Ranks:       []int{0, 0, 0, 0, 0},
		Ability:     dAbility,
		Item:        "",
		Condition:   "",
	}
	cd := &situation.FieldCondition{
		Weather:      "なし",
		Field:        "なし",
		HasReflector: false,
	}
	st, err := builder.ToSituation(level, at, df, move, cd)
	if err != nil {
		return 0
	}
	dmg := NewDamageCalculator()
	damages, _ := dmg.CreateDamage(st)
	return damages.Min()
}

type simple struct {
	sp species.Repository
	mv move.Repository
	ab ability.Repository
	it item.Repository
}
