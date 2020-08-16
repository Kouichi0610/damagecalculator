package ability

import (
	"damagecalculator/domain/ability/situation"
	"damagecalculator/domain/field"
	"damagecalculator/domain/status"
	"damagecalculator/domain/types"
)

// てんきや
// 晴れの時ほのおタイプに、雨の時みずタイプに、あられのときこおりタイプになる
type forecast struct {
	ability
}

func (s *forecast) CorrectStatus(st situation.SituationChecker) *status.StatsCorrectors {
	c := status.NewStatsCorrectors()
	switch {
	case st.IsWeather(field.Sunny):
		c.Types(types.NewTypes(types.Fire))
	case st.IsWeather(field.Rainy):
		c.Types(types.NewTypes(types.Water))
	case st.IsWeather(field.Snow):
		c.Types(types.NewTypes(types.Ice))
	}
	return c
}
