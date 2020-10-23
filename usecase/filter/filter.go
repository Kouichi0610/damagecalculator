/*
	種族を絞り込むためのAPI

	タイプ
	種族値合計(ツボツボは？) (仮想敵絞り込みのため)
	しきい値?(全部69以下を除外)threshold
*/
package filter

import (
	"damagecalculator/domain/species"
	"damagecalculator/domain/types"
	"strconv"
)

const (
	Exclude FilterResult = iota
	Include
)

type (
	// 名前、タイプ、種族値
	Result interface {
		Name() string
		Types() string
		HP() uint
		Attack() uint
		Defense() uint
		SpAttack() uint
		SpDefense() uint
		Speed() uint
	}
	Results []Result

	// 種族絞り込み
	SpeciesFilter interface {
		Filter(s species.Repository) Results
	}
)

/*
	ty フィルタータイプ 指定されたタイプ以外を除く
	total 種族値合計

*/
func NewFilter(ty Type, total string) SpeciesFilter {
	t, err := strconv.Atoi(total)
	var totalFilter TotalFilter
	if err != nil {
		totalFilter = createNoLimitTotalFilter()
	} else {
		totalFilter = createTotalFilter(uint(t))
	}

	return &filterImpl{
		ty: createTypeFilter(ty),
		to: totalFilter,
	}
}

func (r Results) Find(name string) Result {
	for _, s := range r {
		if s.Name() == name {
			return s
		}
	}
	return nil
}

type filterImpl struct {
	ty TypeFilter
	to TotalFilter
}

func (f *filterImpl) Filter(s species.Repository) Results {
	res := make(Results, 0)
	list := s.GetAll()
	for _, sp := range list {
		if f.ty.Filter(sp.Types) == Exclude {
			continue
		}
		if f.to.Filter(&sp) == Exclude {
			continue
		}
		res = append(res, newResult(&sp))
	}
	return res
}

func newResult(s *species.Species) Result {
	tt := types.NewTypes(s.Types...)
	typeString := tt.String()

	return &resultImpl{
		name:  s.Name,
		types: typeString,
		hp:    s.HP,
		at:    s.Attack,
		df:    s.Defense,
		sa:    s.SpAttack,
		sd:    s.SpDefense,
		sp:    s.Speed,
	}
}

type resultImpl struct {
	name                   string
	types                  string
	hp, at, df, sa, sd, sp uint
}

func (r *resultImpl) Name() string {
	return r.name
}
func (r *resultImpl) Types() string {
	return r.types
}
func (r *resultImpl) HP() uint {
	return r.hp
}
func (r *resultImpl) Attack() uint {
	return r.at
}
func (r *resultImpl) Defense() uint {
	return r.df
}
func (r *resultImpl) SpAttack() uint {
	return r.sa
}
func (r *resultImpl) SpDefense() uint {
	return r.sd
}
func (r *resultImpl) Speed() uint {
	return r.sp
}
