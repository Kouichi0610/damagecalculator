package detail

import (
	"damagecalculator/domain/field"
	"damagecalculator/domain/status"
)

/*
	わざの補正を掛けるために必要な情報を取得する
*/
type SituationChecker interface {
	Attacker() status.StatusChecker
	Defender() status.StatusChecker
	// TODO:持ち物、重さ、ダイマックス、壁
	IsWeather(field.Weather) bool
}
