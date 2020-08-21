package ability

import (
	"damagecalculator/domain/ability/correct"
	"damagecalculator/domain/corrector"
	"damagecalculator/domain/field"
	"damagecalculator/domain/move"
	"damagecalculator/domain/move/attribute"
	"damagecalculator/domain/move/count"
	"damagecalculator/domain/types"
	"testing"
)

func Test_生成(t *testing.T) {
	datas := []AbilityBuilder{
		new(NoAbilityData),
		new(MoldBreakerData),
		new(NewTralizingGasData),
		new(WonderGuardData),
		new(StatusCorrectorData),
		new(ProteanData),
		new(FieldStatusCorrectData),
		new(WeatherStatusCorrectData),
		new(ForecastData),
		new(MimicryData),
		new(SkillLinkData),
	}

	if _, ok := datas[0].Create().(*ability); !ok {
		t.Error()
	}
	if _, ok := datas[1].Create().(*moldBreaker); !ok {
		t.Error()
	}
	if _, ok := datas[2].Create().(*newtralizingGas); !ok {
		t.Error()
	}
	if _, ok := datas[3].Create().(*wonderGuard); !ok {
		t.Error()
	}
	if _, ok := datas[4].Create().(*statusCorrector); !ok {
		t.Error()
	}
	if _, ok := datas[5].Create().(*protean); !ok {
		t.Error()
	}
	if _, ok := datas[6].Create().(*fieldStatusCorrector); !ok {
		t.Error()
	}
	if _, ok := datas[7].Create().(*weatherStatusCorrector); !ok {
		t.Error()
	}
	if _, ok := datas[8].Create().(*forecast); !ok {
		t.Error()
	}
	if _, ok := datas[9].Create().(*mimicry); !ok {
		t.Error()
	}
	if _, ok := datas[10].Create().(*skillLink); !ok {
		t.Error()
	}
}

func Test_とくせいなし(t *testing.T) {
	// abilityテスト
	// 特に効果が無い事
}

// PowerCorrector複合のテスト
func Test_powerCorrector(t *testing.T) {
	st := &testSituation{
		action:    attribute.Fang,
		skillType: []types.Type{types.Flying},
	}
	d := &PowerCorrectorData{
		Builders: []correct.PowerCorrectorBuilder{
			&correct.TypeAttackData{[]types.Type{types.Flying}, 2.0},
			&correct.ActionAttackData{attribute.Fang, 3.0},
			&correct.TypeDefenseData{[]types.Type{types.Flying}, 0.0},
		},
	}

	a := d.Create()
	a.setAttacker(true)
	c := a.Correctors(st)

	if len(c) != 2 {
		t.Errorf("%d", len(c))
	}
	if c[0].Correct(100) != 200 {
		t.Error()
	}
	if c[1].Correct(100) != 300 {
		t.Error()
	}
	// 防御側効果確認
	a.setAttacker(false)
	c = a.Correctors(st)
	if len(c) != 1 {
		t.Error()
	}
	if c[0].Correct(100) != 0 {
		t.Error()
	}
}

func createSkillLinkData(min, max uint) *move.MoveFactory {
	cnt, _ := count.NewAttackCount(min, max)
	return &move.MoveFactory{
		Count: cnt,
	}
}

func Test_スキルリンク_AbilityFieldに設定して機能すること(t *testing.T) {
	a := (&SkillLinkData{}).Create()
	d := (&NoAbilityData{}).Create()
	af := NewAbilityField(a, d)

	sd := createSkillLinkData(2, 5)
	act := af.RewriteMoveFactory(*sd)

	if !(act.Count.Min() == 5 && act.Count.Max() == 5) {
		t.Error()
	}

	af = NewAbilityField(d, a)
	sd = createSkillLinkData(2, 5)
	act = af.RewriteMoveFactory(*sd)
	if !(act.Count.Min() == 2 && act.Count.Max() == 5) {
		t.Error()
	}
}

func Test_スキルリンク(t *testing.T) {
	a := (&SkillLinkData{}).Create()
	a.setAttacker(true)
	sd := createSkillLinkData(1, 1)
	act := a.RewriteMoveFactory(*sd)
	if !(act.Count.Min() == 1 && act.Count.Max() == 1) {
		t.Error()
	}

	sd = createSkillLinkData(1, 5)
	act = a.RewriteMoveFactory(*sd)
	if !(act.Count.Min() == 1 && act.Count.Max() == 5) {
		t.Error()
	}
	sd = createSkillLinkData(2, 5)
	act = a.RewriteMoveFactory(*sd)
	if !(act.Count.Min() == 5 && act.Count.Max() == 5) {
		t.Errorf("%v", act)
	}
	// 元のSkillDataは書き換えられていないこと
	if sd.Count.Min() != 2 {
		t.Error()
	}

	// 攻撃側でなければ変化ない事
	a.setAttacker(false)
	sd = createSkillLinkData(2, 5)
	act = a.RewriteMoveFactory(*sd)
	if !(act.Count.Min() == 2 && act.Count.Max() == 5) {
		t.Errorf("%v", act)
	}
}

func Test_ぎたい(t *testing.T) {
	fields := map[field.Field]types.Type{
		field.NoField:       types.Normal,
		field.ElectricField: types.Electric,
		field.PsychoField:   types.Psychic,
		field.GrassField:    types.Grass,
		field.MystField:     types.Fairy,
	}
	a := (&MimicryData{}).Create()
	for f, ex := range fields {
		st := &testSituation{field: f}
		c := a.CorrectStatus(st)
		act := c.CorrectTypes(types.NewTypes(types.Normal))
		if !act.Has(ex) {
			t.Errorf("%d -> %v", f, act)
		}
	}
}

func Test_てんきや(t *testing.T) {
	weathers := map[field.Weather]types.Type{
		field.NoWeather: types.Normal,
		field.Sunny:     types.Fire,
		field.Rainy:     types.Water,
		field.Snow:      types.Ice,
		field.SandStorm: types.Normal,
	}
	a := (&ForecastData{}).Create()
	for f, ex := range weathers {
		st := &testSituation{weather: f}
		c := a.CorrectStatus(st)
		act := c.CorrectTypes(types.NewTypes(types.Normal))
		if !act.Has(ex) {
			t.Errorf("%d -> %v", f, act)
		}
	}
}

func Test_フラワーギフト(t *testing.T) {
	st := &testSituation{
		weather: field.Sunny,
	}
	// 晴れの時こうげき、とくぼう1.5倍
	a := (&WeatherStatusCorrectData{
		Weather:   field.Sunny,
		Attack:    1.5,
		SpDefense: 1.5,
	}).Create()
	c := a.CorrectStatus(st)
	at, _, _, sd, _ := c.Correct(100, 100, 100, 100, 100)
	if at != 150 || sd != 150 {
		t.Error()
	}
	// 晴れでなければ効果が無い事
	st.weather = field.Rainy
	c = a.CorrectStatus(st)
	at, _, _, sd, _ = c.Correct(100, 100, 100, 100, 100)
	if at != 100 || sd != 100 {
		t.Error()
	}
}

func Test_サーフテール(t *testing.T) {
	st := &testSituation{
		field: field.ElectricField,
	}
	// エレキフィールド時、素早さ２倍
	a := (&FieldStatusCorrectData{
		Field: field.ElectricField,
		Speed: 2.0,
	}).Create()
	c := a.CorrectStatus(st)
	_, _, _, _, sp := c.Correct(100, 100, 100, 100, 100)
	if sp != 200 {
		t.Error()
	}

	// 他のフィールドでは効果が無い事
	st.field = field.PsychoField
	c = a.CorrectStatus(st)
	_, _, _, _, sp = c.Correct(100, 100, 100, 100, 100)
	if sp != 100 {
		t.Error()
	}
}

func Test_へんげんじざい(t *testing.T) {
	// 攻撃時、技のタイプを補正にかけること
	st := &testSituation{
		skillType: []types.Type{types.Fire},
	}
	a := (&ProteanData{}).Create()
	a.setAttacker(true)
	c := a.CorrectStatus(st)
	if !c.CorrectTypes(types.NewTypes(types.Water)).Has(types.Fire) {
		t.Error()
	}
	// 攻撃時でなければ変化が無い事
	a.setAttacker(false)
	c = a.CorrectStatus(st)
	if !c.CorrectTypes(types.NewTypes(types.Water)).Has(types.Water) {
		t.Error()
	}
}

func Test_能力補正デフォルト(t *testing.T) {
	st := &testSituation{}
	a := (&StatusCorrectorData{}).Create()
	c := a.CorrectStatus(st)
	at, df, sa, sd, sp := c.Correct(100, 100, 100, 100, 100)

	if at != 100 {
		t.Error()
	}
	if df != 100 {
		t.Error()
	}
	if sa != 100 {
		t.Error()
	}
	if sd != 100 {
		t.Error()
	}
	if sp != 100 {
		t.Error()
	}
	if c.CorrectWeight(100) != 100 {
		t.Error()
	}

	if !c.CorrectTypes(types.NewTypes(types.Fire)).Has(types.Fire) {
		t.Error()
	}
}
func Test_能力補正(t *testing.T) {
	st := &testSituation{}
	a := (&StatusCorrectorData{
		Types:     []types.Type{types.Water},
		Attack:    2.0,
		Defense:   3.0,
		SpAttack:  4.0,
		SpDefense: 5.0,
		Speed:     6.0,
		Weight:    0.5,
	}).Create()
	c := a.CorrectStatus(st)

	at, df, sa, sd, sp := c.Correct(100, 100, 100, 100, 100)
	if at != 200 {
		t.Error()
	}
	if df != 300 {
		t.Error()
	}
	if sa != 400 {
		t.Error()
	}
	if sd != 500 {
		t.Error()
	}
	if sp != 600 {
		t.Error()
	}

	if c.CorrectWeight(100) != 50 {
		t.Error()
	}
	if !c.CorrectTypes(types.NewTypes(types.Fire)).Has(types.Water) {
		t.Error()
	}
}

func Test_かたやぶり(t *testing.T) {
	af := NewAbilityField(&moldBreaker{}, &wonderGuard{})

	// 防御側の特性が消されていること
	if _, ok := af.Defender().(*ability); !ok {
		t.Error()
	}
	if _, ok := af.Attacker().(*moldBreaker); !ok {
		t.Error()
	}

	// 防御側がかたやぶりのとき、攻撃側に影響を与えないこと
	af = NewAbilityField(&wonderGuard{}, &moldBreaker{})
	if _, ok := af.Attacker().(*wonderGuard); !ok {
		t.Error()
	}
}

// 全ての特性が消されていること
func Test_かがくへんかガス(t *testing.T) {
	af := NewAbilityField(&newtralizingGas{}, &wonderGuard{})
	if _, ok := af.Defender().(*ability); !ok {
		t.Error()
	}
	if _, ok := af.Attacker().(*ability); !ok {
		t.Error()
	}

	af = NewAbilityField(&wonderGuard{}, &newtralizingGas{})
	if _, ok := af.Defender().(*ability); !ok {
		t.Error()
	}
	if _, ok := af.Attacker().(*ability); !ok {
		t.Error()
	}
}

func Test_AbilityField_Correctors(t *testing.T) {
	// ひこうタイプ2倍(フィールド飛行)
	a := (&PowerCorrectorData{
		Builders: []correct.PowerCorrectorBuilder{
			&correct.TypeAttackData{[]types.Type{types.Flying}, 2.0},
		},
	}).Create()
	// ふしぎなまもり(効果は普通)
	d := (&WonderGuardData{}).Create()
	af := NewAbilityField(a, d)
	st := &testSituation{
		skillType: []types.Type{types.Flying},
		effective: 1.0,
	}

	c := af.Correctors(st)
	if len(c) != 2 {
		t.Error()
	}

	// 攻撃側と防御側の補正が反映されていること
	if c[0].Category() != corrector.Power {
		t.Error()
	}
	if c[0].Correct(100) != 200 {
		t.Error()
	}
	if c[1].Category() != corrector.Damage {
		t.Error()
	}
	if c[1].Correct(100) != 0 {
		t.Error()
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
	if c[0].Category() != corrector.Damage {
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
