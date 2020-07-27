package stats

import "testing"

func Test_BasePointRange(t *testing.T) {
	min := newBasePoint(0)
	if min != 0 {
		t.Error()
	}
	max := newBasePoint(253)
	if max != 252 {
		t.Error()
	}
	value := newBasePoint(180)
	if value != 180 {
		t.Error()
	}
}

func Test_BasePointStats(t *testing.T) {
	b, _ := NewBasePointStats(10, 20, 30, 40, 50, 100)
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
	if b.Speed() != 100 {
		t.Error()
	}
}
func Test_BasePointStatsOver(t *testing.T) {
	_, err := NewBasePointStats(100, 100, 100, 100, 100, 11)
	if err == nil {
		t.Error()
	}
}
