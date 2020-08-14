package ability

import (
	"damagecalculator/domain/field"
	"damagecalculator/domain/status"
	"damagecalculator/domain/types"
)

// てんきや
// 晴れの時ほのおタイプに、雨の時みずタイプに、あられのときこおりタイプになる
type forecast struct {
	ability
}

func (s *forecast) CorrectStatus(st situationChecker) *status.StatsCorrectors {
	c := status.NewStatsCorrectors()
	switch {
	case st.IsWeather(field.Sunny):
		c.Types([]types.Type{types.Fire})
	case st.IsWeather(field.Rainy):
		c.Types([]types.Type{types.Water})
	case st.IsWeather(field.Snow):
		c.Types([]types.Type{types.Ice})
	}
	return c
}
