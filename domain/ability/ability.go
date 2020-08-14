package ability

import (
	"damagecalculator/domain/corrector"
	"damagecalculator/domain/status"
)

type AbilityField interface {
	Attacker() Ability
	Defender() Ability
}

type Ability interface {
	// 威力補正
	Correctors(situationChecker) []corrector.Corrector
	// 能力補正
	CorrectStatus(situationChecker) *status.StatsCorrectors

	// 場に出たときに効果がある(かがくへんかガス)
	onField(AbilityField) AbilityField
	// 攻撃時効果がある(かたやぶり)
	onAttack(AbilityField) AbilityField
	// AbilityFieldに置かれた時点で設定する
	setAttacker(isAttacker bool)
}

//--------------------------------------------------
func NewAbilityField(at, df Ability) AbilityField {
	at.setAttacker(true)
	df.setAttacker(false)
	var res AbilityField
	res = &abilityField{
		at: at,
		df: df,
	}
	res = res.Attacker().onField(res)
	res = res.Defender().onField(res)
	res = res.Attacker().onAttack(res)
	res = res.Defender().onAttack(res)
	return res
}

type abilityField struct {
	at, df Ability
}

func (a *abilityField) Attacker() Ability {
	return a.at
}
func (a *abilityField) Defender() Ability {
	return a.df
}

//--------------------------------------------------

//--------------------------------
type ability struct {
	isAttacker bool
}

func (a *ability) setAttacker(isAttacker bool) {
	a.isAttacker = isAttacker
}

func (a *ability) onField(f AbilityField) AbilityField {
	return f
}

func (a *ability) onAttack(f AbilityField) AbilityField {
	return f
}
func (a *ability) Correctors(situationChecker) []corrector.Corrector {
	return nil
}
func (a *ability) CorrectStatus(situationChecker) *status.StatsCorrectors {
	return status.NewStatsCorrectors()
}
