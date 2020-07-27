/*
	ダメージ計算
	能力値(statsパッケージで生成)
	ランク補正(能力値に補正を与える)
	わざ(威力、タイプ、命中、追加効果)
	とくせい(ダメージ補正に関係する)
	フィールド(天候、エレキフィールドなど)
*/
package damage

import "damagecalculator/domain/stats"

type Damage uint

func calcDamage(l stats.Level, s SkillPower, at uint, df uint) Damage {
	tmp := uint(l)*2/5 + 2
	tmp = tmp * uint(s) * at / df
	tmp = tmp/50 + 2
	return Damage(tmp)
}
func calcDamageArray(l stats.Level, s SkillPower, at uint, df uint) []Damage {
	res := make([]Damage, 0)
	b := calcDamage(l, s, at, df)
	for mag := 0.85; mag <= 1.0; mag += 0.01 {
		d := uint(float64(b) * mag)
		res = append(res, Damage(d))
	}
	return res
}
