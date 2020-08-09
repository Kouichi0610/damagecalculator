package status

import (
	"damagecalculator/domain/stats"
	"damagecalculator/domain/types"
)

// ポケモンの状態を表す
type Status struct {
	l stats.Level
	t *types.Types
	s *RankedStats
}

func (s *Status) Types() *types.Types {
	return s.t
}

func (s *Status) Level() stats.Level {
	return s.l
}

func (s *Status) HP() *HP {
	return NewHP(s.s.hp)
}

func (s *Status) Attack() *RankedValue {
	return s.s.at
}
func (s *Status) Defense() *RankedValue {
	return s.s.df
}
func (s *Status) SpAttack() *RankedValue {
	return s.s.sa
}
func (s *Status) SpDefense() *RankedValue {
	return s.s.sd
}
func (s *Status) Speed() *RankedValue {
	return s.s.sp
}
