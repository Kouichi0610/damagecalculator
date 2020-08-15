package correct

import (
	"damagecalculator/domain/ability/situation"
	"damagecalculator/domain/corrector"
	"damagecalculator/domain/types"
)

// xxタイプの威力x倍
type typeAttack struct {
	ty []types.Type
	sc float64
}

func (s *typeAttack) Correct(isAttacker bool, st situation.SituationChecker) corrector.Corrector {
	if !isAttacker {
		return nil
	}
	ty := st.SkillTypes()
	for _, t := range s.ty {
		if ty.Has(t) {
			return corrector.NewPower(s.sc)
		}
	}
	return nil
}
