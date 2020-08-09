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

type Corrector interface {
	Caterogy() Category
	Correct(dmg uint) uint
}

type corrector struct {
	t Category      // どのパラメータに補正を賭けるか
	f calculateFunc // 補正計算式
	m uint          // 倍率(4096==1の固定小数点)
}

func newCorrector(t Category, f calculateFunc, numer, denom uint) Corrector {
	return &corrector{
		t: t,
		f: f,
		m: numer * one / denom,
	}
}

func (c *corrector) Caterogy() Category {
	return c.t
}

func (c *corrector) Correct(dmg uint) uint {
	return c.f(dmg, c.m)
}

type calculateFunc func(d uint, m uint) uint

const one = 4096
const half = 2048

// 四捨五入
func drop4_pick5(d uint, m uint) uint {
	res := d * m
	if res%one >= 2048 {
		return res/one + 1
	}
	return res / one
}

// 五捨五超入
func drop5_pick5over(d uint, m uint) uint {
	res := d * m
	if res%one > 2048 {
		return res/one + 1
	}
	return res / one
}

// 切り捨て
func omit(d uint, m uint) uint {
	return d * m / one
}
