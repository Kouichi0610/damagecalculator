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
	Weight                                                          float64
}

func (s *StatusData) Create() StatusChecker {
	l := stats.NewLevel(s.Lv)
	t := types.NewTypes(s.Types...)
	st := NewRankedStats(s.HP, s.Attack, s.Defense, s.SpAttack, s.SpDefense, s.Speed, s.AttackRank, s.DefenseRank, s.SpAttackRank, s.SpDefenseRank, s.SpeedRank)
	w := NewWeight(s.Weight)
	return &Status{
		l: l,
		t: t,
		s: st,
		w: w,
	}
}

// 補正を加えたステータスを生成
func (s *StatsCorrectors) Create(st StatusChecker) StatusChecker {
	res := new(Status)
	res.l = st.Level()
	res.t = st.Types()

	at, df, sa, sd, sp := s.Correct(st.Attack().v, st.Defense().v, st.SpAttack().v, st.SpDefense().v, st.Speed().v)

	hp := st.HP().value
	atr := int(st.Attack().r)
	dfr := int(st.Defense().r)
	sar := int(st.SpAttack().r)
	sdr := int(st.SpDefense().r)
	spr := int(st.Speed().r)
	res.s = NewRankedStats(hp, at, df, sa, sd, sp, atr, dfr, sar, sdr, spr)

	return res
}
