/*
	1.技威力の補正
	2.威力決定(威力*x/4096 五捨五超入、1以上を保証)

	3.攻撃補正
	4.攻撃決定(攻撃*x/4096 五捨五超入、1以上を保証))

	5.防御補正
	6.防御決定(防御*x/4096 五捨五超入、1以上を保証))

	7.ダメージ補正
	8.ダメージ計算(タイプ相性でこうかがない以外は1以上保証)

	威力補正
	攻撃補正
	防御補正
	ダメージ補正

	特殊:急所補正(スナイパー1.5 -> 2.25)
	特殊:タイプ一致補正(てきおうりょく 1.5 -> 2)
	特殊:火傷無視 からげんき(威力2倍)
	特殊:まもる無視 フェイント
	特殊:天候無視 とくせいノーてんき


	威力、攻撃、防御、ダメージ補正はAppend
	特殊系はReplace

	4096==1の固定小数点
*/
package corrector

import "damagecalculator/domain/factor"

type Corrector interface {
	Caterogy() category
	Correct(dmg uint) uint
}

type corrector struct {
	t category     // どのパラメータに補正を賭けるか
	f factor.FixPN // 補正計算
}

func (c *corrector) Caterogy() category {
	return c.t
}

func (c *corrector) Correct(dmg uint) uint {
	return c.f.Correct(dmg)
}
