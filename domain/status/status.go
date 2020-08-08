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

func NewStatus(lv uint, ty []types.Type, hp, at, df, sa, sd, sp uint, atr, dfr, sar, sdr, spr int) StatusChecker {
	l := stats.NewLevel(lv)
	t := types.NewTypes(ty...)
	s := NewRankedStats(hp, at, df, sa, sd, sp, atr, dfr, sar, sdr, spr)

	return &Status{
		l: l,
		t: t,
		s: s,
	}
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
