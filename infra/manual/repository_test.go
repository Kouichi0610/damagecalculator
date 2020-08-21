package manual

import (
	"damagecalculator/domain/field"
	"damagecalculator/domain/move/attribute"
	"damagecalculator/domain/types"
	"reflect"
	"testing"
)

// 取得と機能チェック
func Test_AbilitiesRepository(t *testing.T) {
	rp := new(abilityRepository)
	af := rp.Get("ヨガパワー", "ぎたい")
	st := new(testSituation)

	co := af.CorrectAttackerStatus(st)
	act := co.CorrectArray(100, 100, 100, 100, 100)
	if !reflect.DeepEqual(act, [5]uint{200, 100, 100, 100, 100}) {
		t.Errorf("%v", act)
	}

	af = rp.Get("ヨガパワー", "かがくへんかガス")
	co = af.CorrectAttackerStatus(st)
	act = co.CorrectArray(110, 100, 100, 100, 100)
	if !reflect.DeepEqual(act, [5]uint{110, 100, 100, 100, 100}) {
		t.Errorf("%v", act)
	}

}

type testSituation struct {
	skillType []types.Type
	effective float64
	weather   field.Weather
	field     field.Field
	action    attribute.Action
}

func (st *testSituation) MoveTypes() *types.Types {
	return types.NewTypes(st.skillType...)
}
func (st *testSituation) MoveEffective() types.Effective {
	return types.Effective(st.effective)
}
func (st *testSituation) IsWeather(w field.Weather) bool {
	return st.weather == w
}
func (st *testSituation) IsField(f field.Field) bool {
	return st.field == f
}
func (st *testSituation) MoveAttribute() attribute.Attribute {
	return attribute.NewAttribute(st.action, attribute.NoAttribute)
}
