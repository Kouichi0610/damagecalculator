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

// 補正を加えたステータスを生成
func (s *StatsCorrectors) Create(st StatusChecker) StatusChecker {
	res := new(Status)
	res.l = st.Level()
	res.t = st.Types()

	hp := st.HP().value
	at := s.at.Correct(st.Attack().v)
	atr := int(st.Attack().r)
	df := s.df.Correct(st.Defense().v)
	dfr := int(st.Defense().r)
	sa := s.sa.Correct(st.SpAttack().v)
	sar := int(st.SpAttack().r)
	sd := s.sd.Correct(st.SpDefense().v)
	sdr := int(st.SpDefense().r)
	sp := s.sp.Correct(st.SpDefense().v)
	spr := int(st.Speed().r)
	res.s = NewRankedStats(hp, at, df, sa, sd, sp, atr, dfr, sar, sdr, spr)

	return res
}
