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
	s.Appends(NewAttack(1, 2), NewTypeMatch(2, 1))
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
	s.append(NewDamage(1, 2))
	if len(s.c) != 1 {
		t.Error()
	}

	// 環境補正が追加されていなければ何も追加しないこと
	s.append(NewCritical(2, 3))
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
	s.append(NewCritical(4, 2))
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
		NewAttack(2, 1),
		NewDefense(3, 2),
		NewPower(4, 1),
		NewDamage(2, 1),
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
	c := newCorrector(Attack, drop5_pick5over, 1, 2)
	if c.Caterogy() != Attack {
		t.Error()
	}
	if c.Correct(27) != 13 {
		t.Error()
	}
}

func Test_四捨五入(t *testing.T) {
	// 7.5 -> 8
	a := drop4_pick5(15, 2048)
	if a != 8 {
		t.Errorf("%d", a)
	}
	// 20.4 -> 20
	a = drop4_pick5(34, 2457)
	if a != 20 {
		t.Errorf("%d", a)
	}
}
func Test_五捨五超入(t *testing.T) {
	//drop5_pick5over
	// 7.5 -> 7
	a := drop5_pick5over(15, 2048)
	if a != 7 {
		t.Errorf("%d", a)
	}
	// 2.52 -> 3
	a = drop5_pick5over(36, 287)
	if a != 3 {
		t.Errorf("%d", a)
	}
}
func Test_小数点以下切り捨て(t *testing.T) {
	//omit
	a := omit(49, 409)
	if a != 4 {
		t.Errorf("%d", a)
	}
}

func Test_Factories(t *testing.T) {
	c := NewPower(1, 2)
	if c.Caterogy() != Power {
		t.Error()
	}
	c = NewAttack(1, 2)
	if c.Caterogy() != Attack {
		t.Error()
	}
	c = NewDefense(1, 2)
	if c.Caterogy() != Defense {
		t.Error()
	}
	c = NewDamage(1, 2)
	if c.Caterogy() != Damage {
		t.Error()
	}
	c = NewTypeMatch(1, 2)
	if c.Caterogy() != TypeMatch {
		t.Error()
	}
	c = NewCritical(1, 2)
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
