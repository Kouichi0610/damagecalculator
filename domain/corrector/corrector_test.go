package corrector

import (
	"testing"
)

func Test_Correctors_Rules(t *testing.T) {
}

func Test_Appends(t *testing.T) {
	c := NewCorrectors()
	var add []Corrector
	c.Appends(add...)
	c.Appends()
	// nilを渡しても影響がない事
	if len(c.c) != 0 {
		t.Error()
	}
	add = []Corrector{
		NewAttack(1, 2),
		NewDefense(2, 3),
		NewPower(3, 4),
		NewDamage(2, 1),
	}
	c.Appends(add...)
	if c.CorrectAttack(100) != 50 {
		t.Error()
	}
	if c.CorrectDefense(100) != 67 {
		t.Error()
	}
	if c.CorrectPower(100) != 75 {
		t.Error()
	}
	if c.CorrectDamage(100) != 200 {
		t.Error()
	}
}

// 他のカテゴリの補正に影響を与えないこと
func Test_Correctors(t *testing.T) {
	c := NewCorrectors()
	c.Append(NewAttack(1, 2))
	c.Append(NewDefense(2, 3))
	c.Append(NewPower(3, 4))
	c.Append(NewDamage(2, 1))

	if c.CorrectAttack(100) != 50 {
		t.Error()
	}
	c.Append(NewAttack(1, 2))
	if c.CorrectAttack(100) != 25 {
		t.Error()
	}

	if c.CorrectDefense(100) != 67 {
		t.Error()
	}
	if c.CorrectPower(100) != 75 {
		t.Error()
	}
	if c.CorrectDamage(100) != 200 {
		t.Error()
	}
}

func Test_Corrector(t *testing.T) {
	c := newCorrector(Attack, drop5_pick5over, 1, 2)
	if c.Caterogy() != Attack {
		t.Error()
	}
	if c.Correct(27) != 13 {
		t.Error()
	}
}

func Test_四捨五入(t *testing.T) {
	// 7.5 -> 8
	a := drop4_pick5(15, 2048)
	if a != 8 {
		t.Errorf("%d", a)
	}
	// 20.4 -> 20
	a = drop4_pick5(34, 2457)
	if a != 20 {
		t.Errorf("%d", a)
	}
}
func Test_五捨五超入(t *testing.T) {
	//drop5_pick5over
	// 7.5 -> 7
	a := drop5_pick5over(15, 2048)
	if a != 7 {
		t.Errorf("%d", a)
	}
	// 2.52 -> 3
	a = drop5_pick5over(36, 287)
	if a != 3 {
		t.Errorf("%d", a)
	}
}
func Test_小数点以下切り捨て(t *testing.T) {
	//omit
	a := omit(49, 409)
	if a != 4 {
		t.Errorf("%d", a)
	}
}
