package species

import (
	"damagecalculator/domain/gender"
	"damagecalculator/domain/stats"
	"damagecalculator/domain/status"
	"damagecalculator/domain/types"
	"fmt"
	"reflect"
	"testing"
)

func Test_Repository(t *testing.T) {
	s := NewStatusFactory(newRepository())
	args := &StatusFactoryArgs{Name: "unknown"}

	_, _, _, _, err := s.Create(args)
	if err == nil {
		t.Error()
	}
	basePoints, err := stats.NewBasePointStats(6, 252, 0, 0, 0, 252)
	args = &StatusFactoryArgs{
		Name:       "ツンデツンデ",
		Level:      100,
		Ranks:      [5]int{1, 2, -1, 4, 6},
		Nature:     stats.Adamant,
		Individual: stats.NewIndividualStats(31, 31, 31, 31, 31, 31),
		BasePoint:  basePoints,
	}
	stats, ab, mv, g, err := s.Create(args)
	if err != nil {
		t.Error()
	}
	if g != gender.Unknown {
		t.Error()
	}
	if !reflect.DeepEqual(ab, []string{"ビーストブースト"}) {
		t.Error()
	}
	if !reflect.DeepEqual(mv, []string{"tackle"}) {
		t.Error()
	}

	if stats.Weight() != 820.0 {
		t.Error()
	}
	if stats.Attack().Value() != 595 {
		t.Errorf("%d", stats.Attack().Value())
	}
	if stats.Defense().Value() != 916 {
		t.Errorf("%d", stats.Defense().Value())
	}
	if stats.SpAttack().Value() != 84 {
		t.Errorf("%d", stats.SpAttack().Value())
	}
	if stats.SpDefense().Value() != 714 {
		t.Errorf("%d", stats.SpDefense().Value())
	}
	if stats.Speed().Value() != 500 {
		t.Errorf("%d", stats.Speed().Value())
	}
	hp := status.NewHP(264)
	if !reflect.DeepEqual(stats.HP(), hp) {
		t.Errorf("%v", stats.HP())
	}
}

func newRepository() Repository {
	return new(speciesRepository)
}

type speciesRepository struct {
}

func (s *speciesRepository) Get(name string) (*Species, error) {
	if name == "unknown" {
		return nil, fmt.Errorf("%s not found.", name)
	}

	return &Species{
		Name:      "ツンデツンデ",
		Stats:     stats.NewSpeciesStats(61, 131, 211, 53, 101, 13),
		Weight:    820.0,
		Gender:    gender.Unknown,
		Types:     []types.Type{types.Rock, types.Steel},
		Abilities: []string{"ビーストブースト"},
		Moves: []string{
			"tackle",
		},
	}, nil
}
