package basePoints

import (
	"testing"
)

func Test_BasePoints(t *testing.T) {
	b := New(10, 20, 30, 40, 50, 60)
	if b == nil {
		t.Error()
	}
	if b.HP() != 10 {
		t.Error()
	}
	if b.Attack() != 20 {
		t.Error()
	}
	if b.Defense() != 30 {
		t.Error()
	}
	if b.SpAttack() != 40 {
		t.Error()
	}
	if b.SpDefense() != 50 {
		t.Error()
	}
	if b.Speed() != 60 {
		t.Error()
	}
	if b.Info() != "" {
		t.Error()
	}
}

func Test_Clamp_252以内に収めること(t *testing.T) {
	const expect = 252
	actual := clamp(253)
	if actual != expect {
		t.Error()
	}
}
