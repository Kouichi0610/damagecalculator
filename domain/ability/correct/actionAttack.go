package correct

import (
	"damagecalculator/domain/ability/situation"
	"damagecalculator/domain/corrector"
	"damagecalculator/domain/move/attribute"
)

// 特定のわざで威力x倍
type actionAttack struct {
	ac attribute.Action
	sc float64
}

func (s *actionAttack) Correct(isAttacker bool, st situation.SituationChecker) corrector.Corrector {
	if !isAttacker {
		return nil
	}
	attr := st.MoveAttribute()
	if attr.HasAction(s.ac) {
		return corrector.NewPower(s.sc)
	}
	return nil
}
