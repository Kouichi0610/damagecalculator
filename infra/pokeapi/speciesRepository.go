package pokeapi

import (
	"damagecalculator/domain/gender"
	"damagecalculator/domain/species"
	"damagecalculator/domain/types"
	"damagecalculator/infra/index"

	"github.com/mtslzr/pokeapi-go"
)

type speciesRepository struct {
}

// TODO:Pokeapi経由だと通信時間が
func (s *speciesRepository) GetAll() []species.Species {
	list := index.IndexArray()
	res := make([]species.Species, 0)
	for _, id := range list {
		sp, err := s.Get(id)
		if err != nil {
			continue
		}
		res = append(res, *sp)
	}
	return res
}

func (s *speciesRepository) Get(name string) (*species.Species, error) {
	res := new(species.Species)
	id, err := index.Index(name)
	if err != nil {
		return res, err
	}
	poke, err := pokeapi.Pokemon(id)
	if err != nil {
		return res, err
	}

	res.Name = name

	// 種族値
	var hp, at, df, sa, sd, sp uint
	for _, s := range poke.Stats {
		switch s.Stat.Name {
		case "hp":
			hp = uint(s.BaseStat)
		case "attack":
			at = uint(s.BaseStat)
		case "defense":
			df = uint(s.BaseStat)
		case "special-attack":
			sa = uint(s.BaseStat)
		case "special-defense":
			sd = uint(s.BaseStat)
		case "speed":
			sp = uint(s.BaseStat)
		}
	}
	res.HP = hp
	res.Attack = at
	res.Defense = df
	res.SpAttack = sa
	res.SpDefense = sd
	res.Speed = sp

	res.Weight = float64(poke.Weight) / 10.0

	spData, err := pokeapi.PokemonSpecies(id)
	if err == nil {
		rate := spData.GenderRate
		switch {
		case rate == -1:
			res.Gender = gender.Unknown
		case rate == 0:
			res.Gender = gender.MaleOnly
		case rate == 8:
			res.Gender = gender.FemaleOnly
		default:
			res.Gender = gender.MaleFemale
		}
	}
	// タイプ
	res.Types = make([]types.Type, 0)
	for _, ty := range poke.Types {
		res.Types = append(res.Types, typesMap[ty.Type.Name])
	}

	// とくせい
	res.Abilities = make([]string, 0)
	for _, a := range poke.Abilities {
		ab, err := pokeapi.Ability(a.Ability.Name)
		if err != nil {
			break
		}
		jname := a.Ability.Name
		for _, n := range ab.Names {
			if n.Language.Name == JpName {
				jname = n.Name
				break
			}
		}
		res.Abilities = append(res.Abilities, jname)
	}
	// わざ
	res.Moves = make([]string, 0)
	for _, m := range poke.Moves {
		mv, err := pokeapi.Move(m.Move.Name)
		if err != nil {
			continue
		}
		if mv.DamageClass.Name == "status" {
			continue
		}
		name := m.Move.Name
		res.Moves = append(res.Moves, name)
	}

	return res, nil
}

var typesMap map[string]types.Type = map[string]types.Type{
	"normal":   types.Normal,
	"fire":     types.Fire,
	"water":    types.Water,
	"electric": types.Electric,
	"grass":    types.Grass,
	"ice":      types.Ice,
	"fighting": types.Fighting,
	"poison":   types.Poison,
	"ground":   types.Ground,
	"flying":   types.Flying,
	"psychic":  types.Psychic,
	"bug":      types.Bug,
	"rock":     types.Rock,
	"ghost":    types.Ghost,
	"dragon":   types.Dragon,
	"dark":     types.Dark,
	"steel":    types.Steel,
	"fairy":    types.Fairy,
}

/*
Domain層がInfraを利用してドメインオブジェクトを作る
(Infraはドメインオブジェクトがどうやって成り立つか知らない)
*/
