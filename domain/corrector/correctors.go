package corrector

// わざ、もちもの、とくせい、天候から補正メソッドを収集
// ダメージ計算時に分類、補正を掛ける
// TODO:天候、フィールド補正
type Correctors struct {
	c             []Corrector
	appendDamages []Corrector // 状況によるダメージ補正追加
}

func NewCorrectors() *Correctors {
	return &Correctors{
		c: make([]Corrector, 0),
	}
}

func (c *Correctors) Append(f Corrector) {
	c.c = append(c.c, f)
}

// タイプ一致ダメージ補正(1.5)
func (c *Correctors) AppendTypeMatch() {
	c.appendDamages = append(c.appendDamages, newCorrector(TypeMatch, drop5_pick5over, 3, 2))
}

// 急所ダメージ補正(1.5)
func (c *Correctors) AppendCritical() {
	c.appendDamages = append(c.appendDamages, newCorrector(Critical, drop5_pick5over, 3, 2))
}

// まもるによるダメージ補正(0.25)
func (c *Correctors) AppendProtect() {
	c.appendDamages = append(c.appendDamages, newCorrector(Protect, drop5_pick5over, 1, 4))
}

// 複数対象によるダメージ補正(0.75)
func (c *Correctors) AppendMultiTarget() {
	c.appendDamages = append(c.appendDamages, newCorrector(MultiTarget, drop5_pick5over, 3, 4))
}

// やけどによるダメージ補正(0.5)
func (c *Correctors) AppendBurn() {
	c.appendDamages = append(c.appendDamages, newCorrector(Burn, drop5_pick5over, 1, 2))
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
	for _, f := range c.appendDamages {
		res = f.Correct(res)
	}
	return res
}

func (c *Correctors) correct(n uint, t category) uint {
	res := n
	for _, f := range c.c {
		if f.caterogy() != t {
			continue
		}
		res = f.Correct(res)
	}
	return res
}
