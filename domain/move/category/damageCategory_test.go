package category

import (
	"damagecalculator/domain/stats"
	"damagecalculator/domain/status"
	"damagecalculator/domain/types"
	"testing"
)

func Test_Category_物理(t *testing.T) {
	c, _ := NewCategory(Physical)
	at := newStats(1, 2, 3, 4, 5, 6)
	df := newStats(7, 8, 9, 10, 11, 12)
	av, dv := c(at, df)
	if av.Value() != 2 {
		t.Errorf("%d", av.Value())
	}
	if dv.Value() != 9 {
		t.Errorf("%d", dv.Value())
	}
}
func Test_Category_特殊(t *testing.T) {
	c, _ := NewCategory(Special)
	at := newStats(1, 2, 3, 4, 5, 6)
	df := newStats(7, 8, 9, 10, 11, 12)
	av, dv := c(at, df)
	if av.Value() != 4 {
		t.Errorf("%d", av.Value())
	}
	if dv.Value() != 11 {
		t.Errorf("%d", dv.Value())
	}
}
func Test_Category_サイコショック(t *testing.T) {
	c, _ := NewCategory(PsycoShock)
	at := newStats(1, 2, 3, 4, 5, 6)
	df := newStats(7, 8, 9, 10, 11, 12)
	av, dv := c(at, df)
	if av.Value() != 4 {
		t.Errorf("%d", av.Value())
	}
	if dv.Value() != 9 {
		t.Errorf("%d", dv.Value())
	}
}
func Test_Category_ボディプレス(t *testing.T) {
	c, _ := NewCategory(BodyPress)
	at := newStats(1, 2, 3, 4, 5, 6)
	df := newStats(7, 8, 9, 10, 11, 12)
	av, dv := c(at, df)
	if av.Value() != 3 {
		t.Errorf("%d", av.Value())
	}
	if dv.Value() != 9 {
		t.Errorf("%d", dv.Value())
	}
}
func Test_Category_イカサマ(t *testing.T) {
	c, _ := NewCategory(FoulPlay)
	at := newStats(1, 2, 3, 4, 5, 6)
	df := newStats(7, 8, 9, 10, 11, 12)
	av, dv := c(at, df)
	if av.Value() != 8 {
		t.Errorf("%d", av.Value())
	}
	if dv.Value() != 9 {
		t.Errorf("%d", dv.Value())
	}
}
func Test_Category_シェルアームズ(t *testing.T) {
	c, _ := NewCategory(ShellArms)
	at := attacker(110, 0, 100, 0)
	df := defender(100, 0, 50, 0)

	av, dv := c(at, df)
	if av.Value() != 100 && dv.Value() != 50 {
		t.Errorf("%d %d", av.Value(), dv.Value())
	}

	df = defender(100, 0, 50, 2)
	av, dv = c(at, df)
	if av.Value() != 110 && dv.Value() != 100 {
		t.Errorf("%d %d", av.Value(), dv.Value())
	}
}

func Test_Category_例外(t *testing.T) {
	c, err := NewCategory(DamageCategory(100))
	if c != nil {
		t.Error()
	}
	if err == nil {
		t.Error()
	}
}

func attacker(at uint, ar int, sa uint, sr int) status.StatusChecker {
	return &dummyStats{
		hp: status.NewHP(1),
		at: status.NewRankedValue(at, ar),
		df: status.NewRankedValue(1, 0),
		sa: status.NewRankedValue(sa, sr),
		sd: status.NewRankedValue(1, 0),
		sp: status.NewRankedValue(1, 0),
	}
}
func defender(df uint, dr int, sd uint, sr int) status.StatusChecker {
	return &dummyStats{
		hp: status.NewHP(1),
		at: status.NewRankedValue(1, 0),
		df: status.NewRankedValue(df, dr),
		sa: status.NewRankedValue(1, 0),
		sd: status.NewRankedValue(sd, sr),
		sp: status.NewRankedValue(1, 0),
	}
}

func newStats(hp, at, df, sa, sd, sp uint) status.StatusChecker {
	return &dummyStats{
		hp: status.NewHP(hp),
		at: status.NewRankedValue(at, 0),
		df: status.NewRankedValue(df, 0),
		sa: status.NewRankedValue(sa, 0),
		sd: status.NewRankedValue(sd, 0),
		sp: status.NewRankedValue(sp, 0),
	}
}

type dummyStats struct {
	hp                 *status.HP
	at, df, sa, sd, sp *status.RankedValue
	wt                 float64
}

func (s *dummyStats) Level() stats.Level {
	return stats.NewLevel(50)

}
func (s *dummyStats) Types() *types.Types {
	return types.NewTypes(types.Dark)
}
func (s *dummyStats) HP() *status.HP {
	return s.hp
}
func (s *dummyStats) Attack() *status.RankedValue {
	return s.at
}
func (s *dummyStats) Defense() *status.RankedValue {
	return s.df
}
func (s *dummyStats) SpAttack() *status.RankedValue {
	return s.sa
}
func (s *dummyStats) SpDefense() *status.RankedValue {
	return s.sd
}
func (s *dummyStats) Speed() *status.RankedValue {
	return s.sp
}
func (s *dummyStats) Weight() status.Weight {
	return status.Weight(s.wt)
}
func (s *dummyStats) String() string {
	return ""
}
