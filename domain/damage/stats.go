package damage

// 能力値＋ランク補正
type Stats struct {
	hp, at, df, sa, sd, sp  uint
	atr, dfr, sar, sdr, spr Rank
}

func NewStats(hp, at, df, sa, sd, sp uint) *Stats {
	res := new(Stats)
	res.hp = hp
	res.at = at
	res.df = df
	res.sa = sa
	res.sd = sd
	res.sp = sp

	res.atr = Rank(0)
	res.dfr = Rank(0)
	res.sar = Rank(0)
	res.sdr = Rank(0)
	res.spr = Rank(0)
	return res
}

func (s *Stats) Attack() uint {
	return s.atr.RankedStats(s.at)
}
func (s *Stats) Defense() uint {
	return s.dfr.RankedStats(s.df)
}
func (s *Stats) SpAttack() uint {
	return s.sar.RankedStats(s.sa)
}
func (s *Stats) SpDefense() uint {
	return s.sdr.RankedStats(s.sd)
}
func (s *Stats) Speed() uint {
	return s.spr.RankedStats(s.sp)
}
