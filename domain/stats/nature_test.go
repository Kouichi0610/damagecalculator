package stats

import "testing"

func Test_NatureDescriptions(t *testing.T) {
	list := NatureDescriptions()

	if len(list) != 21 {
		t.Error()
	}
}

func Test_Nature(t *testing.T) {
	const h = 404
	const f = 299
	const u = 328
	const l = 269

	m := map[NatureType][6]uint{
		Bashful: [6]uint{h, f, f, f, f, f},
		Lonely:  [6]uint{h, u, l, f, f, f},
		Adamant: [6]uint{h, u, f, l, f, f},
		Naughty: [6]uint{h, u, f, f, l, f},
		Brave:   [6]uint{h, u, f, f, f, l},
		Bold:    [6]uint{h, l, u, f, f, f},
		Impish:  [6]uint{h, f, u, l, f, f},
		Lax:     [6]uint{h, f, u, f, l, f},
		Relaxed: [6]uint{h, f, u, f, f, l},
		Modest:  [6]uint{h, l, f, u, f, f},
		Mild:    [6]uint{h, f, l, u, f, f},
		Rash:    [6]uint{h, f, f, u, l, f},
		Quiet:   [6]uint{h, f, f, u, f, l},
		Calm:    [6]uint{h, l, f, f, u, f},
		Gentle:  [6]uint{h, f, l, f, u, f},
		Careful: [6]uint{h, f, f, l, u, f},
		Sassy:   [6]uint{h, f, f, f, u, l},
		Timid:   [6]uint{h, l, f, f, f, u},
		Hasty:   [6]uint{h, f, l, f, f, u},
		Jolly:   [6]uint{h, f, f, l, f, u},
		Naive:   [6]uint{h, f, f, f, l, u},
	}

	for key, expect := range m {
		n := NewNature(key)
		actual := testCalcStats(n)
		if actual != expect {
			t.Errorf("%d %v %v\n", key, expect, actual)
		}
	}

	// ヌケニンはHPが1になること
	{
		n := NewNature(Brave)
		n.ToShadinja()
		actual := testCalcStats(n)
		expect := [6]uint{1, u, f, f, f, l}
		if actual != expect {
			t.Errorf("shadinja %v %v\n", expect, actual)
		}
	}
	// ヌケニンのHP計算が使用した性格に影響を与えないこと
	{
		n := NewNature(Brave)
		actual := testCalcStats(n)
		if actual[0] == 1 {
			t.Error()
		}

	}
}

func testCalcStats(n *Nature) [6]uint {
	l := NewLevel(100)
	s := newSpecies(100)
	i := newIndividual(31)
	b := newBasePoint(252)
	res := [6]uint{0, 0, 0, 0, 0, 0}
	res[0] = n.calcHP(l, s, i, b)
	res[1] = n.calcAttack(l, s, i, b)
	res[2] = n.calcDefense(l, s, i, b)
	res[3] = n.calcSpAttack(l, s, i, b)
	res[4] = n.calcSpDefense(l, s, i, b)
	res[5] = n.calcSpeed(l, s, i, b)
	return res
}
