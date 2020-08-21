package detail

import (
	"damagecalculator/domain/ability/situation"
	"damagecalculator/domain/corrector"
	"damagecalculator/domain/move"
	"damagecalculator/domain/status"
)

// とくせい個別インターフェイス
type ability interface {
	// わざ生成データを書き換える
	RewriteMoveFactory(move.MoveFactory) *move.MoveFactory
	// 威力補正
	Correctors(situation.SituationChecker) []corrector.Corrector
	// 能力補正
	CorrectStatus(situation.SituationChecker) *status.StatsCorrectors
	// 場に出たときに効果がある(かがくへんかガス)
	onField(eraser)
	// 攻撃時効果がある(かたやぶり)
	onAttack(eraser)
	// AbilityFieldに置かれた時点で設定する
	setAttacker(isAttacker bool)
}

type abilityImpl struct {
	isAttacker bool
}

func (a *abilityImpl) setAttacker(isAttacker bool) {
	a.isAttacker = isAttacker
}

func (a *abilityImpl) onField(eraser) {
}
func (a *abilityImpl) onAttack(eraser) {
}
func (a *abilityImpl) Correctors(situation.SituationChecker) []corrector.Corrector {
	return []corrector.Corrector{corrector.NewPower(1.0)}
}
func (a *abilityImpl) CorrectStatus(situation.SituationChecker) *status.StatsCorrectors {
	return status.NewStatsCorrectors()
}
func (a *abilityImpl) RewriteMoveFactory(mv move.MoveFactory) *move.MoveFactory {
	return &mv
}
