// ダメージ計算API
package damage

import (
	"damagecalculator/domain/ability"
	"damagecalculator/domain/damage"
	DOMAIN "damagecalculator/domain/damage"
	"damagecalculator/domain/item"
	"damagecalculator/domain/move"
	"damagecalculator/domain/situation"
	"damagecalculator/domain/species"
	"damagecalculator/domain/stats"
)

//
type (
	// TODO:situation.PokeDataの変換クラスを作成
	DamageRateService interface {
		Calculate(level stats.Level, attacker, defender *situation.PokeParams, move string, conditions *situation.FieldCondition) (DOMAIN.Damages, DOMAIN.DamageRate, error)
	}
)

type service struct {
	sp species.Repository
	mv move.Repository
	ab ability.Repository
	it item.Repository
}

func NewService(sp species.Repository, mv move.Repository, ab ability.Repository, it item.Repository) DamageRateService {
	return &service{
		sp: sp,
		mv: mv,
		ab: ab,
		it: it,
	}
}

func (s *service) Calculate(level stats.Level, attacker, defender *situation.PokeParams, move string, conditions *situation.FieldCondition) (DOMAIN.Damages, DOMAIN.DamageRate, error) {
	builder := situation.NewBuilder(s.sp, s.ab, s.mv, s.it)
	st, err := builder.ToSituation(level, attacker, defender, move, conditions)
	if err != nil {
		return nil, nil, err
	}
	dmg := damage.NewDamageCalculator()
	damages, rates := dmg.CreateDamage(st)

	return damages, rates, nil
}
