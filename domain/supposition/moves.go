package supposition

import (
	"damagecalculator/domain/move"
	"damagecalculator/domain/move/category"
	"damagecalculator/domain/species"
)

// TODO:その中で最大火力を(とくせい)

// カテゴリごとに使える技一覧を取得
type (
	MovesService interface {
		Moves(name string, category category.DamageCategory) []string
	}
)

func newMovesService(sp species.Repository, mv move.Repository) MovesService {
	return &mvService{
		sp: sp,
		mv: mv,
	}
}

func (m *mvService) Moves(name string, category category.DamageCategory) []string {
	res := make([]string, 0)
	sp, err := m.sp.Get(name)
	if err != nil {
		return res
	}
	for _, moveName := range sp.Moves {
		move, _ := m.mv.Get(moveName)
		if move.Power < 60 {
			continue
		}
		if move.Category == category {
			res = append(res, moveName)
		}
	}
	return res
}

type mvService struct {
	sp species.Repository
	mv move.Repository
}
