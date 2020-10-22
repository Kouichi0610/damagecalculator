package speed

import (
	"damagecalculator/infra/local"
	"reflect"
	"testing"
)

func Test_Speed(t *testing.T) {
	rp := local.Species()
	g := NewGenerator()
	const level = 50

	res := g.Generate(level, rp)
	if len(res) == 0 {
		t.Error()
	}

	t.Errorf("len[%d]", len(res))

	for _, info := range res {
		t.Errorf("%d %s", info.Speed(), info.Info())
	}
	/*


		if res[0].Speed() != 5 {
			t.Error()
		}
		if res[0].Info() != "ツボツボ " {
			t.Error()
		}

		for _, sp := range res {
			t.Errorf("%d %s", sp.Speed(), sp.Info())
		}
	*/
}

// ソートされていることを確認
func Test_Filter(t *testing.T) {
	rp := local.Species()
	list := totalFilter(rp)

	slowest := list[0].Speed
	l := len(list)
	fastest := list[l-1].Speed

	if slowest > fastest {
		t.Errorf("slowest:%d fastest:%d", slowest, fastest)
	}
}

func Test_Groups(t *testing.T) {
	rp := local.Species()
	grp := newSpeedGroups(rp)
	if len(grp.list) == 0 {
		t.Error()
	}

	// 最後の候補が登録されていること
	list := totalFilter(rp)
	last := list[len(list)-1].Name
	grplast := grp.list[len(grp.list)-1].names[0]
	if last != grplast {
		t.Error()
	}
}

func Test_InfoMaker(t *testing.T) {
	names := []string{"A", "B", "C", "D"}
	const lv = 100
	const sp = 130

	res, actual := checkInfoMaker(lv, sp, names, new(highest), "130族最速 A B C", 394)
	if !res {
		t.Errorf("%t", actual)
	}
	res, actual = checkInfoMaker(lv, sp, names, new(lowest), "130族最遅 A B C", 238)
	if !res {
		t.Errorf("%t", actual)
	}
	res, actual = checkInfoMaker(lv, sp, names, new(raw), "130族補正なし+0 A B C", 296)
	if !res {
		t.Errorf("%t", actual)
	}
	res, actual = checkInfoMaker(lv, sp, names, new(raw4), "130族補正なし+4 A B C", 297)
	if !res {
		t.Errorf("%t", actual)
	}
	res, actual = checkInfoMaker(lv, sp, names, new(raw252), "130族補正なし+252 A B C", 359)
	if !res {
		t.Errorf("%t", actual)
	}
}

func checkInfoMaker(lv uint, sp uint, names []string, m infoMaker, expectInfo string, expectSpeed uint) (bool, SpeedInfo) {
	expect := &infoImpl{info: expectInfo, speed: expectSpeed}
	actual := m.create(names, lv, sp)
	return reflect.DeepEqual(actual, expect), actual
}
