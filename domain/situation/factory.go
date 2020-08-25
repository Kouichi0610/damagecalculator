package situation

import (
	"damagecalculator/domain/ability"
	"damagecalculator/domain/field"
	"damagecalculator/domain/item"
	"damagecalculator/domain/move"
	"damagecalculator/domain/species"
	"damagecalculator/domain/stats"
	"damagecalculator/domain/status"
)

type Individuals struct {
	HP, Attack, Defense, SpAttack, SpDefense, Speed uint
}
type BasePoints struct {
	HP, Attack, Defense, SpAttack, SpDefense, Speed uint
}
type Ranks struct {
	Attack, Defense, SpAttack, SpDefense, Speed int
}

type PokeData struct {
	Name        string
	Level       uint
	Individuals Individuals
	BasePoints  BasePoints
	Ranks       Ranks
	Nature      stats.NatureType
	Ability     string
	Item        string

	// repositories
}

func (p *PokeData) Create(sp species.Repository) (*status.StatusData, error) {
	d, err := sp.Get(p.Name)
	if err != nil {
		return nil, err
	}

	l := stats.NewLevel(p.Level)
	n := stats.NewNature(p.Nature)
	s := stats.NewSpeciesStats(d.HP, d.Attack, d.Defense, d.SpAttack, d.SpDefense, d.Speed)
	i := stats.NewIndividualStats(p.Individuals.HP, p.Individuals.Attack, p.Individuals.Defense, p.Individuals.SpAttack, p.Individuals.SpDefense, p.Individuals.Speed)
	b, err := stats.NewBasePointStats(p.BasePoints.HP, p.BasePoints.Attack, p.BasePoints.Defense, p.BasePoints.SpAttack, p.BasePoints.SpDefense, p.BasePoints.Speed)
	if err != nil {
		return nil, err
	}
	st := stats.NewStats(l, s, i, b, n)

	return &status.StatusData{
		Lv:            p.Level,
		HP:            st.HP,
		Attack:        st.Attack,
		Defense:       st.Defense,
		SpAttack:      st.SpAttack,
		SpDefense:     st.SpDefense,
		Speed:         st.Speed,
		Types:         d.Types,
		AttackRank:    p.Ranks.Attack,
		DefenseRank:   p.Ranks.Defense,
		SpAttackRank:  p.Ranks.SpAttack,
		SpDefenseRank: p.Ranks.SpDefense,
		SpeedRank:     p.Ranks.Speed,
		Weight:        d.Weight,
	}, nil
}

type SituationData struct {
	Move string

	Attacker, Defender PokeData

	Weather field.Weather
	Field   field.Field

	// Condition
	IsCritical    bool
	IsBurn        bool
	IsReflector   bool
	IsLightScreen bool
}

func (s *SituationData) Create(mv move.Repository, sp species.Repository, ab ability.Repository, it item.Repository) (SituationChecker, error) {
	mf, err := mv.Get(s.Move)
	if err != nil {
		return nil, err
	}
	move, err := mf.Create()
	if err != nil {
		return nil, err
	}
	aitem := it.Get(s.Attacker.Item, true)
	ditem := it.Get(s.Defender.Item, false)
	af := ab.Get(s.Attacker.Ability, s.Defender.Ability)

	adata, err := s.Attacker.Create(sp)
	if err != nil {
		return nil, err
	}
	ddata, err := s.Defender.Create(sp)
	if err != nil {
		return nil, err
	}
	return &situation{
		at:         adata.Create(),
		df:         ddata.Create(),
		sk:         move,
		mv:         field.NewFields(s.Field, s.Weather),
		atItem:     aitem,
		dfItem:     ditem,
		abilities:  af,
		isCritical: s.IsCritical,
	}, nil
}
