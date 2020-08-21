package detail

import (
	"damagecalculator/domain/corrector"
	"damagecalculator/domain/field"
	"damagecalculator/domain/types"
)

/*
	ウェザーボール
	天候 晴れ、雨、あられ、砂嵐の時に威力2倍

	晴れの時ほのおタイプ
	雨の時みずタイプ
	あられの時こおりタイプ
	砂嵐のときにいわタイプ
*/

type weatherBall struct {
	*defaultMove
}

func (s *weatherBall) Correctors(st SituationChecker) []corrector.Corrector {
	switch {
	case st.IsWeather(field.Sunny):
		return []corrector.Corrector{corrector.NewPower(2.0)}
	case st.IsWeather(field.Rainy):
		return []corrector.Corrector{corrector.NewPower(2.0)}
	case st.IsWeather(field.Snow):
		return []corrector.Corrector{corrector.NewPower(2.0)}
	case st.IsWeather(field.SandStorm):
		return []corrector.Corrector{corrector.NewPower(2.0)}
	}
	return []corrector.Corrector{corrector.NewPower(1.0)}
}
func (s *weatherBall) Types(st SituationChecker) *types.Types {
	switch {
	case st.IsWeather(field.Sunny):
		return types.NewTypes(types.Fire)
	case st.IsWeather(field.Rainy):
		return types.NewTypes(types.Water)
	case st.IsWeather(field.Snow):
		return types.NewTypes(types.Ice)
	case st.IsWeather(field.SandStorm):
		return types.NewTypes(types.Rock)
	}
	return s.defaultMove.Types(st)
}
