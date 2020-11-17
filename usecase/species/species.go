/*
	種族に関する情報

	args name, nature(?)

	・種族値
	・能力値計算式
	・タイプ

	・おもさ

	・計算式一式(なので性格も渡す)

	// ダメージ計算関連で
	・しんかのきせきの有効

	・わざ一覧
	・とくせい一覧
*/
package species

import (
	"damagecalculator/domain/species"
	"damagecalculator/domain/types"
)

type (
	Species interface {
		Name() string
		Types() []string
		Species() []uint
		Weight() float64
		Abilities() []string
		Moves() []string
	}

	Loader interface {
		Get(name string) (Species, error)
	}
)

func NewLoader(rp species.Repository) Loader {
	return &loader{
		rp: rp,
	}
}

type speciesImpl struct {
	name      string
	types     []string
	species   []uint
	weight    float64
	abilities []string
	moves     []string
}

func (s *speciesImpl) Name() string {
	return s.name
}
func (s *speciesImpl) Types() []string {
	return s.types
}
func (s *speciesImpl) Species() []uint {
	return s.species
}
func (s *speciesImpl) Weight() float64 {
	return s.weight
}
func (s *speciesImpl) Abilities() []string {
	return s.abilities
}
func (s *speciesImpl) Moves() []string {
	return s.moves
}

type loader struct {
	rp species.Repository
}

func (g *loader) Get(name string) (Species, error) {
	s, err := g.rp.Get(name)
	if err != nil {
		return nil, err
	}

	ty := types.NewTypes(s.Types...)

	sp := []uint{s.HP, s.Attack, s.Defense, s.SpAttack, s.SpDefense, s.Speed}

	res := &speciesImpl{
		name:      s.Name,
		types:     ty.Strings(),
		weight:    s.Weight,
		abilities: s.Abilities,
		species:   sp,
		moves:     s.Moves,
	}

	return res, nil
}
