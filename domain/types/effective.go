package types

/*
	ダメージに補正をかける
	・小数点切り捨て
*/
func (e Effective) Correct(dmg uint) uint {
	res := float32(dmg) * float32(e)
	return uint(res)
}

func (e Effective) IsFlat() bool {
	return e == Flat
}
func (e Effective) IsSuper() bool {
	return e > Flat
}
func (e Effective) IsNotVery() bool {
	return e > NoEffective && e < Flat
}
func (e Effective) IsNoEffective() bool {
	return e == NoEffective
}
