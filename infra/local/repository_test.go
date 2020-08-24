package local

import (
	"damagecalculator/domain/field"
	"damagecalculator/domain/move/attribute"
	"damagecalculator/domain/move/category"
	"damagecalculator/domain/types"
	"reflect"
	"testing"
)

func Test_MovesRepository(t *testing.T) {
	rp := new(movesRepository)
	mv, err := rp.Get("イカサマ")
	if err != nil {
		t.Error()
	}
	if mv.Category != category.FoulPlay {
		t.Error()
	}
}

func Test_SpeciesRepository(t *testing.T) {
	rp := new(speciesRepository)
	sp, err := rp.Get("ランクルス")
	if err != nil {
		t.Error()
	}
	if sp == nil {
		t.Error()
	}
	if sp.Name != "ランクルス" {
		t.Error()
	}
}

func Test_ItemRepository(t *testing.T) {
	rp := new(itemRepository)
	item := rp.Get("こだわりハチマキ", true)
	co := item.Correct()

	act := co.Correct(100, 100, 100, 100, 100)
	if !reflect.DeepEqual(act, [5]uint{150, 100, 100, 100, 100}) {
		t.Error()
	}

	item = rp.Get("unknown", true)
	if item == nil {
		t.Error()
	}
}

// 取得と機能チェック
func Test_AbilitiesRepository(t *testing.T) {
	rp := new(abilityRepository)
	af := rp.Get("ヨガパワー", "ぎたい")
	st := new(testSituation)

	co := af.CorrectAttackerStatus(st)
	act := co.Correct(100, 100, 100, 100, 100)
	if !reflect.DeepEqual(act, [5]uint{200, 100, 100, 100, 100}) {
		t.Errorf("%v", act)
	}

	af = rp.Get("ヨガパワー", "かがくへんかガス")
	co = af.CorrectAttackerStatus(st)
	act = co.Correct(110, 100, 100, 100, 100)
	if !reflect.DeepEqual(act, [5]uint{110, 100, 100, 100, 100}) {
		t.Errorf("%v", act)
	}
}

// とくにpanicなど起こさないこと
func Test_例外チェック(t *testing.T) {
	rp := new(abilityRepository)
	af := rp.Get("Unknown", "Unknown")
	if af == nil {
		t.Error()
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
