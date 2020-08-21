package correct

import (
	"damagecalculator/domain/ability/situation"
	"damagecalculator/domain/corrector"
	"damagecalculator/domain/move/attribute"
)

// 特定のわざのダメージ補正(防御側)
type actionDefense struct {
	at attribute.Action
	sc float64
}

func (s *actionDefense) Correct(isAttacker bool, st situation.SituationChecker) corrector.Corrector {
	if isAttacker {
		return nil
	}
	attr := st.MoveAttribute()
	if attr.HasAction(s.at) {
		return corrector.NewDamage(s.sc)
	}
	return nil
}
