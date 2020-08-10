package corrector

/*
	急所補正など、条件によって追加されるダメージ補正

*/
type Rules struct {
	c []Corrector
}

func (r *Rules) Replace(c Category, numer, denom uint) {
	for i := 0; i < len(r.c); i++ {
		if r.c[i].Caterogy() == c {
			r.c[i] = newCorrector(c, drop5_pick5over, numer, denom)
		}
	}
}

// タイプ一致ダメージ補正(1.5)
func (r *Rules) AppendTypeMatch() {
	r.c = append(r.c, newCorrector(TypeMatch, drop5_pick5over, 3, 2))
}

// 急所ダメージ補正(1.5)
func (r *Rules) AppendCritical() {
	r.c = append(r.c, newCorrector(Critical, drop5_pick5over, 3, 2))
}

// まもるによるダメージ補正(0.25)
// TODO:ダイマックスわざでない限り無効なので省いていいかも
func (r *Rules) AppendProtect() {
	r.c = append(r.c, newCorrector(Protect, drop5_pick5over, 1, 4))
}

// 複数対象によるダメージ補正(0.75)
func (r *Rules) AppendMultiTarget() {
	r.c = append(r.c, newCorrector(MultiTarget, drop5_pick5over, 3, 4))
}

// やけどによるダメージ補正(0.5)
func (r *Rules) AppendBurn() {
	r.c = append(r.c, newCorrector(Burn, drop5_pick5over, 1, 2))
}
