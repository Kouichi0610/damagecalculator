/*
	わざ
*/
package damage

import (
	"damagecalculator/domain/corrector"
	"damagecalculator/domain/situation"
	"damagecalculator/domain/skill"
)

// ダメージ計算の手続き
type DamageCalculator interface {
	CreateDamage(situation.SituationChecker) *Damages
}

type damageCalculator struct {
	s skill.Skill
}

func (d *damageCalculator) CreateDamage(st situation.SituationChecker) *Damages {
	c := corrector.NewCorrectors()
	level := uint(st.Attacker().Level())
	power := d.s.Power(st)
	attack := d.s.AttackStats(st)
	defense := d.s.DefenseStats(st)
	// 補正などを

	// calculateArray(level, power, attack, defense uint) []uint {
	return nil
}

/*
	・急所時、ダメージ1.5倍
	・急所時、攻撃側のマイナスランクを0にする
	・急所時、防御側のプラスランクを0にする
*/
