package ability

import (
	"damagecalculator/domain/corrector"
	"damagecalculator/domain/skill"
)

// 特定のわざで威力x倍
type actionPowerCorrector struct {
	ability
	ac skill.Action
	sc float64
}

func (s *actionPowerCorrector) Correctors(st situationChecker) []corrector.Corrector {
	if st.SkillAction() == s.ac {
		return []corrector.Corrector{corrector.NewPower(s.sc)}
	}
	return []corrector.Corrector{corrector.NewPower(1.0)}
}
