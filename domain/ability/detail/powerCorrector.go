package detail

import (
	"damagecalculator/domain/ability/correct"
	"damagecalculator/domain/ability/situation"
	"damagecalculator/domain/corrector"
)

// ダメージ、威力補正をまとめたもの
type powerCorrector struct {
	abilityImpl
	c []correct.PowerCorrector
}

func (s *powerCorrector) Correctors(st situation.SituationChecker) []corrector.Corrector {
	res := make([]corrector.Corrector, 0)
	for _, c := range s.c {
		cr := c.Correct(s.isAttacker, st)
		if cr == nil {
			continue
		}
		res = append(res, cr)
	}
	if len(res) == 0 {
		res = []corrector.Corrector{corrector.NewPower(1.0)}
	}
	return res
}
