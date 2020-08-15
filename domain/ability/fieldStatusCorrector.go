package ability

import (
	"damagecalculator/domain/ability/situation"
	"damagecalculator/domain/field"
	"damagecalculator/domain/status"
)

// 特定のフィールドで能力補正
type fieldStatusCorrector struct {
	ability
	statusCorrector
	fl field.Field
}

func (s *fieldStatusCorrector) CorrectStatus(st situation.SituationChecker) *status.StatsCorrectors {
	c := status.NewStatsCorrectors()
	if !st.IsField(s.fl) {
		return c
	}
	return s.statusCorrector.CorrectStatus(st)
}
