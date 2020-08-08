package types

import "testing"

func Test_ダメージ倍率(t *testing.T) {
	a := NewTypes([]Type{Fire, Flying})
	d := NewTypes([]Type{Bug})

	if a.Magnification(d) != Effective(4.0) {
		t.Error()
	}
}
func Test_ダメージ倍率_指定なし(t *testing.T) {
	a := NewTypes([]Type{})
	d := NewTypes([]Type{Bug})

	if a.Magnification(d) != flatEffective() {
		t.Error()
	}
}

func Test_部分一致(t *testing.T) {
	ty := NewTypes([]Type{Fire, Water, Grass})
	o := NewTypes([]Type{Water})
	o2 := NewTypes([]Type{Dragon})

	if ty.PartialMatch(o) == false {
		t.Error()
	}
	if ty.PartialMatch(o2) == true {
		t.Error()
	}

}

func Test_Has(t *testing.T) {
	ty := NewTypes([]Type{Fire, Water, Grass})
	if ty.Has(Ghost) {
		t.Error()
	}
	if !ty.Has(Water) {
		t.Error()
	}
}
