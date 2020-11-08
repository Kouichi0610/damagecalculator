package stats

import (
	"damagecalculator/domain/basepoints"
	"testing"
)

// 能力値を生成すること
func Test_Stats(t *testing.T) {
	l := NewLevel(50)
	n := NewNature(Bold)
	s := NewSpeciesStats(95, 109, 105, 75, 85, 56)
	i := IndividualTypeMax.Create()
	b := basePoints.New(6, 252, 0, 0, 0, 252)

	stats := NewStats(l, s, i, b, n)

	if stats.HP != 171 {
		t.Errorf("%d", stats.HP)
	}
	if stats.Attack != 144 {
		t.Errorf("%d", stats.Attack)
	}
	if stats.Defense != 137 {
		t.Errorf("%d", stats.Defense)
	}
	if stats.SpAttack != 95 {
		t.Errorf("%d", stats.SpAttack)
	}
	if stats.SpDefense != 105 {
		t.Errorf("%d", stats.SpDefense)
	}
	if stats.Speed != 108 {
		t.Errorf("%d", stats.Speed)
	}
}
