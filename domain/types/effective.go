package types

type Effective float32

/*
	ダメージに補正をかける
	・小数点切り捨て
*/
func (e Effective) Correct(dmg uint) uint {
	res := float32(dmg) * float32(e)
	return uint(res)
}

func (e Effective) IsFlat() bool {
	return e == 1.0
}
func (e Effective) IsSuper() bool {
	return e > 1.0
}
func (e Effective) IsNotVery() bool {
	return e > 0.0 && e < 1.0
}
func (e Effective) IsNoEffective() bool {
	return e == 0.0
}

func flatEffective() Effective {
	return 1.0
}
func superEffective() Effective {
	return 2.0
}
func notVeryEffective() Effective {
	return 0.5
}
func noEffective() Effective {
	return 0.0
}
