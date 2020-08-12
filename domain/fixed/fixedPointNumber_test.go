package fixed

import (
	"testing"
)

func Test_四捨五入(t *testing.T) {
	f, _ := NewFixPN(0.5, Drop4Pick5)
	a := f.Correct(15)
	// 7.5 -> 8
	if a != 8 {
		t.Errorf("%d", a)
	}
	f, _ = NewFixPN(0.6, Drop4Pick5)
	a = f.Correct(34)
	// 20.4 -> 20
	if a != 20 {
		t.Errorf("%d", a)
	}
}
func Test_五捨五超入(t *testing.T) {
	f, _ := NewFixPN(0.5, Drop5Pick5Over)
	a := f.Correct(15)
	// 7.5 -> 8
	if a != 7 {
		t.Errorf("%d", a)
	}
	f, _ = NewFixPN(0.07, Drop4Pick5)
	a = f.Correct(36)
	// 2.52 -> 3
	if a != 3 {
		t.Errorf("%d", a)
	}
}
func Test_切り捨て(t *testing.T) {
	f, _ := NewFixPN(0.1, Omit)
	a := f.Correct(49)
	// 4.9 - 4
	if a != 4 {
		t.Errorf("%d", a)
	}
}

func Test_生成(t *testing.T) {
	f, err := NewFixPN(1.0, Drop4Pick5)
	if f == nil {
		t.Error()
	}
	if err != nil {
		t.Error()
	}

	f, err = NewFixPN(-0.1, Drop4Pick5)
	if f != nil {
		t.Error()
	}
	if err == nil {
		t.Error()
	}

	f, err = NewFixPN(1.2, Category(3))
	if f != nil {
		t.Error()
	}
	if err == nil {
		t.Error()
	}
}
