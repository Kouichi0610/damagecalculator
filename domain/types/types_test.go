package types

import (
	"testing"
)

func Test_Typesダメージ倍率(t *testing.T) {
	a := NewTypes(Fire, Flying)
	d := NewTypes(Bug)

	if a.Magnification(d) != Effective(4.0) {
		t.Error()
	}
}
func Test_Typesダメージ倍率_指定なし(t *testing.T) {
	a := NewTypes()
	d := NewTypes(Bug)

	if a.Magnification(d) != Flat {
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

func Test_全一致(t *testing.T) {
	ty := NewTypes(Fire, Water, Grass)
	o := NewTypes(Water, Grass, Fire)
	if !ty.Equal(o) {
		t.Error()
	}

	o = NewTypes(Fire, Water, Grass, Ice)
	if ty.Equal(o) {
		t.Error()
	}
	o = NewTypes(Fire, Water, Water)
	if ty.Equal(o) {
		t.Error()
	}

	ty = NewTypes()
	o = NewTypes()
	if !ty.Equal(o) {
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

// 小数点以下は切り捨てること
func Test_Effective_Correct(t *testing.T) {
	e := NotVery
	d := e.Correct(99)
	if d != 49 {
		t.Errorf("%d", d)
	}
}

func Test_Effective判定(t *testing.T) {
	if Effective(0.99).IsNotVery() == false {
		t.Error()
	}
	if Effective(0.0).IsNoEffective() == false {
		t.Error()
	}
	if Effective(1.0).IsFlat() == false {
		t.Error()
	}
	if Effective(1.1).IsSuper() == false {
		t.Error()
	}
}

func Test_Effective乗算(t *testing.T) {
	calc := Super * NotVery
	if calc != 1.0 {
		t.Error()
	}
}

// Typeを*Typesに変換できること
func Test_TypeToTypes(t *testing.T) {
	if !Electric.Types().Has(Electric) {
		t.Error()
	}
}

func Test_TypesString(t *testing.T) {
	if Fire.Types().String() != "ほのお" {
		t.Error()
	}

	if NewTypes().String() != "(none)" {
		t.Error()
	}

	if NewTypes(Fire, Water, Grass).String() != "ほのお/みず/くさ" {
		t.Error()
	}
}

func Test_string(t *testing.T) {
	expects := map[Type]string{
		Normal:   "ノーマル",
		Fire:     "ほのお",
		Water:    "みず",
		Electric: "でんき",
		Grass:    "くさ",
		Ice:      "こおり",
		Fighting: "かくとう",
		Poison:   "どく",
		Ground:   "じめん",
		Flying:   "ひこう",
		Psychic:  "エスパー",
		Bug:      "むし",
		Rock:     "いわ",
		Ghost:    "ゴースト",
		Dragon:   "ドラゴン",
		Dark:     "あく",
		Steel:    "はがね",
		Fairy:    "フェアリー",
	}
	for ty, expect := range expects {
		if ty.String() != expect {
			t.Errorf("Expect:%s Actual:%s", expect, ty.String())
		}
	}
}
