package damage

import "testing"

func Test_Category_Physical(t *testing.T) {
	at, df := createTestStats()
	p := NewCategory(Physical)
	a, d := p(at, df)
	if a != 2 {
		t.Error()
	}
	if d != 13 {
		t.Error()
	}
}

func Test_Category_Special(t *testing.T) {
	at, df := createTestStats()
	p := NewCategory(Special)
	a, d := p(at, df)
	if a != 4 {
		t.Error()
	}
	if d != 15 {
		t.Error()
	}
}
func Test_Category_PsycoShock(t *testing.T) {
	at, df := createTestStats()
	p := NewCategory(PsycoShock)
	a, d := p(at, df)
	if a != 4 {
		t.Error()
	}
	if d != 13 {
		t.Error()
	}
}
func Test_Category_BodyPress(t *testing.T) {
	at, df := createTestStats()
	p := NewCategory(BodyPress)
	a, d := p(at, df)
	if a != 3 {
		t.Error()
	}
	if d != 13 {
		t.Error()
	}
}
func Test_Category_FoulPlay(t *testing.T) {
	at, df := createTestStats()
	p := NewCategory(FoulPlay)
	a, d := p(at, df)
	if a != 12 {
		t.Error()
	}
	if d != 13 {
		t.Error()
	}
}
func Test_Category_ShellArms_Physical(t *testing.T) {
	at := NewStats(1, 100, 1, 120, 1, 1)
	df := NewStats(1, 1, 40, 1, 30, 1)
	p := NewCategory(ShellArms)
	a, d := p(at, df)
	if a != 120 {
		t.Error()
	}
	if d != 30 {
		t.Error()
	}
}
func Test_Category_ShellArms_Special(t *testing.T) {
	at := NewStats(1, 180, 1, 120, 1, 1)
	df := NewStats(1, 1, 40, 1, 30, 1)
	p := NewCategory(ShellArms)
	a, d := p(at, df)
	if a != 180 {
		t.Errorf("%d", a)
	}
	if d != 40 {
		t.Errorf("%d", d)
	}
}

func createTestStats() (at, df *Stats) {
	at = NewStats(1, 2, 3, 4, 5, 6)
	df = NewStats(11, 12, 13, 14, 15, 16)
	return at, df
}
