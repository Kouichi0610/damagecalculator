package stats

import "testing"

// 能力値を生成すること
func Test_Stats(t *testing.T) {
	l := NewLevel(100)
	s := NewSpeciesStats(80, 130, 85, 70, 75, 103)
	i := NewIndividualStats(31, 31, 31, 31, 31, 31)
	b, _ := NewBasePointStats(6, 252, 0, 0, 0, 252)
	n := NewNature(Jolly)

	stats := NewStats(l, s, i, b, n)

	if stats.HP() != 302 {
		t.Errorf("HP %d\n", stats.HP())
	}
	if stats.Attack() != 359 {
		t.Errorf("AT %d\n", stats.Attack())
	}
	if stats.Defense() != 206 {
		t.Errorf("DF %d\n", stats.Defense())
	}
	if stats.SpAttack() != 158 {
		t.Errorf("SA %d\n", stats.SpAttack())
	}
	if stats.SpDefense() != 186 {
		t.Errorf("SD %d\n", stats.SpDefense())
	}
	if stats.Speed() != 335 {
		t.Errorf("SP %d\n", stats.Speed())
	}
}
