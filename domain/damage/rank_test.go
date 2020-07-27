package damage

import "testing"

func Test_Rank上限_下限(t *testing.T) {
	min := NewRank(-7)
	if min != -6 {
		t.Error()
	}
	max := NewRank(7)
	if max != 6 {
		t.Error()
	}
}

// actual 実際
// expect 期待
func Test_Rank(t *testing.T) {
	var s uint = 100

	if NewRank(-6).RankedStats(s) != 25 {
		t.Error()
	}
	if NewRank(-5).RankedStats(s) != 28 {
		t.Error()
	}
	if NewRank(-4).RankedStats(s) != 33 {
		t.Error()
	}
	if NewRank(-3).RankedStats(s) != 40 {
		t.Error()
	}
	if NewRank(-2).RankedStats(s) != 50 {
		t.Error()
	}
	if NewRank(-1).RankedStats(s) != 66 {
		t.Error()
	}
	if NewRank(0).RankedStats(s) != 100 {
		t.Error()
	}
	if NewRank(1).RankedStats(s) != 150 {
		t.Error()
	}
	if NewRank(2).RankedStats(s) != 200 {
		t.Error()
	}
	if NewRank(3).RankedStats(s) != 250 {
		t.Error()
	}
	if NewRank(4).RankedStats(s) != 300 {
		t.Error()
	}
	if NewRank(5).RankedStats(s) != 350 {
		t.Error()
	}
	if NewRank(6).RankedStats(s) != 400 {
		t.Error()
	}
}
