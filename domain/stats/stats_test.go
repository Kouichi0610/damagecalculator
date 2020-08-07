package stats

import "testing"

// 能力値を生成すること
func Test_Stats(t *testing.T) {
	l := NewLevel(50)
	n := NewNature(Bold)
	s := NewSpeciesStats(95, 109, 105, 75, 85, 56)
	i := NewIndividualStats(31, 31, 31, 31, 31, 31)
	b, _ := NewBasePointStats(6, 252, 0, 0, 0, 252)

	stats := NewStats(l, s, i, b, n)

	if stats.hp != 171 {
		t.Errorf("%d", stats.hp)
	}
	if stats.at != 144 {
		t.Errorf("%d", stats.at)
	}
	if stats.df != 137 {
		t.Errorf("%d", stats.df)
	}
	if stats.sa != 95 {
		t.Errorf("%d", stats.sa)
	}
	if stats.sd != 105 {
		t.Errorf("%d", stats.sd)
	}
	if stats.sp != 108 {
		t.Errorf("%d", stats.sp)
	}
}
