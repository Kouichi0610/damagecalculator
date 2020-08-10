/*
	ダメージ計算の手続き
*/
package damage

import (
	"damagecalculator/domain/corrector"
	"damagecalculator/domain/situation"
	_ "damagecalculator/domain/skill"
	"damagecalculator/domain/status"
)

// ダメージ計算の手続き
type (
	DamageCalculator interface {
		CreateDamage(situation.SituationChecker) *Damages
	}
	impl struct {
	}
)

func NewDamageCalculator() DamageCalculator {
	return &impl{}
}

func (d *impl) CreateDamage(st situation.SituationChecker) *Damages {
	c := corrector.NewCorrectors()
	r := corrector.NewRules()
	level := uint(st.Attacker().Level())
	skill := st.Skill()

	// TODO:Situationからまとめて取得してしまうか
	c.Appends(skill.Correctors(st)...)

	// タイプ相性
	skillType := skill.Types(st)
	attacker := st.Attacker()
	defender := st.Defender()
	if skillType.PartialMatch(attacker.Types()) {
		r.AppendTypeMatch()
	}
	ef := skillType.Magnification(defender.Types())

	power := c.CorrectPower(skill.Power(st))
	at, df := skill.PickStats(st)
	attack, defense := criticalStats(st.IsCritical(), at, df)

	attack = c.CorrectAttack(attack)
	defense = c.CorrectDefense(defense)

	dmgs := skill.Calculate(level, power, attack, defense)

	// TODO:タイプ相性
	// TODO:タイプ一致

	for i := 0; i < len(dmgs); i++ {
		dmgs[i] = ef.Correct(dmgs[i])
		dmgs[i] = c.CorrectDamage(dmgs[i])
		dmgs[i] = r.Correct(dmgs[i])
	}
	return NewDamages(dmgs)
}

func criticalStats(isCritical bool, at, df *status.RankedValue) (a, d uint) {
	if isCritical {
		return at.IgnoreMinusValue(), df.IgnorePlusValue()
	}
	return at.Value(), df.Value()
}

/*
	・急所時、ダメージ1.5倍
	・急所時、攻撃側のマイナスランクを0にする
	・急所時、防御側のプラスランクを0にする
*/
