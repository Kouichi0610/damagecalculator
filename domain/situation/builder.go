package situation

import (
	"damagecalculator/domain/ability"
	"damagecalculator/domain/condition"
	"damagecalculator/domain/field"
	"damagecalculator/domain/individuals"
	"damagecalculator/domain/item"
	"damagecalculator/domain/move"
	"damagecalculator/domain/species"
	"damagecalculator/domain/stats"
	"fmt"
)

/*
	TODO:factory.go を置き換え、こっちに統一
*/
type (
	PokeParams struct {
		Name        string
		Nature      string
		Individuals string
		BasePoints  []uint
		Ranks       []int
		Ability     string
		Item        string
		Condition   string
	}
	FieldCondition struct {
		Weather      string
		Field        string
		HasReflector bool // リフレクター、ひかりのかべ
	}
	Builder interface {
		ToSituation(level stats.Level, attacker, defender *PokeParams, move string, cd *FieldCondition) (SituationChecker, error)
	}
)

func (p *PokeParams) AttackersInfo() string {
	res := p.Name
	res += " " + p.Nature
	if p.BasePoints[1] > 0 {
		res += fmt.Sprintf(" 攻撃%d", p.BasePoints[1])
	}
	if p.BasePoints[3] > 0 {
		res += fmt.Sprintf(" 特攻%d", p.BasePoints[3])
	}
	return res
}
func (p *PokeParams) DefendersInfo() string {
	res := p.Name
	res += " " + p.Nature
	if p.BasePoints[0] > 0 {
		res += fmt.Sprintf(" HP%d", p.BasePoints[0])
	}
	if p.BasePoints[2] > 0 {
		res += fmt.Sprintf(" 防御%d", p.BasePoints[2])
	}
	if p.BasePoints[4] > 0 {
		res += fmt.Sprintf(" 特防%d", p.BasePoints[4])
	}
	return res
}

func NewBuilder(sp species.Repository, ab ability.Repository, mv move.Repository, it item.Repository) Builder {
	return &builder{
		sp: sp,
		ab: ab,
		mv: mv,
		it: it,
	}
}

type builder struct {
	sp species.Repository
	ab ability.Repository
	mv move.Repository
	it item.Repository
}

func (b *builder) ToSituation(level stats.Level, attacker, defender *PokeParams, move string, cd *FieldCondition) (SituationChecker, error) {
	d := &SituationData{
		Move: move,
		Attacker: PokeData{
			Name:        attacker.Name,
			Level:       uint(level),
			Individuals: individuals.ToIndividualType(attacker.Individuals),
			BasePoints:  toBasePoints(attacker.BasePoints),
			Ranks:       Ranks{0, 0, 0, 0, 0},
			Ability:     attacker.Ability,
			Item:        attacker.Item,
			Nature:      stats.NameToNatureType(attacker.Nature),
			Condition:   condition.FromString(attacker.Condition),
		},
		Defender: PokeData{
			Name:        defender.Name,
			Level:       uint(level),
			Individuals: individuals.ToIndividualType(defender.Individuals),
			BasePoints:  toBasePoints(defender.BasePoints),
			Ranks:       Ranks{0, 0, 0, 0, 0},
			Ability:     defender.Ability,
			Item:        defender.Item,
			Nature:      stats.NameToNatureType(defender.Nature),
			Condition:   condition.FromString(defender.Condition),
		},
		Weather:       field.ToWeather(cd.Weather),
		Field:         field.ToField(cd.Field),
		IsCritical:    false,
		IsReflector:   cd.HasReflector,
		IsLightScreen: cd.HasReflector,
	}
	st, err := d.Create(b.mv, b.sp, b.ab, b.it)
	if err != nil {
		return nil, err
	}
	return st, nil
}

func toBasePoints(p []uint) BasePoints {
	return BasePoints{
		HP:        p[0],
		Attack:    p[1],
		Defense:   p[2],
		SpAttack:  p[3],
		SpDefense: p[4],
		Speed:     p[5],
	}
}
