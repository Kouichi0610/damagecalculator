package ability

import (
	"damagecalculator/domain/corrector"
	"damagecalculator/domain/types"
)

// xxタイプの威力x倍

type typePowerCorrector struct {
	ability
	ty []types.Type
	sc float64
}

func (s *typePowerCorrector) Correctors(st situationChecker) []corrector.Corrector {
	if !s.ability.isAttacker {
		return []corrector.Corrector{corrector.NewPower(1.0)}
	}
	ty := st.SkillTypes()
	for _, t := range s.ty {
		if ty.Has(t) {
			return []corrector.Corrector{corrector.NewPower(s.sc)}
		}
	}
	return []corrector.Corrector{corrector.NewPower(1.0)}
}
