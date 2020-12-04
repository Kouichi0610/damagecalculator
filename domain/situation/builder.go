package situation

import (
	"damagecalculator/domain/ability"
	"damagecalculator/domain/basepoints"
	_ "damagecalculator/domain/condition"
	"damagecalculator/domain/field"
	"damagecalculator/domain/individuals"
	"damagecalculator/domain/item"
	"damagecalculator/domain/move"
	"damagecalculator/domain/species"
	"damagecalculator/domain/stats"
	"damagecalculator/domain/status"
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
		IsCritical   bool
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
	mv, err := b.move(move)
	if err != nil {
		return nil, err
	}
	wt := field.ToWeather(cd.Weather)
	fl := field.ToField(cd.Field)
	at, err := attacker.toStatus(b.sp, level)
	if err != nil {
		return nil, err
	}
	df, err := defender.toStatus(b.sp, level)
	if err != nil {
		return nil, err
	}

	return &situation{
		attacker:      at,
		defender:      df,
		move:          mv,
		fields:        field.NewFields(fl, wt),
		attackersItem: b.atItem(attacker.Item),
		defendersItem: b.dfItem(defender.Item),
		abilities:     b.ab.Get(attacker.Ability, defender.Ability),
		isCritical:    cd.IsCritical,
	}, nil
}

func (b *builder) move(move string) (move.Move, error) {
	factory, err := b.mv.Get(move)
	if err != nil {
		return nil, err
	}
	res, err := factory.Create()
	return res, err
}

func (b *builder) atItem(name string) item.Item {
	return b.it.Get(name, true)
}

func (b *builder) dfItem(name string) item.Item {
	return b.it.Get(name, false)
}

func (p *PokeParams) toStatus(rp species.Repository, level stats.Level) (status.StatusChecker, error) {
	sp, err := rp.Get(p.Name)
	if err != nil {
		return nil, err
	}
	s := stats.NewSpeciesStats(sp.HP, sp.Attack, sp.Defense, sp.SpAttack, sp.SpDefense, sp.Speed)
	n := stats.NameToNature(p.Nature)
	i := individuals.ToIndividuals(p.Individuals)
	b := basePoints.New(p.BasePoints[0], p.BasePoints[1], p.BasePoints[2], p.BasePoints[3], p.BasePoints[4], p.BasePoints[5])
	stats := stats.NewStats(level, s, i, b, n)

	data := &status.StatusData{
		Lv:    uint(level),
		Types: sp.Types,

		HP:        stats.HP,
		Attack:    stats.Attack,
		Defense:   stats.Defense,
		SpAttack:  stats.SpAttack,
		SpDefense: stats.SpDefense,
		Speed:     stats.Speed,

		AttackRank:    p.Ranks[0],
		DefenseRank:   p.Ranks[1],
		SpAttackRank:  p.Ranks[2],
		SpDefenseRank: p.Ranks[3],
		SpeedRank:     p.Ranks[4],
		Weight:        sp.Weight,
	}
	return data.Create(), nil
}
