package ability

import (
	"damagecalculator/domain/field"
	"damagecalculator/domain/status"
)

// 特定のフィールドで能力補正
type fieldStatusCorrector struct {
	ability
	statusCorrector
	fl field.Field
}

func (s *fieldStatusCorrector) CorrectStatus(st situationChecker) *status.StatsCorrectors {
	c := status.NewStatsCorrectors()
	if !st.IsField(s.fl) {
		return c
	}
	return s.statusCorrector.CorrectStatus(st)
}
