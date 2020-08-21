package correct

import (
	"damagecalculator/domain/corrector"
	"damagecalculator/domain/field"
	"damagecalculator/domain/move/attribute"
	"damagecalculator/domain/types"
	"testing"
)

func Test_生成確認(t *testing.T) {
	datas := []PowerCorrectorBuilder{
		&TypeAttackData{},
		&ActionAttackData{},
		&TypeDefenseData{},
		&ActionDefenseData{},
	}
	if _, ok := datas[0].Create().(*typeAttack); !ok {
		t.Error()
	}
	if _, ok := datas[1].Create().(*actionAttack); !ok {
		t.Error()
	}
	if _, ok := datas[2].Create().(*typeDefense); !ok {
		t.Error()
	}
	if _, ok := datas[3].Create().(*actionDefense); !ok {
		t.Error()
	}
}

func Test_タイプ防御補正(t *testing.T) {
	st := &testSituation{
		skillType: []types.Type{types.Water},
	}
	a := (&TypeDefenseData{Types: []types.Type{types.Water}, Scale: 0.5}).Create()

	// 条件が一致するタイプに補正を掛けること
	c := a.Correct(false, st)
	if c.Category() != corrector.Damage {
		t.Error()
	}
	if c.Correct(100) != 50 {
		t.Error()
	}

	// 一致しなければ補正を掛けないこと
	st.skillType = []types.Type{types.Fighting}
	c = a.Correct(false, st)
	if c != nil {
		t.Error()
	}

	// 防御側でなければ効果がないこと
	st.skillType = []types.Type{types.Water}
	c = a.Correct(true, st)
	if c != nil {
		t.Error()
	}
}
func Test_アクション防御補正(t *testing.T) {
	st := &testSituation{
		action: attribute.Fang,
	}
	// 特定のわざアクションの時補正を掛けること
	a := (&ActionDefenseData{attribute.Fang, 0.5}).Create()
	c := a.Correct(false, st)
	if c.Category() != corrector.Damage {
		t.Error()
	}
	if c.Correct(100) != 50 {
		t.Error()
	}

	// それ以外では補正を掛けないこと
	st.action = attribute.Contact
	c = a.Correct(false, st)
	if c != nil {
		t.Error()
	}

	// 防御側では補正を掛けないこと
	st.action = attribute.Fang
	c = a.Correct(true, st)
	if c != nil {
		t.Error()
	}
}

func Test_タイプ攻撃補正(t *testing.T) {
	st := &testSituation{
		skillType: []types.Type{types.Dragon},
	}
	a := (&TypeAttackData{Types: []types.Type{types.Dragon, types.Steel}, Scale: 1.5}).Create()

	// 条件が一致するタイプに補正を掛けること
	c := a.Correct(true, st)
	if c.Category() != corrector.Power {
		t.Error()
	}
	if c.Correct(100) != 150 {
		t.Error()
	}

	st.skillType = []types.Type{types.Steel}
	c = a.Correct(true, st)
	if c.Correct(100) != 150 {
		t.Error()
	}

	// 一致しなければ補正を掛けないこと
	st.skillType = []types.Type{types.Fighting}
	c = a.Correct(true, st)
	if c != nil {
		t.Error()
	}

	// 攻撃側でなければ効果がないこと
	st.skillType = []types.Type{types.Steel}
	c = a.Correct(false, st)
	if c != nil {
		t.Error()
	}
}
func Test_アクション攻撃補正(t *testing.T) {
	st := &testSituation{
		action: attribute.Fang,
	}
	// 特定のわざアクションの時補正を掛けること
	a := (&ActionAttackData{attribute.Fang, 1.5}).Create()
	c := a.Correct(true, st)
	if c.Category() != corrector.Power {
		t.Error()
	}
	if c.Correct(100) != 150 {
		t.Error()
	}

	// それ以外では補正を掛けないこと
	st.action = attribute.Contact
	c = a.Correct(true, st)
	if c != nil {
		t.Error()
	}

	// 防御側では補正を掛けないこと
	st.action = attribute.Fang
	c = a.Correct(false, st)
	if c != nil {
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
