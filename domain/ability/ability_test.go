package ability

import (
	"damagecalculator/domain/corrector"
	"damagecalculator/domain/field"
	"damagecalculator/domain/types"
	"reflect"
	"testing"
)

func Test_かたやぶり(t *testing.T) {
	af := NewAbilityField(&moldBreaker{}, &wonderGuard{})

	// 防御側の特性が消されていること
	if reflect.TypeOf(af.Defender()) != reflect.TypeOf(&ability{}) {
		t.Errorf("%v", reflect.TypeOf(af.Defender()))
	}
	if reflect.TypeOf(af.Attacker()) != reflect.TypeOf(&moldBreaker{}) {
		t.Errorf("%v", reflect.TypeOf(af.Attacker()))
	}

	// 防御側がかたやぶりのとき、攻撃側に影響を与えないこと
	af = NewAbilityField(&wonderGuard{}, &moldBreaker{})
	if reflect.TypeOf(af.Attacker()) != reflect.TypeOf(&wonderGuard{}) {
		t.Errorf("%v", reflect.TypeOf(af.Attacker()))
	}
}

// 全ての特性が消されていること
func Test_かがくへんかガス(t *testing.T) {
	af := NewAbilityField(&newtralizingGas{}, &wonderGuard{})
	if reflect.TypeOf(af.Defender()) != reflect.TypeOf(&ability{}) {
		t.Errorf("%v", reflect.TypeOf(af.Defender()))
	}
	if reflect.TypeOf(af.Attacker()) != reflect.TypeOf(&ability{}) {
		t.Errorf("%v", reflect.TypeOf(af.Attacker()))
	}

	af = NewAbilityField(&wonderGuard{}, &newtralizingGas{})
	if reflect.TypeOf(af.Defender()) != reflect.TypeOf(&ability{}) {
		t.Errorf("%v", reflect.TypeOf(af.Defender()))
	}
	if reflect.TypeOf(af.Attacker()) != reflect.TypeOf(&ability{}) {
		t.Errorf("%v", reflect.TypeOf(af.Attacker()))
	}
}

func Test_ふしぎなまもり(t *testing.T) {
	st := &testSituation{effective: 2.0}
	a := &wonderGuard{}
	a.setAttacker(false)
	c := a.Correctors(st)
	// こうかがばつぐんなら補正が無い事
	if c != nil {
		t.Error()
	}
	st = &testSituation{effective: 1.0}
	c = a.Correctors(st)

	// こうかがばつぐんでないダメージを0にすること
	if c == nil {
		t.Error()
	}
	if c[0].Caterogy() != corrector.Damage {
		t.Error()
	}
	if c[0].Correct(100) != 0 {
		t.Error()
	}

	// 攻撃側なら効果が無い事
	a.setAttacker(true)
	c = a.Correctors(st)
	if c != nil {
		t.Error()
	}
}

type testSituation struct {
	effective float64
	weather   field.Weather
	field     field.Field
}

func (st *testSituation) SkillEffective() types.Effective {
	return types.Effective(st.effective)
}
func (st *testSituation) IsWeather(w field.Weather) bool {
	return st.weather == w
}
func (st *testSituation) IsField(f field.Field) bool {
	return st.field == f
}
