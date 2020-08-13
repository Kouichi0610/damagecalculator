package status

import (
	"damagecalculator/domain/factor"
	"damagecalculator/domain/types"
)

type StatsCorrectors struct {
	ty                 []types.Type
	at, df, sa, sd, sp factor.FixPN
	weight             float64
}

/*
	能力値に補正を掛ける
	チェーンメソッドで設定する

	c := NewStatsCorrectors().Attack(1.5)
*/
func NewStatsCorrectors() *StatsCorrectors {
	at, _ := factor.NewFixPN(1.0, factor.Drop4Pick5)
	df, _ := factor.NewFixPN(1.0, factor.Drop4Pick5)
	sa, _ := factor.NewFixPN(1.0, factor.Drop4Pick5)
	sd, _ := factor.NewFixPN(1.0, factor.Drop4Pick5)
	sp, _ := factor.NewFixPN(1.0, factor.Drop4Pick5)
	weight := 1.0
	return &StatsCorrectors{
		ty:     nil,
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

// TODO:immutableにしておきたい
func (s *StatsCorrectors) Types(ty []types.Type) *StatsCorrectors {
	s.ty = ty
	return s
}
func (s *StatsCorrectors) Attack(m float64) *StatsCorrectors {
	if m < 0 {
		m = 1.0
	}
	s.at, _ = factor.NewFixPN(m, factor.Drop4Pick5)
	return s
}
func (s *StatsCorrectors) Defense(m float64) *StatsCorrectors {
	if m < 0 {
		m = 1.0
	}
	s.df, _ = factor.NewFixPN(m, factor.Drop4Pick5)
	return s
}
func (s *StatsCorrectors) SpAttack(m float64) *StatsCorrectors {
	if m < 0 {
		m = 1.0
	}
	s.sa, _ = factor.NewFixPN(m, factor.Drop4Pick5)
	return s
}
func (s *StatsCorrectors) SpDefense(m float64) *StatsCorrectors {
	if m < 0 {
		m = 1.0
	}
	s.sd, _ = factor.NewFixPN(m, factor.Drop4Pick5)
	return s
}
func (s *StatsCorrectors) Speed(m float64) *StatsCorrectors {
	if m < 0 {
		m = 1.0
	}
	s.sp, _ = factor.NewFixPN(m, factor.Drop4Pick5)
	return s
}

func (s *StatsCorrectors) Weight(m float64) *StatsCorrectors {
	if m < 0 {
		m = 1.0
	}
	s.weight = m
	return s
}
