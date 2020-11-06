// ダメージ計算API
package damage

import (
	"damagecalculator/domain/ability"
	"damagecalculator/domain/condition"
	"damagecalculator/domain/damage"
	DOMAIN "damagecalculator/domain/damage"
	"damagecalculator/domain/field"
	"damagecalculator/domain/item"
	"damagecalculator/domain/move"
	"damagecalculator/domain/situation"
	"damagecalculator/domain/species"
	"damagecalculator/domain/stats"
)

type PokeParams struct {
	Name        string
	Nature      string
	Individuals string
	BasePoints  []int
	Ranks       []int
	Ability     string
	Item        string
	Condition   string
}
type FieldCondition struct {
	Weather      string
	Field        string
	HasReflector bool // リフレクター、ひかりのかべ
}

//
type (
	DamageRateService interface {
		Calculate(level int, attacker *PokeParams, defender *PokeParams, move string, conditions *FieldCondition) (DOMAIN.Damages, DOMAIN.DamageRate, error)
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

func toBasePoints(p []int) situation.BasePoints {
	return situation.BasePoints{
		HP:        uint(p[0]),
		Attack:    uint(p[1]),
		Defense:   uint(p[2]),
		SpAttack:  uint(p[3]),
		SpDefense: uint(p[4]),
		Speed:     uint(p[5]),
	}
}

func (s *service) Calculate(level int, attacker *PokeParams, defender *PokeParams, move string, conditions *FieldCondition) (DOMAIN.Damages, DOMAIN.DamageRate, error) {
	d := &situation.SituationData{
		Move: move,
		Attacker: situation.PokeData{
			Name:        attacker.Name,
			Level:       uint(level),
			Individuals: stats.ToIndividualType(attacker.Individuals),
			BasePoints:  toBasePoints(attacker.BasePoints),
			Ranks:       situation.Ranks{0, 0, 0, 0, 0},
			Ability:     attacker.Ability,
			Item:        attacker.Item,
			Nature:      stats.NameToNatureType(attacker.Nature),
			Condition:   condition.FromString(attacker.Condition),
		},
		Defender: situation.PokeData{
			Name:        defender.Name,
			Level:       uint(level),
			Individuals: stats.ToIndividualType(defender.Individuals),
			BasePoints:  toBasePoints(defender.BasePoints),
			Ranks:       situation.Ranks{0, 0, 0, 0, 0},
			Ability:     defender.Ability,
			Item:        defender.Item,
			Nature:      stats.NameToNatureType(defender.Nature),
			Condition:   condition.FromString(defender.Condition),
		},
		Weather:       field.ToWeather(conditions.Weather),
		Field:         field.ToField(conditions.Field),
		IsCritical:    false,
		IsReflector:   conditions.HasReflector,
		IsLightScreen: conditions.HasReflector,
	}
	st, err := d.Create(s.mv, s.sp, s.ab, s.it)
	if err != nil {
		return nil, nil, err
	}
	dmg := damage.NewDamageCalculator()
	damages, rates := dmg.CreateDamage(st)

	return damages, rates, nil
}
