package skill

import (
	"damagecalculator/domain/status"
)

/*
	わざの補正を掛けるために必要な情報を取得する
	(situation.Situationに実装する。)

	goはパッケージの相互参照を禁止しているため
	situation.Situationを直接参照できない。
*/
type SituationChecker interface {
	Attacker() status.StatusChecker
	Defender() status.StatusChecker
	// 天候、持ち物、重さ、ダイマックス、壁
}
