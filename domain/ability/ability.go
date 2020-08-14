package ability

import (
	"damagecalculator/domain/corrector"
)

type AbilityField interface {
	Attacker() Ability
	Defender() Ability
}

type Ability interface {
	// 場に出たときに効果がある(かがくへんかガス)
	onField(AbilityField) AbilityField
	// 攻撃時効果がある(かたやぶり)
	onAttack(AbilityField) AbilityField

	Correctors(situationChecker) []corrector.Corrector

	// AbilityFieldで使用
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

//--------------------------------
// かたやぶり
// 攻撃時、防御側の特性を無視する
type moldBreaker struct {
	ability
}

func (a *moldBreaker) onAttack(f AbilityField) AbilityField {
	if !a.ability.isAttacker {
		return f
	}
	return &abilityField{
		at: a,
		df: &ability{false},
	}
}

//--------------------------------
// かがくへんかガス
// 全てのとくせいを無効にする
type newtralizingGas struct {
	ability
}

func (a *newtralizingGas) onField(f AbilityField) AbilityField {
	return &abilityField{
		at: &ability{true},
		df: &ability{false},
	}
}

//--------------------------------
// ふしぎなまもり
// 効果抜群以外のダメージを受けない
type wonderGuard struct {
	ability
}

func (a *wonderGuard) Correctors(st situationChecker) []corrector.Corrector {
	// 防御側でのみ有効
	if a.ability.isAttacker {
		return nil
	}

	ef := st.SkillEffective()
	if ef.IsSuper() {
		return nil
	}
	return []corrector.Corrector{corrector.NewDamage(0)}
}

//--------------------------------
