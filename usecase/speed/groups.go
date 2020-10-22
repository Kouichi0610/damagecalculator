package speed

import (
	"damagecalculator/domain/species"
	"sort"
)

// 素早さ種族値が同じポケモンをグループ化
type speedGroup struct {
	species uint     // 種族値
	names   []string // 名前一覧
}

type speedGroups struct {
	list []*speedGroup
}

func (g *speedGroup) createInfos(level uint) SpeedInfos {
	return createInfos(g.names, level, g.species)
}
func (groups *speedGroups) createInfos(level uint) SpeedInfos {
	res := make(SpeedInfos, 0)
	for _, g := range groups.list {
		res = append(res, g.createInfos(level)...)
	}
	return res
}

func newSpeedGroups(s species.Repository) *speedGroups {
	res := new(speedGroups)
	list := totalFilter(s)
	i := 0
	for i < len(list) {
		sp := list[i].Speed
		grp := &speedGroup{
			species: sp,
			names:   make([]string, 0),
		}
		sameSpecies := make([]species.Species, 0)

		for sp == list[i].Speed {
			sameSpecies = append(sameSpecies, list[i])
			i++
			if i == len(list) {
				break
			}
		}
		sort.SliceStable(sameSpecies, func(i, j int) bool { return total(&sameSpecies[i]) > total(&sameSpecies[j]) })
		for j := 0; j < len(sameSpecies); j++ {
			grp.names = append(grp.names, sameSpecies[j].Name)
		}
		res.list = append(res.list, grp)
	}
	return res
}
