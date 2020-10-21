/*
	速度調整のための仮想敵リストを作成
	TODO:手動でデータベースにしたほうが良いと思う
*/
package speed

import (
	"damagecalculator/domain/species"
	"sort"
)

/*
	TODO:最速補正あり 無振り　+4
	　　　　(65以上)

*/
type (
	// 種族値が同じポケモンをまとめて
	Param interface {
		Info() string // names, 種族値, 基礎ポイント
		Speed() uint  // 素早さ実数
	}
	Params []Param

	ListGenerator interface {
		Generate(s species.Repository) Params
	}
)

func NewGenerator() ListGenerator {
	return new(generatorImpl)
}

type generatorImpl struct {
}

const speciesLower = 460
const speciesUpper = 610

func total(s *species.Species) uint {
	return s.HP + s.Attack + s.Defense + s.SpAttack + s.SpDefense + s.Speed
}

//
// リスト保管
type speedGroups struct {
	list []species.Species
}

func newSpeedGroups(s species.Repository) *speedGroups {
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
	return &speedGroups{list: list}
}
func (s *speedGroups) group(speed uint) []species.Species {
	res := make([]species.Species, 0)
	for _, sp := range s.list {
		if sp.Speed == speed {
			res = append(res, sp)
		}
	}
	return res
}

func (g *generatorImpl) Generate(s species.Repository) Params {
	res := make(Params, 0)
	grp := newSpeedGroups(s)

	for spd := 5; spd <= 250; spd++ {
		list := grp.group(uint(spd))
		if len(list) == 0 {
			continue
		}
		sort.SliceStable(list, func(i, j int) bool { return total(&list[i]) > total(&list[j]) })

		info := ""
		const nameCount = 3
		for i, sp := range list {
			info += sp.Name
			if i == nameCount-1 {
				break
			}
			info += " "
		}

		res = append(res, newParam(uint(spd), info))
	}
	return res
}

type paramImpl struct {
	info  string
	speed uint
}

func newParam(speed uint, info string) Param {
	return &paramImpl{
		info:  info,
		speed: speed,
	}
}

func (p *paramImpl) Info() string {
	return p.info
}
func (p *paramImpl) Speed() uint {
	return p.speed
}
