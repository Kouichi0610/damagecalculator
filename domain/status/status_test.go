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
