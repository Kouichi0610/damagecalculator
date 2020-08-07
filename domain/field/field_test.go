package field

import (
	"damagecalculator/domain/types"
	"testing"
)

func Test_Weather_晴れ(t *testing.T) {
	w := Sunny
	plus := weatherToPlusCorrectMap(w)
	minus := weatherToMinusCorrectMap(w)
	if correctCount(plus) != 1 {
		t.Error()
	}
	if correctCount(minus) != 1 {
		t.Error()
	}
	if plus[types.Fire] != true {
		t.Error()
	}
	if minus[types.Water] != true {
		t.Error()
	}
}
func Test_Weather_雨(t *testing.T) {
	w := Rainy
	plus := weatherToPlusCorrectMap(w)
	minus := weatherToMinusCorrectMap(w)
	if correctCount(plus) != 1 {
		t.Error()
	}
	if correctCount(minus) != 1 {
		t.Error()
	}
	if plus[types.Water] != true {
		t.Error()
	}
	if minus[types.Fire] != true {
		t.Error()
	}
}
func Test_Weather_それ以外(t *testing.T) {
	weathers := []Weather{Snow, SandStorm, NoWeather}
	for _, w := range weathers {
		p := correctCount(weatherToPlusCorrectMap(w))
		m := correctCount(weatherToMinusCorrectMap(w))
		if p != 0 {
			t.Error()
		}
		if m != 0 {
			t.Error()
		}
	}
}

// でんきタイプのみ補正がかかること
func Test_Field_エレキフィールド(t *testing.T) {
	f := ElectricField
	actual := fieldToCorrectMap(f)
	count := correctCount(actual)
	if count != 1 {
		t.Error()
	}
	if actual[types.Electric] != true {
		t.Error()
	}
}

// エスパータイプのみ補正がかかること
func Test_Field_サイコフィールド(t *testing.T) {
	f := PsychoField
	actual := fieldToCorrectMap(f)
	count := correctCount(actual)
	if count != 1 {
		t.Error()
	}
	if actual[types.Psychic] != true {
		t.Error()
	}
}

// 草タイプのみ補正がかかること
func Test_Field_グラスフィールド(t *testing.T) {
	f := GrassField
	actual := fieldToCorrectMap(f)
	count := correctCount(actual)
	if count != 1 {
		t.Error()
	}
	if actual[types.Grass] != true {
		t.Error()
	}
}

// フェアリータイプのみ補正がかかること
func Test_Field_ミストフィールド(t *testing.T) {
	f := MystField
	actual := fieldToCorrectMap(f)
	count := correctCount(actual)
	if count != 1 {
		t.Error()
	}
	if actual[types.Fairy] != true {
		t.Error()
	}
}

// どのタイプにも補正がかからないこと
func Test_Field_フィールドなし(t *testing.T) {
	f := NoField
	actual := fieldToCorrectMap(f)
	count := correctCount(actual)
	if count != 0 {
		t.Error()
	}
}

func Test_Fields(t *testing.T) {
	f := NewFields(ElectricField, Sunny)
	if f.HasFieldPlus(types.NewTypes(types.Electric)) == false {
		t.Error()
	}
	if f.HasWeather(Sunny) == false {
		t.Error()
	}
	if f.HasWeatherPlus(types.NewTypes(types.Fire)) == false {
		t.Error()
	}
	if f.HasWeatherMinus(types.NewTypes(types.Water)) == false {
		t.Error()
	}

	if f.HasSpDefensePlus(types.NewTypes(types.Rock)) == true {
		t.Error()
	}

	f = NewFields(ElectricField, SandStorm)
	if f.HasSpDefensePlus(types.NewTypes(types.Rock)) == false {
		t.Error()
	}
}

var typeArray = []types.Type{
	types.Normal,
	types.Fire,
	types.Water,
	types.Electric,
	types.Grass,
	types.Ice,
	types.Fighting,
	types.Poison,
	types.Ground,
	types.Flying,
	types.Psychic,
	types.Bug,
	types.Rock,
	types.Ghost,
	types.Dragon,
	types.Dark,
	types.Steel,
	types.Fairy,
}

func weatherToPlusCorrectMap(w Weather) map[types.Type]bool {
	res := make(map[types.Type]bool, 0)
	for _, t := range typeArray {
		res[t] = w.hasPlus(types.NewTypes(t))
	}
	return res
}
func weatherToMinusCorrectMap(w Weather) map[types.Type]bool {
	res := make(map[types.Type]bool, 0)
	for _, t := range typeArray {
		res[t] = w.hasMinus(types.NewTypes(t))
	}
	return res
}

func fieldToCorrectMap(f Field) map[types.Type]bool {
	res := make(map[types.Type]bool, 0)
	for _, t := range typeArray {
		res[t] = f.hasPlus(types.NewTypes(t))
	}
	return res
}
func correctCount(m map[types.Type]bool) int {
	var res int
	for _, v := range m {
		if v {
			res++
		}
	}
	return res
}
