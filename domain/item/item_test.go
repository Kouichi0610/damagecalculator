package item

import (
	"damagecalculator/domain/corrector"
	"damagecalculator/domain/status"
	"damagecalculator/domain/types"
	"reflect"
	"testing"
)

func Test_Item(t *testing.T) {
	/* TODO:ItemCreatotからItemを生成できること
	var items = []ItemCreator{
		&StatsCorrectData{1.0, 1.0, 1.0, 1.0, 1.0},
		&TypeMatchCorrectData{1.5},
	}
	*/

}

// 威力補正
func Test_defaultPowerCorrector(t *testing.T) {
	s := defaultPowerCorrector()
	c := s.CorrectPower(types.NewTypes(types.Bug), types.NewTypes(types.Fire), types.NewTypes(types.Rock))
	if c.Caterogy() != corrector.Power {
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
	if c.Caterogy() != corrector.Power {
		t.Error()
	}
	if c.Correct(100) != 100 {
		t.Error()
	}
	c = s.CorrectPower(types.NewTypes(types.Bug), types.NewTypes(types.Water), types.NewTypes(types.Fire))
	if c.Caterogy() != corrector.Power {
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
	if c.Caterogy() != corrector.Power {
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
	if c.Caterogy() != corrector.Power {
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
	expects := correctArray(c)
	if !reflect.DeepEqual(expects, []uint{100, 100, 100, 100, 100}) {
		t.Errorf("%v", expects)
	}
}

// 指定が無いところに補正を掛けないこと
func Test_createStatsCorrector一部指定なし(t *testing.T) {
	d := StatsCorrectData{Attack: 1.5, Defense: 2.0}
	s := d.createStatsCorrector()
	c := s.Correct()
	expects := correctArray(c)
	if !reflect.DeepEqual(expects, []uint{150, 200, 100, 100, 100}) {
		t.Errorf("%v", expects)
	}
}

// 指定通りの補正を掛けること
func Test_createStatsCorrector全て指定(t *testing.T) {
	d := StatsCorrectData{Attack: 1.5, Defense: 2.0, SpAttack: 3.0, SpDefense: 4.0, Speed: 0.5}
	s := d.createStatsCorrector()
	c := s.Correct()
	expects := correctArray(c)
	if !reflect.DeepEqual(expects, []uint{150, 200, 300, 400, 50}) {
		t.Errorf("%v", expects)
	}
}

func correctArray(s *status.StatsCorrectors) []uint {
	res := make([]uint, 0)
	a, b, c, d, e := s.Correct(100, 100, 100, 100, 100)
	res = append(res, a, b, c, d, e)
	return res
}
