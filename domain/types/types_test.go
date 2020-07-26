package types

import "testing"

func Test_ダメージ倍率(t *testing.T) {
	a := NewTypes(Fire, Flying)
	d := NewTypes(Bug)

	if a.Magnification(d) != Effective(4.0) {
		t.Error()
	}
}
func Test_ダメージ倍率_指定なし(t *testing.T) {
	a := NewTypes()
	d := NewTypes(Bug)

	if a.Magnification(d) != FlatEffective() {
		t.Error()
	}
}

func Test_部分一致(t *testing.T) {
	ty := NewTypes(Fire, Water, Grass)
	o := NewTypes(Water)
	o2 := NewTypes(Dragon)

	if ty.PartialMatch(o) == false {
		t.Error()
	}
	if ty.PartialMatch(o2) == true {
		t.Error()
	}

}
