package ability

import (
	"damagecalculator/domain/corrector"
	"damagecalculator/domain/field"
	"damagecalculator/domain/skill"
	"damagecalculator/domain/skill/count"
	"damagecalculator/domain/types"
	"reflect"
	"testing"
)

func Test_生成(t *testing.T) {
	datas := []AbilityBuilder{
		new(NoEffectData),
		new(MoldBreakerData),
		new(NewTralizingGasData),
		new(WonderGuardData),
		new(StatusCorrectorData),
		new(ProteanData),
		new(FieldStatusCorrectData),
		new(WeatherStatusCorrectData),
		new(ForecastData),
		new(MimicryData),
		new(TypePowerCorrectData),
		new(ActionPowerCorrectData),
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
	if _, ok := datas[10].Create().(*typePowerCorrector); !ok {
		t.Error()
	}
	if _, ok := datas[11].Create().(*actionPowerCorrector); !ok {
		t.Error()
	}
	if _, ok := datas[12].Create().(*skillLink); !ok {
		t.Error()
	}
}

func Test_とくせいなし(t *testing.T) {
	// abilityテスト
	// 特に効果が無い事
}

func Test_スキルリンク(t *testing.T) {
	a := (&SkillLinkData{}).Create()
	count := func(min, max uint) *count.AttackCount {
		res, _ := count.NewAttackCount(min, max)
		return res
	}
	// 攻撃時、2～5回を5回にすること
	a.setAttacker(true)
	c := a.AttackCount(count(1, 1))
	if !(c.Min() == 1 && c.Max() == 1) {
		t.Error()
	}
	c = a.AttackCount(count(1, 5))
	if !(c.Min() == 1 && c.Max() == 5) {
		t.Error()
	}
	c = a.AttackCount(count(2, 5))
	if !(c.Min() == 5 && c.Max() == 5) {
		t.Errorf("%d - %d", c.Min(), c.Max())
	}

	// 防御時、変化ないこと
	a.setAttacker(false)
	c = a.AttackCount(count(2, 5))
	if !(c.Min() == 2 && c.Max() == 5) {
		t.Errorf("%d - %d", c.Min(), c.Max())
	}

}

func Test_わざアクション威力補正(t *testing.T) {
	st := &testSituation{
		action: skill.Fang,
	}
	// 特定のわざアクションの時補正を掛けること
	a := (&ActionPowerCorrectData{skill.Fang, 1.5}).Create()
	a.setAttacker(true)
	c := a.Correctors(st)
	if c[0].Caterogy() != corrector.Power {
		t.Error()
	}
	if c[0].Correct(100) != 150 {
		t.Error()
	}

	// それ以外では補正を掛けないこと
	st.action = skill.Contact
	c = a.Correctors(st)
	if c[0].Caterogy() != corrector.Power {
		t.Error()
	}
	if c[0].Correct(100) != 100 {
		t.Error()
	}

	// 防御側では補正を掛けないこと
	a.setAttacker(false)
	st.action = skill.Fang
	c = a.Correctors(st)
	if c[0].Correct(100) != 100 {
		t.Error()
	}
}
func Test_タイプ威力補正(t *testing.T) {
	st := &testSituation{
		skillType: []types.Type{types.Dragon},
	}
	a := (&TypePowerCorrectData{Types: []types.Type{types.Dragon, types.Steel}, Scale: 1.5}).Create()
	a.setAttacker(true)

	// 条件が一致するタイプに補正を掛けること
	c := a.Correctors(st)
	if c[0].Caterogy() != corrector.Power {
		t.Error()
	}
	if c[0].Correct(100) != 150 {
		t.Error()
	}

	st.skillType = []types.Type{types.Steel}
	c = a.Correctors(st)
	if c[0].Correct(100) != 150 {
		t.Error()
	}

	// 一致しなければ補正を掛けないこと
	st.skillType = []types.Type{types.Fighting}
	c = a.Correctors(st)
	if c[0].Correct(100) != 100 {
		t.Error()
	}

	// 攻撃側でなければ効果がないこと
	st.skillType = []types.Type{types.Steel}
	a.setAttacker(false)
	c = a.Correctors(st)
	if c[0].Correct(100) != 100 {
		t.Error()
	}
}

func Test_ぎたい(t *testing.T) {
	fields := map[field.Field][]types.Type{
		field.NoField:       []types.Type{types.Normal},
		field.ElectricField: []types.Type{types.Electric},
		field.PsychoField:   []types.Type{types.Psychic},
		field.GrassField:    []types.Type{types.Grass},
		field.MystField:     []types.Type{types.Fairy},
	}
	a := (&MimicryData{}).Create()
	for f, ex := range fields {
		st := &testSituation{field: f}
		c := a.CorrectStatus(st)
		act := c.CorrectTypes([]types.Type{types.Normal})
		if !reflect.DeepEqual(act, ex) {
			t.Errorf("%d -> %v", f, act)
		}
	}
}

func Test_てんきや(t *testing.T) {
	weathers := map[field.Weather][]types.Type{
		field.NoWeather: []types.Type{types.Normal},
		field.Sunny:     []types.Type{types.Fire},
		field.Rainy:     []types.Type{types.Water},
		field.Snow:      []types.Type{types.Ice},
		field.SandStorm: []types.Type{types.Normal},
	}
	a := (&ForecastData{}).Create()
	for f, ex := range weathers {
		st := &testSituation{weather: f}
		c := a.CorrectStatus(st)
		act := c.CorrectTypes([]types.Type{types.Normal})
		if !reflect.DeepEqual(act, ex) {
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
	if !reflect.DeepEqual(c.CorrectTypes([]types.Type{types.Water}), []types.Type{types.Fire}) {
		t.Error()
	}
	// 攻撃時でなければ変化が無い事
	a.setAttacker(false)
	c = a.CorrectStatus(st)
	if !reflect.DeepEqual(c.CorrectTypes([]types.Type{types.Water}), []types.Type{types.Water}) {
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

	if !reflect.DeepEqual(c.CorrectTypes([]types.Type{types.Fire}), []types.Type{types.Fire}) {
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
	if !reflect.DeepEqual(c.CorrectTypes([]types.Type{types.Fire}), []types.Type{types.Water}) {
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
	skillType []types.Type
	effective float64
	weather   field.Weather
	field     field.Field
	action    skill.Action
}

func (st *testSituation) SkillTypes() *types.Types {
	return types.NewTypes(st.skillType...)
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
func (st *testSituation) SkillAction() skill.Action {
	return st.action
}
