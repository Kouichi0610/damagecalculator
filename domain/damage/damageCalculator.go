// 状況から与えうるダメージを生成
package damage

/*
	・急所時、ダメージ1.5倍
	・急所時、攻撃側のマイナスランクを0にする
	・急所時、防御側のプラスランクを0にする
*/
type DamageCalculator interface {
	CreateDamage() *Damages
}

type damageCalculator struct {
}

/*
func (d *damageCalculator) CreateDamage(dp *DamageParams) *Damages {
	dmgs := calculateArray(dp.l, dp.p, dp.a, dp.d)
	for i := 0; i < len(dmgs); i++ {
		dmgs[i] = dp.Correct(dmgs[i])
	}
	return NewDamages(dmgs)
}
*/

func calculate(level, power, attack, defense uint) uint {
	a := level*2/5 + 2
	a = a * power * attack / defense
	a = a/50 + 2
	return a
}

func calculateArray(level, power, attack, defense uint) []uint {
	d := calculate(level, power, attack, defense)
	res := make([]uint, 0)
	for i := 0.85; i <= 1.0; i += 0.01 {
		t := uint(i * float64(d))
		res = append(res, t)
	}
	return res
}
