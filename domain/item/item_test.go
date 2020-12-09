package item

import (
	"damagecalculator/domain/corrector"
	"damagecalculator/domain/status"
	"damagecalculator/domain/types"
	"reflect"
	"testing"
)

func Test_ItemCreators(t *testing.T) {
	items := []ItemCreator{
		&NoItem{},
		&StatsCorrectData{1.0, 1.0, 1.0, 1.0, 1.0},
		&WeightCorrectData{0.5},
		&TypeCorrectData{types.Dark, 1.5},
		&PowerCorrectData{1.3},
		&SuperEffectiveCorrectData{1.5},
	}
	for i, c := range items {
		item := c.Create(true)
		if item == nil {
			t.Errorf("Failed:%d %v", i, c)
		}
	}
}

func Test_持ち物なし(t *testing.T) {
	item := (&NoItem{}).Create(true)
	res := testItem(item, 100, [5]uint{100, 100, 100, 100, 100}, 200.52, t)
	if !res {
		t.Error()
	}
}
func Test_能力値補正(t *testing.T) {
	item := (&StatsCorrectData{1.1, 1.2, 1.3, 1.4, 1.5}).Create(true)
	res := testItem(item, 100, [5]uint{110, 120, 130, 140, 150}, 200.5, t)
	if !res {
		t.Error()
	}
}
func Test_重さ補正(t *testing.T) {
	item := (&WeightCorrectData{0.5}).Create(true)
	res := testItem(item, 100, [5]uint{100, 100, 100, 100, 100}, 100.25, t)
	if !res {
		t.Error()
	}
}
func Test_タイプ威力補正(t *testing.T) {
	item := (&TypeCorrectData{types.Water, 1.5}).Create(true)
	res := testItem(item, 150, [5]uint{100, 100, 100, 100, 100}, 200.5, t)
	if !res {
		t.Error()
	}
	// タイプが異なれば補正はかからないこと
	p := item.CorrectPower(types.NewTypes(types.Bug), types.NewTypes(types.Fire), types.NewTypes(types.Dark))
	if p.Correct(100) != 100 {
		t.Error()
	}
}
func Test_威力補正(t *testing.T) {
	item := (&PowerCorrectData{1.5}).Create(true)
	res := testItem(item, 150, [5]uint{100, 100, 100, 100, 100}, 200.5, t)
	if !res {
		t.Error()
	}
	// タイプに関係なく補正がかかること
	p := item.CorrectPower(types.NewTypes(types.Bug), types.NewTypes(types.Fire), types.NewTypes(types.Dark))
	if p.Correct(100) != 150 {
		t.Error()
	}
}
func Test_こうかばつぐん補正(t *testing.T) {
	item := (&SuperEffectiveCorrectData{1.5}).Create(true)
	res := testItem(item, 150, [5]uint{100, 100, 100, 100, 100}, 200.5, t)
	if !res {
		t.Error()
	}
	// こうかばつぐんでなければ補正はかからないこと
	p := item.CorrectPower(types.NewTypes(types.Bug), types.NewTypes(types.Fire), types.NewTypes(types.Dark))
	if p.Correct(100) != 100 {
		t.Error()
	}
}

func Test_威力補正は攻撃側出なければ効果が無いこと(t *testing.T) {
	items := []ItemCreator{
		&TypeCorrectData{types.Water, 1.5},
		&PowerCorrectData{1.3},
		&SuperEffectiveCorrectData{1.5},
	}
	for i, d := range items {
		item := d.Create(false)
		p := item.CorrectPower(types.NewTypes(types.Bug), types.NewTypes(types.Fire), types.NewTypes(types.Water))
		if p.Correct(100) != 100 {
			t.Errorf("failed %v", items[i])
		}
	}
}

func testItem(item Item, powerActual uint, statsActual [5]uint, weightActual float64, t *testing.T) bool {
	p := item.CorrectPower(types.NewTypes(types.Bug), types.NewTypes(types.Fire), types.NewTypes(types.Water))
	if p.Correct(100) != powerActual {
		t.Error()
		return false
	}
	c := item.Correct()
	expects := c.Correct(100, 100, 100, 100, 100)
	if !reflect.DeepEqual(expects, statsActual) {
		t.Errorf("%v", expects)
		return false
	}
	if c.CorrectWeight(200.5) != status.Weight(weightActual) {
		t.Error()
		return false
	}
	return true
}

// 威力補正
func Test_defaultPowerCorrector(t *testing.T) {
	s := defaultPowerCorrector()
	c := s.CorrectPower(types.NewTypes(types.Bug), types.NewTypes(types.Fire), types.NewTypes(types.Rock))
	if c.Category() != corrector.Power {
		t.Error()
	}
	if c.Correct(100) != 100 {
		t.Error()
	}
}
func Test_createTypeCorrector(t *testing.T) {
	d := &TypeCorrectData{Scale: 1.5, Type: types.Fire}
	s := d.createTypeCorrector()
	c := s.CorrectPower(types.NewTypes(types.Bug), types.NewTypes(types.Water), types.NewTypes(types.Rock))
	if c.Category() != corrector.Power {
		t.Error()
	}
	if c.Correct(100) != 100 {
		t.Error()
	}
	c = s.CorrectPower(types.NewTypes(types.Bug), types.NewTypes(types.Water), types.NewTypes(types.Fire))
	if c.Category() != corrector.Power {
		t.Error()
	}
	if c.Correct(100) != 150 {
		t.Error()
	}
}

func Test_createPowerCorrector(t *testing.T) {
	d := &PowerCorrectData{Scale: 1.5}
	s := d.createPowerCorrector()
	c := s.CorrectPower(types.NewTypes(types.Bug), types.NewTypes(types.Water), types.NewTypes(types.Rock))
	if c.Category() != corrector.Power {
		t.Error()
	}
	if c.Correct(100) != 150 {
		t.Error()
	}
}

func Test_superEffectiveCorrect(t *testing.T) {
	d := &SuperEffectiveCorrectData{Scale: 1.5}
	s := d.createPowerCorrector()
	c := s.CorrectPower(types.NewTypes(types.Bug), types.NewTypes(types.Ground), types.NewTypes(types.Rock))
	if c.Category() != corrector.Power {
		t.Error()
	}
	if c.Correct(100) != 100 {
		t.Error()
	}
	c = s.CorrectPower(types.NewTypes(types.Bug), types.NewTypes(types.Bug), types.NewTypes(types.Rock))
	if c.Correct(100) != 150 {
		t.Error()
	}
	c = s.CorrectPower(types.NewTypes(types.Bug), types.NewTypes(types.Normal), types.NewTypes(types.Ghost))
	if c.Correct(100) != 100 {
		t.Error()
	}
}

// 能力値補正
func Test_defaultStatsCorrector(t *testing.T) {
	s := defaultStatsCorrector()
	c := s.Correct()
	expects := c.Correct(100, 100, 100, 100, 100)
	if !reflect.DeepEqual(expects, [5]uint{100, 100, 100, 100, 100}) {
		t.Errorf("%v", expects)
	}
}

// 指定が無いところに補正を掛けないこと
func Test_createStatsCorrector一部指定なし(t *testing.T) {
	d := StatsCorrectData{Attack: 1.5, Defense: 2.0}
	s := d.createStatsCorrector()
	c := s.Correct()
	expects := c.Correct(100, 100, 100, 100, 100)
	if !reflect.DeepEqual(expects, [5]uint{150, 200, 100, 100, 100}) {
		t.Errorf("%v", expects)
	}
}

// 指定通りの補正を掛けること
func Test_createStatsCorrector全て指定(t *testing.T) {
	d := StatsCorrectData{Attack: 1.5, Defense: 2.0, SpAttack: 3.0, SpDefense: 4.0, Speed: 0.5}
	s := d.createStatsCorrector()
	c := s.Correct()
	expects := c.Correct(100, 100, 100, 100, 100)
	if !reflect.DeepEqual(expects, [5]uint{150, 200, 300, 400, 50}) {
		t.Errorf("%v", expects)
	}
}

func Test_WeightCorrector(t *testing.T) {
	d := WeightCorrectData{Scale: 2.0}
	s := d.createStatsCorrector()
	c := s.Correct()
	expects := c.Correct(100, 100, 100, 100, 100)
	if !reflect.DeepEqual(expects, [5]uint{100, 100, 100, 100, 100}) {
		t.Errorf("%v", expects)
	}
	if c.CorrectWeight(200.5) != 401 {
		t.Error()
	}
}
