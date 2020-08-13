package status

import (
	"damagecalculator/domain/types"
	"testing"
)

func Test_Status(t *testing.T) {
	sd := &StatusData{
		Lv:            50,
		Types:         []types.Type{types.Bug},
		HP:            999,
		Attack:        10,
		Defense:       20,
		SpAttack:      300,
		SpDefense:     40,
		Speed:         50,
		AttackRank:    -1,
		DefenseRank:   +1,
		SpAttackRank:  0,
		SpDefenseRank: +4,
		SpeedRank:     -6,
		Weight:        100.0,
	}

	stats := sd.Create()

	if stats.Level() != 50 {
		t.Error()
	}
	if !stats.Types().Has(types.Bug) {
		t.Error()
	}
	if stats.HP().value != 999 {
		t.Errorf("%d", stats.HP())
	}
	if stats.Attack().Value() != 6 {
		t.Errorf("%d", stats.Attack().Value())
	}
	if stats.Defense().Value() != 30 {
		t.Errorf("%d", stats.Defense().Value())
	}
	if stats.SpAttack().Value() != 300 {
		t.Errorf("%d", stats.SpAttack().Value())
	}
	if stats.SpDefense().Value() != 120 {
		t.Errorf("%d", stats.SpDefense().Value())
	}
	if stats.Speed().Value() != 12 {
		t.Errorf("%d", stats.Speed().Value())
	}
	if stats.Weight() != 100.0 {
		t.Error()
	}
}

func Test_NewRankedStats(t *testing.T) {
	s := NewRankedStats(999, 10, 20, 300, 40, 50, -1, +1, 0, +4, -6)

	if s.hp != 999 {
		t.Error()
	}
	if s.at.Value() != 6 {
		t.Errorf("%d", s.at.Value())
	}
	if s.df.Value() != 30 {
		t.Errorf("%d", s.df.Value())
	}
	if s.sa.Value() != 300 {
		t.Errorf("%d", s.sa.Value())
	}
	if s.sd.Value() != 120 {
		t.Errorf("%d", s.sd.Value())
	}
	if s.sp.Value() != 12 {
		t.Errorf("%d", s.sp.Value())
	}
}

// 上昇補正を無視すること
func Test_RankedValue_IgnorePlus(t *testing.T) {
	r := NewRankedValue(100, 1)
	if r.Value() != 150 {
		t.Error()
	}
	if r.IgnoreMinusValue() != 150 {
		t.Error()
	}
	if r.IgnorePlusValue() != 100 {
		t.Error()
	}
}

// 下降補正を無視すること
func Test_RankedValue_IgnoreMinus(t *testing.T) {
	r := NewRankedValue(100, -1)
	if r.Value() != 66 {
		t.Error()
	}
	if r.IgnoreMinusValue() != 100 {
		t.Error()
	}
	if r.IgnorePlusValue() != 66 {
		t.Error()
	}
}

// ランクの上限下限を超えて設定できないこと
func Test_Rank上限_下限(t *testing.T) {
	min := newRank(-7)
	if min != -6 {
		t.Error()
	}
	max := newRank(7)
	if max != 6 {
		t.Error()
	}
}

// ランク補正値が正しく取れていること
func Test_Rank(t *testing.T) {
	var s uint = 100

	if newRank(-6).RankedStats(s) != 25 {
		t.Error()
	}
	if newRank(-5).RankedStats(s) != 28 {
		t.Error()
	}
	if newRank(-4).RankedStats(s) != 33 {
		t.Error()
	}
	if newRank(-3).RankedStats(s) != 40 {
		t.Error()
	}
	if newRank(-2).RankedStats(s) != 50 {
		t.Error()
	}
	if newRank(-1).RankedStats(s) != 66 {
		t.Error()
	}
	if newRank(0).RankedStats(s) != 100 {
		t.Error()
	}
	if newRank(1).RankedStats(s) != 150 {
		t.Error()
	}
	if newRank(2).RankedStats(s) != 200 {
		t.Error()
	}
	if newRank(3).RankedStats(s) != 250 {
		t.Error()
	}
	if newRank(4).RankedStats(s) != 300 {
		t.Error()
	}
	if newRank(5).RankedStats(s) != 350 {
		t.Error()
	}
	if newRank(6).RankedStats(s) != 400 {
		t.Error()
	}
}

func Test_StatsCorrectors(t *testing.T) {
	s := NewStatsCorrectors()
	if s.at.Correct(100) != 100 {
		t.Error()
	}
	if s.df.Correct(100) != 100 {
		t.Error()
	}
	if s.sa.Correct(100) != 100 {
		t.Error()
	}
	if s.sd.Correct(100) != 100 {
		t.Error()
	}
	if s.sp.Correct(100) != 100 {
		t.Error()
	}

	s.Attack(2.0).Defense(3.0).SpAttack(4.0).SpDefense(5.0).Speed(6.0)
	if s.at.Correct(100) != 200 {
		t.Error()
	}
	if s.df.Correct(100) != 300 {
		t.Error()
	}
	if s.sa.Correct(100) != 400 {
		t.Error()
	}
	if s.sd.Correct(100) != 500 {
		t.Error()
	}
	if s.sp.Correct(100) != 600 {
		t.Error()
	}

	if s.CorrectWeight(128.0) != 128.0 {
		t.Error()
	}
	s.Weight(2.0)
	if s.CorrectWeight(128.0) != 256.0 {
		t.Error()
	}
}

// 能力値に補正がかかること
// ランクなどそれ以外に影響ない事
func Test_Status補正(t *testing.T) {
	sd := &StatusData{
		Lv:            50,
		Types:         []types.Type{types.Bug},
		HP:            999,
		Attack:        100,
		Defense:       100,
		SpAttack:      100,
		SpDefense:     100,
		Speed:         100,
		AttackRank:    -4,
		DefenseRank:   -2,
		SpAttackRank:  0,
		SpDefenseRank: +4,
		SpeedRank:     +6,
	}
	st := sd.Create()
	s := NewStatsCorrectors().Attack(2.0).Defense(3.0).SpAttack(4.0).SpDefense(5.0).Speed(6.0)

	ex := s.Create(st)
	if ex.Level() != 50 {
		t.Error()
	}
	if !ex.Types().Has(types.Bug) {
		t.Error()
	}
	if ex.HP().value != 999 {
		t.Error()
	}
	if ex.Attack().v != 200 {
		t.Error()
	}
	if ex.Attack().r != -4 {
		t.Error()
	}
	if ex.Defense().v != 300 {
		t.Error()
	}
	if ex.Defense().r != -2 {
		t.Error()
	}
	if ex.SpAttack().v != 400 {
		t.Error()
	}
	if ex.SpAttack().r != 0 {
		t.Error()
	}
	if ex.SpDefense().v != 500 {
		t.Error()
	}
	if ex.SpDefense().r != 4 {
		t.Error()
	}
	if ex.Speed().v != 600 {
		t.Error()
	}
	if ex.Speed().r != 6 {
		t.Error()
	}
}
