package status

import (
	"damagecalculator/domain/stats"
	"damagecalculator/domain/types"
	"fmt"
)

// ポケモンの状態を表す
type Status struct {
	l stats.Level
	t *types.Types
	s *RankedStats
	w Weight
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
func (s *Status) Weight() Weight {
	return s.w
}

func (s *Status) String() string {
	return fmt.Sprintf("Lv:%d %s WT:%f HP:%s AT:%s DF:%s SA:%s SD:%s SP:%s", s.Level(), s.Types().String(), s.Weight(), s.HP(), s.Attack(), s.Defense(), s.SpAttack(), s.SpDefense(), s.Speed())
}
