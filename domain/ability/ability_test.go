package ability

import (
	"damagecalculator/domain/ability/detail"
	"testing"
)

func Test_生成(t *testing.T) {
	a := new(detail.MoldBreakderBuilder)
	b := new(detail.MimicryBuilder)

	af := NewAbilityField(a, b)
	if af == nil {
		t.Error()
	}
}

func Test_OwnerSpeedCorrectorGenerator(t *testing.T) {
	s := NewOwnerSpeedCorrectorGenerator()

	a := s.Create("かそく")
	if a.Rank != 1 {
		t.Error()
	}
	if a.StatsMagnification != 1.0 {
		t.Error()
	}
	a = s.Create("ようりょくそ")
	if a.Rank != 0 {
		t.Error()
	}
	if a.StatsMagnification != 2.0 {
		t.Error()
	}

	a = s.Create("もうか")
	if a.Rank != 0 {
		t.Error()
	}
	if a.StatsMagnification != 1.0 {
		t.Error()
	}
}
func Test_OtherSpeedCorrectorGenerator(t *testing.T) {
	s := NewOtherSpeedCorrectorGenerator()
	a := s.Create("ぬめぬめ")
	if a.Rank != -1 {
		t.Error()
	}
	if a.StatsMagnification != 1.0 {
		t.Error()
	}
	a = s.Create("しんりょく")
	if a.Rank != 0 {
		t.Error()
	}
	if a.StatsMagnification != 1.0 {
		t.Error()
	}
}
