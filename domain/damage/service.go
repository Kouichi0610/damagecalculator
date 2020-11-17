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
	ダメージ計算サービスクラス

*/
type (
	DamageService interface {
		Calculate(level stats.Level, attacker, defender *situation.PokeParams, move string, conditions *situation.FieldCondition) (Damages, DamageRate, error)
	}
)

func NewDamageService(sp species.Repository, mv move.Repository, ab ability.Repository, it item.Repository) DamageService {
	return &service{
		sp: sp,
		mv: mv,
		ab: ab,
		it: it,
	}
}

func (s *service) Calculate(level stats.Level, attacker, defender *situation.PokeParams, move string, conditions *situation.FieldCondition) (Damages, DamageRate, error) {
	builder := situation.NewBuilder(s.sp, s.ab, s.mv, s.it)
	st, err := builder.ToSituation(level, attacker, defender, move, conditions)
	if err != nil {
		return nil, nil, err
	}
	dmg := NewDamageCalculator()
	damages, rates := dmg.CreateDamage(st)

	return damages, rates, nil
}

type service struct {
	sp species.Repository
	mv move.Repository
	ab ability.Repository
	it item.Repository
}
