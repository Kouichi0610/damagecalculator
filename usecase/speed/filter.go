package speed

import (
	"damagecalculator/domain/species"
	"sort"
)

const speciesLower = 460
const speciesUpper = 610

// 候補を減らすために、種族値合計の極端に低いものと高いもの(伝説)を除外
// TODO:ケッキングが除外されるのでどうにかする
func total(s *species.Species) uint {
	return s.HP + s.Attack + s.Defense + s.SpAttack + s.SpDefense + s.Speed
}
func totalFilter(s species.Repository) []species.Species {
	allList := s.GetAll()
	list := make([]species.Species, 0)

	for _, sp := range allList {
		t := total(&sp)
		if t < speciesLower {
			continue
		}
		if t > speciesUpper {
			continue
		}
		list = append(list, sp)
	}
	sort.SliceStable(list, func(i, j int) bool { return list[i].Speed < list[j].Speed })
	return list
}
