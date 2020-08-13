package corrector

import (
	"testing"
)

func Test_NewStatsCorrector(t *testing.T) {
	s := NewStatsCorrector()
	if len(s.c) != 0 {
		t.Error()
	}
	if len(s.env) != 0 {
		t.Error()
	}
}
func Test_Appends(t *testing.T) {
	s := NewStatsCorrector()
	s.ApplyTypeMatch()

	s.Appends(NewAttack(0.5), NewTypeMatch(2.0))
	if len(s.c) != 1 {
		t.Error()
	}
	_, ok := s.env[TypeMatch]
	if !ok {
		t.Error()
	}
}
func Test_append(t *testing.T) {
	s := NewStatsCorrector()
	// 追加できること
	s.append(NewDamage(0.5))
	if len(s.c) != 1 {
		t.Error()
	}

	// 環境補正が追加されていなければ何も追加しないこと
	s.append(NewCritical(0.66))
	if len(s.c) != 1 {
		t.Error()
	}
	_, ok := s.env[Critical]
	if ok {
		t.Error()
	}

	// 環境補正が追加されているなら置き換えること
	s.ApplyCritical()
	if s.env[Critical].Correct(100) != 150 {
		t.Error()
	}
	s.append(NewCritical(2.0))
	_, ok = s.env[Critical]
	if !ok {
		t.Error()
	}

	if s.env[Critical].Correct(100) != 200 {
		t.Error()
	}
}

func Test_Apply(t *testing.T) {
	s := NewStatsCorrector()
	s.ApplyTypeMatch()
	_, ok := s.env[TypeMatch]
	if !ok {
		t.Error()
	}
	s.ApplyMultiTarget()
	_, ok = s.env[MultiTarget]
	if !ok {
		t.Error()
	}
	s.ApplyCritical()
	_, ok = s.env[Critical]
	if !ok {
		t.Error()
	}
	s.ApplyBurn()
	_, ok = s.env[Burn]
	if !ok {
		t.Error()
	}
}

func Test_Corrects(t *testing.T) {
	s := NewStatsCorrector()
	a := []Corrector{
		NewAttack(2.0),
		NewDefense(1.5),
		NewPower(4.0),
		NewDamage(2.0),
	}
	s.Appends(a...)
	s.ApplyCritical()
	s.ApplyBurn()
	s.ApplyMultiTarget()
	s.ApplyTypeMatch()
	// 1.5 * 0.5 * 0.75 * 1.5
	// 2.25

	if s.CorrectPower(100) != 400 {
		t.Errorf("%d", s.CorrectPower(100))
	}
	if s.CorrectAttack(100) != 200 {
		t.Errorf("%d", s.CorrectAttack(100))
	}
	if s.CorrectDefense(100) != 150 {
		t.Errorf("%d", s.CorrectDefense(100))
	}
	if s.CorrectDamage(100) != 168 {
		t.Errorf("%d", s.CorrectDamage(100))
	}
}

func Test_Corrector(t *testing.T) {
	c := NewAttack(0.5)
	if c.Caterogy() != Attack {
		t.Error()
	}
	if c.Correct(27) != 14 {
		t.Errorf("%d", c.Correct(27))
	}
}

func Test_Factories(t *testing.T) {
	c := NewPower(0.5)
	if c.Caterogy() != Power {
		t.Error()
	}
	c = NewAttack(0.5)
	if c.Caterogy() != Attack {
		t.Error()
	}
	c = NewDefense(0.5)
	if c.Caterogy() != Defense {
		t.Error()
	}
	c = NewDamage(0.5)
	if c.Caterogy() != Damage {
		t.Error()
	}
	c = NewTypeMatch(0.5)
	if c.Caterogy() != TypeMatch {
		t.Error()
	}
	c = NewCritical(0.5)
	if c.Caterogy() != Critical {
		t.Error()
	}
}

func Test_IsEnvironment(t *testing.T) {
	expects := map[category]bool{
		Power:       false,
		Attack:      false,
		Defense:     false,
		Damage:      false,
		TypeMatch:   true,
		Critical:    true,
		Protect:     true,
		MultiTarget: true,
		Burn:        true,
	}
	for c, e := range expects {
		if c.IsEnvironment() != e {
			t.Errorf("%d", c)
		}
	}
}
