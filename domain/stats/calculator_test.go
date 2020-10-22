package stats

import "testing"

func Test_外部用(t *testing.T) {
	l := NewLevel(100)
	s := newSpecies(130)

	highest := Highest(l, s)
	lowest := Lowest(l, s)
	raw := Raw(l, s)
	raw4 := Raw4(l, s)
	raw252 := Raw252(l, s)

	if highest != 394 {
		t.Errorf("%d", highest)
	}
	if lowest != 238 {
		t.Errorf("%d", lowest)
	}
	if raw != 296 {
		t.Errorf("%d", raw)
	}
	if raw4 != 297 {
		t.Errorf("%d", raw4)
	}
	if raw252 != 359 {
		t.Errorf("%d", raw252)
	}

}

func Test_ヌケニンHP(t *testing.T) {
	l, s, i, b := getStats()
	res := calcShadinjaHP(l, s, i, b)
	if res != 1 {
		t.Error()
	}
}
func Test_HP(t *testing.T) {
	l, s, i, b := getStats()
	res := calcHP(l, s, i, b)
	if res != 187 {
		t.Errorf("Error%d", res)
	}
}
func Test_Flat(t *testing.T) {
	l, s, i, b := getStats()
	res := calcFlat(l, s, i, b)
	if res != 132 {
		t.Errorf("Error%d", res)
	}
}
func Test_Upper(t *testing.T) {
	l, s, i, b := getStats()
	res := calcUpper(l, s, i, b)
	if res != 145 {
		t.Errorf("Error%d", res)
	}
}
func Test_Lower(t *testing.T) {
	l, s, i, b := getStats()
	res := calcLower(l, s, i, b)
	if res != 118 {
		t.Errorf("Error%d", res)
	}
}

func getStats() (Level, Species, Individual, BasePoint) {
	return NewLevel(50), newSpecies(80), newIndividual(31), newBasePoint(252)
}
