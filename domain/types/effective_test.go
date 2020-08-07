package types

import (
	"testing"
)

// 小数点以下は切り捨てること
func Test_Effective_Correct(t *testing.T) {
	e := notVeryEffective()
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

func Test_Effective(t *testing.T) {
	// 通常
	flat := flatEffective()
	if flat != 1.0 {
		t.Error()
	}
	// 効果はばつぐん
	super := superEffective()
	if super != 2.0 {
		t.Error()
	}
	// いまひとつ
	notVery := notVeryEffective()
	if notVery != 0.5 {
		t.Error()
	}
	// こうかなし
	noEffective := noEffective()
	if noEffective != 0.0 {
		t.Error()
	}

	calc := super * notVery
	if calc != 1.0 {
		t.Error()
	}
}
