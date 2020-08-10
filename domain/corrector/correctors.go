package corrector

/*
	TODO:補正に関するモデリング詰める
	・威力、攻撃、防御、ダメージに補正を掛ける
	・ルール自体書き換える(スナイパーは急所ダメージを1.5->2.25 てきおうりょく タイプ一致補正を1.5 -> 2.0)
	同じ個所に置くべきではない
*/

// わざ、もちもの、とくせい、天候から補正メソッドを収集
// ダメージ計算時に分類、補正を掛ける
// TODO:天候、フィールド補正
type Correctors struct {
	c []Corrector
}

func NewCorrectors() *Correctors {
	return &Correctors{
		c: make([]Corrector, 0),
	}
}

func (c *Correctors) Append(f Corrector) {
	c.c = append(c.c, f)
}
func (c *Correctors) Appends(args ...Corrector) {
	if args == nil {
		return
	}
	for _, f := range args {
		c.Append(f)
	}
}

func (c *Correctors) CorrectPower(n uint) uint {
	return c.correct(n, Power)
}
func (c *Correctors) CorrectAttack(n uint) uint {
	return c.correct(n, Attack)
}
func (c *Correctors) CorrectDefense(n uint) uint {
	return c.correct(n, Defense)
}
func (c *Correctors) CorrectDamage(n uint) uint {
	res := c.correct(n, Damage)
	return res
}

func (c *Correctors) correct(n uint, t Category) uint {
	res := n
	for _, f := range c.c {
		if f.Caterogy() != t {
			continue
		}
		res = f.Correct(res)
	}
	return res
}
