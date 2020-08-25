/*
	ダメージ計算の手続き
*/
package damage

import (
	"damagecalculator/domain/corrector"
	"damagecalculator/domain/situation"
	"damagecalculator/domain/status"
	"damagecalculator/domain/types"
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
	c := corrector.NewStatsCorrector()
	level := uint(st.Attacker().Level())
	move := st.Move()

	// タイプ相性
	skillType := move.Types(st)
	attacker := st.Attacker()
	defender := st.Defender()
	if skillType.PartialMatch(attacker.Types()) {
		c.ApplyTypeMatch()
	}
	ef := skillType.Magnification(defender.Types())

	// 効果なしなら0を
	if ef.IsNoEffective() {
		return NoDamage()
	}

	at, df := move.PickStats(st)
	attack, defense := criticalStats(st.IsCritical(), at, df)
	if st.IsCritical() {
		c.ApplyCritical()
	}

	// TODO:c.ApplyCriticalを実行する前にc.Appendで急所補正を追加しても機能しない。
	//      StatsCorrector側で巻き取るかここで済ませるか
	// 補正
	c.Appends(st.Correctors()...)

	power := c.CorrectPower(move.Power(st))
	attack = c.CorrectAttack(attack)
	defense = c.CorrectDefense(defense)

	dmgs := move.Calculate(level, power, attack, defense)
	// 固定ダメージなら"こうかがない"以外の相性を無効にする
	if len(dmgs) == 1 {
		ef = types.Flat
	}

	for i := 0; i < len(dmgs); i++ {
		dmgs[i] = ef.Correct(dmgs[i])
		dmgs[i] = c.CorrectDamage(dmgs[i])
	}
	return NewDamages(dmgs)
}

/*
	・急所時、攻撃側のマイナスランクを0にする
	・急所時、防御側のプラスランクを0にする
*/
func criticalStats(isCritical bool, at, df *status.RankedValue) (a, d uint) {
	if isCritical {
		return at.IgnoreMinusValue(), df.IgnorePlusValue()
	}
	return at.Value(), df.Value()
}
