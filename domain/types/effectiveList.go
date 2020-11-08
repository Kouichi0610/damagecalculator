package types

/*
	攻撃側として有効なタイプ一覧を取得
*/
func (t Type) AttackEffectiveTypes() []Type {
	list := List()
	res := make([]Type, 0)
	for _, ty := range list {
		ef := t.Magnification(ty)
		if ef < Flat {
			continue
		}
		res = append(res, ty)
	}
	return res
}
