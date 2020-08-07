package status

// ランク補正付きの能力値
type RankedStats struct {
	hp                 uint
	at, df, sa, sd, sp *RankedValue
}

func NewRankedStats(hp, at, df, sa, sd, sp uint, atr, dfr, sar, sdr, spr int) *RankedStats {
	res := new(RankedStats)
	res.hp = hp
	res.at = NewRankedValue(at, atr)
	res.df = NewRankedValue(df, dfr)
	res.sa = NewRankedValue(sa, sar)
	res.sd = NewRankedValue(sd, sdr)
	res.sp = NewRankedValue(sp, spr)
	return res
}
