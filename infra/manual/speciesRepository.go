package manual

import (
	"damagecalculator/domain/gender"
	"damagecalculator/domain/species"
	"damagecalculator/domain/types"
)

type speciesRepository struct {
}

func (s *speciesRepository) Get(name string) (*species.Species, error) {
	return nil, nil
}

var speciesMap map[string]*species.Species = map[string]*species.Species{
	"ツンデツンデ": {
		Name:      "ツンデツンデ",
		HP:        61,
		Attack:    131,
		Defense:   211,
		SpAttack:  53,
		SpDefense: 101,
		Speed:     13,
		Weight:    820.0,
		Gender:    gender.Unknown,
		Types:     []types.Type{types.Rock, types.Steel},
		Abilities: []string{"ビーストブースト"},
		Moves: []string{
			"tackle",
			"take-down",
			"double-edge",
			"rock-throw",
			"earthquake",
			"bide",
			"rock-slide",
			"return",
			"frustration",
			"hidden-power",
			"facade",
			"rock-tomb",
			"rock-blast",
			"gyro-ball",
			"giga-impact",
			"flash-cannon",
			"iron-head",
			"stone-edge",
			"smack-down",
			"round",
			"bulldoze",
			"infestation",
			"brutal-swing",
		},
	},
}
