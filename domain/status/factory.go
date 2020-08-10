package status

import (
	"damagecalculator/domain/stats"
	"damagecalculator/domain/types"
)

type StatusData struct {
	Lv                                                              uint
	Types                                                           []types.Type
	HP, Attack, Defense, SpAttack, SpDefense, Speed                 uint
	AttackRank, DefenseRank, SpAttackRank, SpDefenseRank, SpeedRank int
}

func (s *StatusData) Create() StatusChecker {
	l := stats.NewLevel(s.Lv)
	t := types.NewTypes(s.Types...)
	st := NewRankedStats(s.HP, s.Attack, s.Defense, s.SpAttack, s.SpDefense, s.Speed, s.AttackRank, s.DefenseRank, s.SpAttackRank, s.SpDefenseRank, s.SpeedRank)
	return &Status{
		l: l,
		t: t,
		s: st,
	}
}
