package pokeapi

import (
	"damagecalculator/infra/index"

	"fmt"
	"strings"

	"github.com/mtslzr/pokeapi-go"
)

/*
	進化先が存在するか検索
	しんかのきせき用
*/
type (
	Evolution interface {
		HasEvolution(name string) bool
	}
)

func NewEvolution() Evolution {
	return &evolution{}
}

type evolution struct {
}

func (ev *evolution) HasEvolution(name string) bool {
	id, err := index.Index(name)
	if err != nil {
		return false
	}
	s, err := pokeapi.PokemonSpecies(id)
	if err != nil {
		return false
	}
	// EvolutionChainのID取得
	url := strings.ReplaceAll(s.EvolutionChain.URL, "/", " ")
	var version string
	var evolveId string
	_, err = fmt.Sscanf(url, "https:  pokeapi.co api %s evolution-chain %s \n", &version, &evolveId)
	if err != nil {
		return false
	}
	chain, err := pokeapi.EvolutionChain(evolveId)
	if err != nil {
		return false
	}
	if chain.Chain.Species.Name == id {
		return len(chain.Chain.EvolvesTo) > 0
	}
	for _, c := range chain.Chain.EvolvesTo {
		if c.Species.Name == id {
			return len(c.EvolvesTo) > 0
		}
		for _, cc := range c.EvolvesTo {
			if cc.Species.Name == id {
				return false
			}
		}
	}
	return false
}
