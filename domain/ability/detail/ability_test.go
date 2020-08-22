package detail

import (
	"damagecalculator/domain/ability/correct"
	"damagecalculator/domain/ability/situation"
	"damagecalculator/domain/corrector"
	"damagecalculator/domain/field"
	"damagecalculator/domain/move"
	"damagecalculator/domain/move/attribute"
	"damagecalculator/domain/status"
	"damagecalculator/domain/types"
	"reflect"
	"testing"
)

func Test_AbilityField機能テスト(t *testing.T) {
	a := &testAbility{}
	d := &testAbility{}
	af := NewAbilityField(a, d).(*abilityField)
	st := new(testSituation)

	if af == nil {
		t.Error()
	}
	// 攻撃側と防御側のCorrectorを合成すること
	c := af.Correctors(st)
	if len(c) != 2 {
		t.Error()
	}
	dmg := uint(100)
	for _, co := range c {
		dmg = co.Correct(dmg)
	}
	if dmg != 600 {
		t.Error()
	}

	// わざの効果自体を書き換えること
	mv := createSkillLinkData(2, 5)
	amv := af.RewriteMoveFactory(*mv)
	if amv.Type != types.Ground && amv.Accuracy != 50 {
		t.Error()
	}

	// 能力値補正を掛けること
	ac := af.CorrectAttackerStatus(st)
	dc := af.CorrectDefenderStatus(st)
	corrected := ac.Correct(100, 100, 100, 100, 100)
	if !reflect.DeepEqual(corrected, [5]uint{200, 100, 100, 100, 100}) {
		t.Errorf("%v", corrected)
	}
	corrected = dc.Correct(100, 100, 100, 100, 100)
	if !reflect.DeepEqual(corrected, [5]uint{100, 200, 100, 100, 100}) {
		t.Errorf("%v", corrected)
	}

	// erase実行後、デフォルトの特性に書き換えられていること
	af.erase()
	_, ok := af.at.(*abilityImpl)
	if !ok {
		t.Error()
	}
	_, ok = af.df.(*abilityImpl)
	if !ok {
		t.Error()
	}
}

// デフォルトのとくせい とくに何もしないこと
func Test_abilityImpl(t *testing.T) {
	a := &abilityImpl{}
	e := new(testEraser)
	if a.isAttacker {
		t.Error()
	}
	a.setAttacker(true)
	if !a.isAttacker {
		t.Error()
	}

	a.onField(e)
	a.onAttack(e)
	if e.erased {
		t.Error()
	}

	st := &testSituation{}
	c := a.Correctors(st)
	if len(c) != 1 {
		t.Error()
	}
	if c[0].Correct(100) != 100 {
		t.Error()
	}

	// 書き換えないこと
	mv := createSkillLinkData(2, 5)
	nextMove := a.RewriteMoveFactory(*mv)
	if !reflect.DeepEqual(mv, nextMove) {
		t.Errorf("%v", nextMove)
	}

	sc := a.CorrectStatus(st)
	corrected := sc.Correct(100, 100, 100, 100, 100)
	if !reflect.DeepEqual(corrected, [5]uint{100, 100, 100, 100, 100}) {
		t.Error()
	}
}

// 天候に応じてタイプを変えること
func Test_てんきや(t *testing.T) {
	a := new(forecast)
	st := &testSituation{}

	expects := map[field.Weather]*types.Types{
		field.NoWeather: types.NewTypes(types.Normal),
		field.Sunny:     types.NewTypes(types.Fire),
		field.Rainy:     types.NewTypes(types.Water),
		field.Snow:      types.NewTypes(types.Ice),
		field.SandStorm: types.NewTypes(types.Normal),
	}

	dty := types.NewTypes(types.Normal)
	for weather, expect := range expects {
		st.weather = weather
		co := a.CorrectStatus(st)
		actual := co.CorrectTypes(dty)
		if !actual.Equal(expect) {
			t.Errorf("actual[%s] expect[%s]", actual.String(), expect.String())
		}
	}
}

// フィールドに応じてタイプを変えること
func Test_ぎたい(t *testing.T) {
	a := new(mimicry)
	st := &testSituation{}
	expects := map[field.Field]*types.Types{
		field.NoField:       types.NewTypes(types.Normal),
		field.ElectricField: types.NewTypes(types.Electric),
		field.GrassField:    types.NewTypes(types.Grass),
		field.MystField:     types.NewTypes(types.Fairy),
		field.PsychoField:   types.NewTypes(types.Psychic),
	}
	dty := types.NewTypes(types.Normal)
	for fl, expect := range expects {
		st.field = fl
		co := a.CorrectStatus(st)
		actual := co.CorrectTypes(dty)
		if !actual.Equal(expect) {
			t.Errorf("actual[%s] expect[%s]", actual.String(), expect.String())
		}
	}
}

// 使う技のタイプに応じてタイプを変えること
func Test_へんげんじざい(t *testing.T) {
	a := new(protean)
	a.setAttacker(true)
	st := &testSituation{}
	dty := types.NewTypes(types.Normal)
	st.skillType = []types.Type{types.Ghost}
	co := a.CorrectStatus(st)

	if !co.CorrectTypes(dty).Equal(types.NewTypes(types.Ghost)) {
		t.Error()
	}

	// 攻撃側でなければ変化ないこと
	a.setAttacker(false)
	co = a.CorrectStatus(st)
	if !co.CorrectTypes(dty).Equal(types.NewTypes(types.Normal)) {
		t.Error()
	}
}

func Test_スキルリンク(t *testing.T) {
	a := new(skillLink)
	a.setAttacker(true)
	// 2～5を5回に書き換えること
	mf := &move.MoveFactory{
		CountMin: 2,
		CountMax: 5,
	}
	actual := a.RewriteMoveFactory(*mf)
	if !(actual.CountMin == 5 && actual.CountMax == 5) {
		t.Error()
	}

	// 2～5回でなければ変更無い事
	mf = &move.MoveFactory{
		CountMin: 1,
		CountMax: 5,
	}
	actual = a.RewriteMoveFactory(*mf)
	if !(actual.CountMin == 1 && actual.CountMax == 5) {
		t.Error()
	}

	// 防御側では効果ない事
	a.setAttacker(false)
	mf = &move.MoveFactory{
		CountMin: 2,
		CountMax: 5,
	}
	actual = a.RewriteMoveFactory(*mf)
	if !(actual.CountMin == 2 && actual.CountMax == 5) {
		t.Error()
	}
}

func Test_フラワーギフト(t *testing.T) {
	st := &testSituation{
		weather: field.Sunny,
	}
	// 晴れの時こうげき、とくぼう1.5倍
	a := (&WeatherStatusCorrectorBuilder{
		Weather:   field.Sunny,
		Attack:    1.5,
		SpDefense: 1.5,
	}).Create()
	c := a.CorrectStatus(st)
	corrected := c.Correct(100, 100, 100, 100, 100)
	if !reflect.DeepEqual(corrected, [5]uint{150, 100, 100, 150, 100}) {
		t.Error()
	}
	// 晴れでなければ効果が無い事
	st.weather = field.Rainy
	c = a.CorrectStatus(st)
	corrected = c.Correct(100, 100, 100, 100, 100)
	if !reflect.DeepEqual(corrected, [5]uint{100, 100, 100, 100, 100}) {
		t.Error()
	}
}
func Test_サーフテール(t *testing.T) {
	st := &testSituation{
		field: field.ElectricField,
	}
	// エレキフィールド時、素早さ２倍
	a := (&FieldStatusCorrectorBuilder{
		Field: field.ElectricField,
		Speed: 2.0,
	}).Create()
	c := a.CorrectStatus(st)
	corrected := c.Correct(100, 100, 100, 100, 100)
	if !reflect.DeepEqual(corrected, [5]uint{100, 100, 100, 100, 200}) {
		t.Error()
	}

	// 他のフィールドでは効果が無い事
	st.field = field.PsychoField
	c = a.CorrectStatus(st)
	corrected = c.Correct(100, 100, 100, 100, 100)
	if !reflect.DeepEqual(corrected, [5]uint{100, 100, 100, 100, 100}) {
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

func Test_powerCorrector(t *testing.T) {
	st := &testSituation{
		action:    attribute.Fang,
		skillType: []types.Type{types.Flying},
	}
	d := &PowerCorrectorBuilder{
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

func Test_能力補正(t *testing.T) {
	a := (&StatusCorrectorBuilder{
		Types:     []types.Type{types.Water},
		Attack:    2.0,
		Defense:   3.0,
		SpAttack:  4.0,
		SpDefense: 5.0,
		Speed:     6.0,
		Weight:    0.5,
	}).Create()
	st := new(testSituation)
	c := a.CorrectStatus(st)

	corrected := c.Correct(100, 100, 100, 100, 100)
	if !reflect.DeepEqual(corrected, [5]uint{200, 300, 400, 500, 600}) {
		t.Error()
	}

	if c.CorrectWeight(100) != 50 {
		t.Error()
	}
	if !c.CorrectTypes(types.NewTypes(types.Fire)).Equal(types.NewTypes(types.Water)) {
		t.Error()
	}
}

// 攻撃側でも防御側でもとくせいをリセットすること
func Test_かがくへんかガス(t *testing.T) {
	a := new(newtralizingGas)
	a.setAttacker(true)
	e := new(testEraser)
	a.onField(e)
	if !e.erased {
		t.Error()
	}

	a.setAttacker(false)
	e = new(testEraser)
	a.onField(e)
	if !e.erased {
		t.Error()
	}
}

func Test_かたやぶり(t *testing.T) {
	a := new(moldBreaker)
	// 攻撃側で、攻撃時にとくせいを無効化すること
	a.setAttacker(true)
	e := new(testEraser)
	a.onField(e)
	if e.erased {
		t.Error()
	}
	a.onAttack(e)
	if !e.erased {
		t.Error()
	}

	// 防御側では特に効果が無い事
	a.setAttacker(false)
	e = new(testEraser)
	a.onAttack(e)
	if e.erased {
		t.Error()
	}
}

func createSkillLinkData(min, max uint) *move.MoveFactory {
	return &move.MoveFactory{
		Type:     types.Ice,
		Accuracy: 100,
		CountMin: min,
		CountMax: max,
	}
}

type testEraser struct {
	erased bool
}

func (e *testEraser) erase() {
	e.erased = true
}

// テスト用SituationChecker
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

// テスト用ability
type testAbility struct {
	isAttacker bool
}

func (a *testAbility) setAttacker(isAttacker bool) {
	a.isAttacker = isAttacker
}

func (a *testAbility) onField(eraser) {
}
func (a *testAbility) onAttack(eraser) {
}
func (a *testAbility) Correctors(situation.SituationChecker) []corrector.Corrector {
	if a.isAttacker {
		return []corrector.Corrector{corrector.NewPower(2.0)}
	}
	return []corrector.Corrector{corrector.NewPower(3.0)}
}
func (a *testAbility) CorrectStatus(situation.SituationChecker) *status.StatsCorrectors {
	if a.isAttacker {
		return status.NewStatsCorrectors().Attack(2.0)
	}
	return status.NewStatsCorrectors().Defense(2.0)
}
func (a *testAbility) RewriteMoveFactory(mv move.MoveFactory) *move.MoveFactory {
	if a.isAttacker {
		mv.Type = types.Ground
		return &mv
	}
	mv.Accuracy = 50
	return &mv
}
