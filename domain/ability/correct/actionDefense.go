package correct

import (
	"damagecalculator/domain/ability/situation"
	"damagecalculator/domain/corrector"
	"damagecalculator/domain/skill"
)

// 特定のわざのダメージ補正(防御側)
type actionDefense struct {
	ac skill.Action
	sc float64
}

func (s *actionDefense) Correct(isAttacker bool, st situation.SituationChecker) corrector.Corrector {
	if isAttacker {
		return nil
	}
	if st.SkillAction() == s.ac {
		return corrector.NewDamage(s.sc)
	}
	return nil
}
