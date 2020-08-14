package ability

import (
	"damagecalculator/domain/status"
	"damagecalculator/domain/types"
)

// 能力に補正を掛ける
// 条件指定なし

type statusCorrector struct {
	ability
	ty                 []types.Type
	at, df, sa, sd, sp float64
	wt                 float64
}

func (s *statusCorrector) CorrectStatus(situationChecker) *status.StatsCorrectors {
	c := status.NewStatsCorrectors()

	if s.ty != nil {
		c.Types(s.ty)
	}
	if s.at > 0 {
		c.Attack(s.at)
	}
	if s.df > 0 {
		c.Defense(s.df)
	}
	if s.sa > 0 {
		c.SpAttack(s.sa)
	}
	if s.sd > 0 {
		c.SpDefense(s.sd)
	}
	if s.sp > 0 {
		c.Speed(s.sp)
	}
	if s.wt > 0 {
		c.Weight(s.wt)
	}
	return c
}
