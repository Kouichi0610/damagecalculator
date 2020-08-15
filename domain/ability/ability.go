package ability

import (
	"damagecalculator/domain/ability/situation"
	"damagecalculator/domain/corrector"
	"damagecalculator/domain/skill"
	"damagecalculator/domain/status"
)

type AbilityField interface {
	Attacker() Ability
	Defender() Ability
}

type PowerCorrector interface {
	Correct(isAttacker bool, st situation.SituationChecker) corrector.Corrector
}

type Ability interface {
	// スキル生成データを書き換える
	RewriteSkillData(skill.SkillData) *skill.SkillData

	// 威力補正
	Correctors(situation.SituationChecker) []corrector.Corrector

	// 能力補正
	CorrectStatus(situation.SituationChecker) *status.StatsCorrectors

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
func (a *ability) Correctors(situation.SituationChecker) []corrector.Corrector {
	return []corrector.Corrector{corrector.NewPower(1.0)}
}
func (a *ability) CorrectStatus(situation.SituationChecker) *status.StatsCorrectors {
	return status.NewStatsCorrectors()
}
func (a *ability) RewriteSkillData(sk skill.SkillData) *skill.SkillData {
	return &sk
}
