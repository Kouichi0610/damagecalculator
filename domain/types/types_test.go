package types

import (
	"reflect"
	"testing"
)

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

	if a.Magnification(d) != flatEffective() {
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

func Test_Has(t *testing.T) {
	ty := NewTypes(Fire, Water, Grass)
	if ty.Has(Ghost) {
		t.Error()
	}
	if !ty.Has(Water) {
		t.Error()
	}
}

func Test_TypeArray(t *testing.T) {
	ty := NewTypes(Fire, Water, Grass)
	ar := ty.TypeArray()

	if !reflect.DeepEqual(ar, []Type{Fire, Water, Grass}) {
		t.Error()
	}

	// 変更しても影響ない事
	ar = append(ar, Rock)
	ar2 := ty.TypeArray()
	if !reflect.DeepEqual(ar2, []Type{Fire, Water, Grass}) {
		t.Error()
	}
}

func Test_Types生成時の引数を書き換えても影響がない事(t *testing.T) {
	args := []Type{Fire, Water}
	ty := NewTypes(args...)

	args[0] = Grass

	if ty.TypeArray()[0] == Grass {
		t.Error()
	}
}
