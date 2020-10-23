// 種族値合計による絞り込み
package filter

import (
	"damagecalculator/domain/species"
)

type TotalFilter interface {
	Filter(s *species.Species) FilterResult
}

func createNoLimitTotalFilter() TotalFilter {
	return createTotalFilter(0)
}

func createTotalFilter(total uint) TotalFilter {
	return &totalFilter{total: total}
}

type totalFilter struct {
	total uint
}

func (f *totalFilter) Filter(s *species.Species) FilterResult {
	t := s.HP + s.Attack + s.Defense + s.SpAttack + s.SpDefense + s.Speed
	if t >= f.total {
		return Include
	}
	return Exclude
}
