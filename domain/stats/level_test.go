package stats

import "testing"

func Test_LevelRange(t *testing.T) {
	min := NewLevel(0)
	if min != 1 {
		t.Error()
	}

	max := NewLevel(101)
	if max != 100 {
		t.Error()
	}

	level := NewLevel(50)
	if level != 50 {
		t.Error()
	}
}
