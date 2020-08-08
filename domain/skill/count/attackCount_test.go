package count

import "testing"

func Test_AttackCount(t *testing.T) {
}

func Test_AttackCount_エラーチェック(t *testing.T) {
	a, e := NewAttackCount(0, 1)
	if a != nil {
		t.Error()
	}
	if e == nil {
		t.Error()
	}
	a, e = NewAttackCount(1, 0)
	if a != nil {
		t.Error()
	}
	if e == nil {
		t.Error()
	}
	
	a, = NewAttackCount(1,1)
	if a == nil {
		t.Error()
	}
	if e != nil {
		t.Error()
	}
}
