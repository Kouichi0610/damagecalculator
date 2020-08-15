package correct

import (
	"damagecalculator/domain/ability/situation"
	"damagecalculator/domain/corrector"
	"damagecalculator/domain/skill"
)

// 特定のわざで威力x倍
type actionAttack struct {
	ac skill.Action
	sc float64
}

func (s *actionAttack) Correct(isAttacker bool, st situation.SituationChecker) corrector.Corrector {
	if !isAttacker {
		return nil
	}
	if st.SkillAction() == s.ac {
		return corrector.NewPower(s.sc)
	}
	return nil
}
