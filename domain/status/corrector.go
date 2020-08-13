package status

import "damagecalculator/domain/fixed"

type StatsCorrectors struct {
	at, df, sa, sd, sp fixed.FixPN
	weight             float64
}

/*
	能力値に補正を掛ける
	チェーンメソッドで設定する

	c := NewStatsCorrectors().Attack(1.5)
*/
func NewStatsCorrectors() *StatsCorrectors {
	at, _ := fixed.NewFixPN(1.0, fixed.Drop4Pick5)
	df, _ := fixed.NewFixPN(1.0, fixed.Drop4Pick5)
	sa, _ := fixed.NewFixPN(1.0, fixed.Drop4Pick5)
	sd, _ := fixed.NewFixPN(1.0, fixed.Drop4Pick5)
	sp, _ := fixed.NewFixPN(1.0, fixed.Drop4Pick5)
	weight := 1.0
	return &StatsCorrectors{
		at:     at,
		df:     df,
		sa:     sa,
		sd:     sd,
		sp:     sp,
		weight: weight,
	}
}

func (s *StatsCorrectors) Correct(at, df, sa, sd, sp uint) (a, b, c, d, e uint) {
	a = s.at.Correct(at)
	b = s.df.Correct(df)
	c = s.sa.Correct(sa)
	d = s.sd.Correct(sd)
	e = s.sp.Correct(sp)
	return
}

func (s *StatsCorrectors) CorrectWeight(w float64) Weight {
	return NewWeight(w * s.weight)
}

// TODO:immutable
func (s *StatsCorrectors) Attack(m float64) *StatsCorrectors {
	s.at, _ = fixed.NewFixPN(m, fixed.Drop4Pick5)
	return s
}
func (s *StatsCorrectors) Defense(m float64) *StatsCorrectors {
	s.df, _ = fixed.NewFixPN(m, fixed.Drop4Pick5)
	return s
}
func (s *StatsCorrectors) SpAttack(m float64) *StatsCorrectors {
	s.sa, _ = fixed.NewFixPN(m, fixed.Drop4Pick5)
	return s
}
func (s *StatsCorrectors) SpDefense(m float64) *StatsCorrectors {
	s.sd, _ = fixed.NewFixPN(m, fixed.Drop4Pick5)
	return s
}
func (s *StatsCorrectors) Speed(m float64) *StatsCorrectors {
	s.sp, _ = fixed.NewFixPN(m, fixed.Drop4Pick5)
	return s
}

func (s *StatsCorrectors) Weight(m float64) *StatsCorrectors {
	s.weight = m
	return s
}
