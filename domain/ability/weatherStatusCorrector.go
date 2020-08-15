package ability

import (
	"damagecalculator/domain/ability/situation"
	"damagecalculator/domain/field"
	"damagecalculator/domain/status"
)

// 特定の天候で能力補正
type weatherStatusCorrector struct {
	ability
	statusCorrector
	wt field.Weather
}

func (s *weatherStatusCorrector) CorrectStatus(st situation.SituationChecker) *status.StatsCorrectors {
	c := status.NewStatsCorrectors()
	if !st.IsWeather(s.wt) {
		return c
	}
	return s.statusCorrector.CorrectStatus(st)
}
