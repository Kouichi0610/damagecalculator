package calculator

import (
	"damagecalculator/domain/ability"
	"damagecalculator/domain/item"
	"damagecalculator/domain/move"
	"damagecalculator/domain/pokenames"
	"damagecalculator/domain/situation"
	"damagecalculator/domain/species"
	"damagecalculator/domain/stats"
	"fmt"
)

type IndividualType uint

const (
	AllMax IndividualType = iota
	Slowest
	Weakest
)

// PokeDataを作成するためのインターフェイス
type PokeDataBuilder interface {
	PokeList() Names

	SetName(string)
	AbilityList() Names
	SetAbility(string)
	SetLevel(uint)
	SetNature(stats.NatureType)
	SetBasePoints(hp, at, df, sa, sd, sp uint)
	SetRanks(at, df, sa, sd, sp int)
	SetIndividual(IndividualType)
	ItemList() Names
	SetItem(string)

	Build() (*situation.PokeData, error)
}

type pokeDataBuilder struct {
	name        string
	level       uint
	nature      stats.NatureType
	basePoints  *situation.BasePoints
	individuals *situation.Individuals
	ranks       *situation.Ranks
	ability     string
	item        string

	pk pokenames.Repository
	mv move.Repository
	sp species.Repository
	ab ability.Repository
	it item.Repository
}

func (b *pokeDataBuilder) PokeList() Names {
	return b.pk.Get()
}

func (b *pokeDataBuilder) SetName(name string) {
	b.name = name
}
func (b *pokeDataBuilder) AbilityList() Names {
	sp, err := b.sp.Get(b.name)
	if err != nil {
		return nil
	}
	res := make([]string, 0)
	res = append(res, sp.Abilities...)
	return res
}
func (b *pokeDataBuilder) SetNature(n stats.NatureType) {
	b.nature = n
}
func (b *pokeDataBuilder) SetAbility(name string) {
	list := b.AbilityList()
	if !list.Has(name) {
		return
	}
	b.ability = name
}
func (b *pokeDataBuilder) SetLevel(level uint) {
	b.level = level
}
func (b *pokeDataBuilder) SetBasePoints(hp, at, df, sa, sd, sp uint) {
	b.basePoints = &situation.BasePoints{hp, at, df, sa, sd, sp}
}
func (b *pokeDataBuilder) SetRanks(at, df, sa, sd, sp int) {
	b.ranks = &situation.Ranks{at, df, sa, sd, sp}
}
func (b *pokeDataBuilder) SetIndividual(ty IndividualType) {
	var st *situation.Individuals
	switch ty {
	case AllMax:
		st = &situation.Individuals{31, 31, 31, 31, 31, 31}
	case Slowest:
		st = &situation.Individuals{31, 31, 31, 31, 31, 0}
	case Weakest:
		st = &situation.Individuals{31, 0, 31, 31, 31, 31}
	}
	b.individuals = st
}
func (b *pokeDataBuilder) ItemList() Names {
	return b.it.List()
}
func (b *pokeDataBuilder) SetItem(name string) {
	list := b.ItemList()
	if !list.Has(name) {
		return
	}
	b.item = name
}

func (b *pokeDataBuilder) Build() (*situation.PokeData, error) {
	err := b.errorCheck()
	if err != nil {
		return nil, err
	}
	return &situation.PokeData{
		Name:        b.name,
		Level:       b.level,
		Nature:      b.nature,
		Individuals: *b.individuals,
		BasePoints:  *b.basePoints,
		Ranks:       *b.ranks,
		Ability:     b.ability,
		Item:        b.item,
	}, nil
}

func (b *pokeDataBuilder) errorCheck() error {
	if len(b.name) == 0 {
		return fmt.Errorf("no name")
	}
	if b.level == 0 {
		return fmt.Errorf("level error")
	}
	if b.individuals == nil {
		return fmt.Errorf("no individuals")
	}
	if b.basePoints == nil {
		return fmt.Errorf("no basepoints")
	}
	if b.ranks == nil {
		return fmt.Errorf("no ranks")
	}
	if len(b.ability) == 0 {
		return fmt.Errorf("no ability")
	}
	return nil
}
