package situation

import (
	"damagecalculator/domain/field"
	"damagecalculator/domain/stats"
	"damagecalculator/domain/types"
	"damagecalculator/infra/local"
	"encoding/json"
	"reflect"
	"testing"
)

func Test_Situation生成(t *testing.T) {
	mv := local.Move()
	sp := local.Species()
	ab := local.Ability()
	it := local.Item()
	d := &SituationData{
		Move: "かみなりパンチ",
		Attacker: PokeData{
			Name:        "ピカチュウ",
			Level:       50,
			Individuals: stats.IndividualTypeMax,
			BasePoints:  BasePoints{6, 252, 0, 0, 0, 252},
			Ranks:       Ranks{1, 2, 3, 4, 5},
			Ability:     "none",
			Item:        "none",
		},
		Defender: PokeData{
			Name:        "ゼニガメ",
			Level:       50,
			Individuals: stats.IndividualTypeMax,
			BasePoints:  BasePoints{252, 0, 252, 0, 6, 0},
			Ranks:       Ranks{-1, -2, -3, -4, -5},
			Ability:     "none",
			Item:        "none",
		},
		Weather:       field.Sunny,
		Field:         field.ElectricField,
		IsCritical:    false,
		IsBurn:        false,
		IsReflector:   false,
		IsLightScreen: false,
	}

	st, err := d.Create(mv, sp, ab, it)
	if err != nil {
		t.Error()
	}
	if st == nil {
		t.Error()
	}

	// check params
	atk := st.Attacker()
	aval := [6]uint{atk.HP().Value(), atk.Attack().Value(), atk.Defense().Value(), atk.SpAttack().Value(), atk.SpDefense().Value(), atk.Speed().Value()}
	aival := [6]uint{atk.HP().Value(), atk.Attack().IgnorePlusValue(), atk.Defense().IgnorePlusValue(), atk.SpAttack().IgnorePlusValue(), atk.SpDefense().IgnorePlusValue(), atk.Speed().IgnorePlusValue()}
	def := st.Defender()
	dval := [6]uint{def.HP().Value(), def.Attack().Value(), def.Defense().Value(), def.SpAttack().Value(), def.SpDefense().Value(), def.Speed().Value()}
	dival := [6]uint{def.HP().Value(), def.Attack().IgnoreMinusValue(), def.Defense().IgnoreMinusValue(), def.SpAttack().IgnoreMinusValue(), def.SpDefense().IgnoreMinusValue(), def.Speed().IgnoreMinusValue()}

	if !reflect.DeepEqual(aval, [6]uint{111, 160, 120, 175, 210, 497}) {
		t.Errorf("%v", aval)
	}
	if !reflect.DeepEqual(aival, [6]uint{111, 107, 60, 70, 70, 142}) {
		t.Errorf("%v", aival)
	}
	if !reflect.DeepEqual(dval, [6]uint{151, 45, 58, 28, 28, 18}) {
		t.Errorf("%v", dval)
	}
	if !reflect.DeepEqual(dival, [6]uint{151, 68, 117, 70, 85, 63}) {
		t.Errorf("%v", dival)
	}

	if !st.MoveTypes().Equal(types.NewTypes(types.Electric)) {
		t.Error()
	}
	effective := st.MoveEffective()
	if !effective.IsSuper() {
		t.Error()
	}

	move := st.Move()

	if move.Power(st) != 75 {
		t.Errorf("%d", move.Power(st))
	}

	// json化->復元テスト
	byteData, err := json.Marshal(d)
	if err != nil {
		t.Error()
	}

	re := &SituationData{}
	err = json.Unmarshal(byteData, re)
	if err != nil {
		t.Error()
	}
	if !reflect.DeepEqual(d, re) {
		t.Errorf("%v", d)
		t.Errorf("%v", re)
	}
}
