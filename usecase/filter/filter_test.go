package filter

import (
	"damagecalculator/domain/species"
	"damagecalculator/domain/types"
	"damagecalculator/infra/local"
	"testing"
)

func Test_Filter_種族値制限なし(t *testing.T) {
	rp := local.Species()
	filter := NewFilter(Rock, "570")
	res := filter.Filter(rp)

	filteredCount := len(res)

	filter = NewFilter(Rock, "")
	res = filter.Filter(rp)
	noFilteredCount := len(res)

	if !(noFilteredCount > filteredCount) {
		t.Error()
	}
}

func Test_Filter(t *testing.T) {
	rp := local.Species()

	// いわ570属リストを作成
	filter := NewFilter(Rock, "570")
	res := filter.Filter(rp)

	if len(res) == 0 {
		t.Errorf("No Species.")
	}
	sample := res.Find("ツンデツンデ")
	if sample == nil {
		t.Error()
	}
	if sample.Types() != "いわ/はがね" {
		t.Error()
	}
	if sample.HP() != 61 {
		t.Errorf("HP:%d", sample.HP())
	}
	if sample.Attack() != 131 {
		t.Errorf("AT:%d", sample.Attack())
	}
	if sample.Defense() != 211 {
		t.Errorf("DF:%d", sample.Defense())
	}
	if sample.SpAttack() != 53 {
		t.Errorf("SA:%d", sample.SpAttack())
	}
	if sample.SpDefense() != 101 {
		t.Errorf("SD:%d", sample.SpDefense())
	}
	if sample.Speed() != 13 {
		t.Errorf("SP:%d", sample.Speed())
	}

	// 検索できない場合も確認
	if res.Find("カイリュー") != nil {
		t.Error()
	}
}

func Test_TotalFilter(t *testing.T) {
	s := &species.Species{HP: 60, Attack: 60, Defense: 60, SpAttack: 60, SpDefense: 60, Speed: 60}
	f := createTotalFilter(360)
	a := f.Filter(s)
	if a != Include {
		t.Error()
	}
	s.HP = 59
	a = f.Filter(s)
	if a != Exclude {
		t.Error()
	}

	// 合計種族値0でも通すこと
	f = createNoLimitTotalFilter()
	s = &species.Species{HP: 0, Attack: 0, Defense: 0, SpAttack: 0, SpDefense: 0, Speed: 0}
	a = f.Filter(s)
	if a != Include {
		t.Error()
	}

}

func Test_TypeFilter(t *testing.T) {
	typeList := [][]types.Type{
		{types.Normal},
		{types.Fire},
		{types.Water},
		{types.Electric},
		{types.Grass},
		{types.Ice},
		{types.Fighting},
		{types.Poison},
		{types.Ground},
		{types.Flying},
		{types.Psychic},
		{types.Bug},
		{types.Rock},
		{types.Ghost},
		{types.Dragon},
		{types.Dark},
		{types.Steel},
		{types.Fairy},
	}
	expects := map[Type]int{
		All:      18,
		Normal:   1,
		Fire:     1,
		Water:    1,
		Electric: 1,
		Grass:    1,
		Ice:      1,
		Fighting: 1,
		Poison:   1,
		Ground:   1,
		Flying:   1,
		Psychic:  1,
		Bug:      1,
		Rock:     1,
		Ghost:    1,
		Dragon:   1,
		Dark:     1,
		Steel:    1,
		Fairy:    1,
	}

	for initType, expect := range expects {
		filter := createTypeFilter(initType)
		var actual int = 0
		for _, checkType := range typeList {
			if filter.Filter(checkType) == Include {
				actual++
			}
		}
		if actual != expect {
			t.Errorf("failed: %d actual:%d expect:%d", initType, actual, expect)
		}
	}
}
