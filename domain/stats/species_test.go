package stats

import "testing"

func Test_SpeciesRange(t *testing.T) {
	min := newSpecies(0)
	if min != 1 {
		t.Error()
	}
	max := newSpecies(256)
	if max != 255 {
		t.Error()
	}
	s := newSpecies(100)
	if s != 100 {
		t.Error()
	}
}

func Test_SpeciesStats(t *testing.T) {
	s := NewSpeciesStats(80, 100, 60, 20, 70, 200)
	if s.HP() != 80 {
		t.Error()
	}
	if s.Attack() != 100 {
		t.Error()
	}
	if s.Defense() != 60 {
		t.Error()
	}
	if s.SpAttack() != 20 {
		t.Error()
	}
	if s.SpDefense() != 70 {
		t.Error()
	}
	if s.Speed() != 200 {
		t.Error()
	}
}
