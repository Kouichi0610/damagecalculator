package skill

import (
	"damagecalculator/domain/corrector"
	"damagecalculator/domain/field"
	"damagecalculator/domain/skill/category"
	_ "damagecalculator/domain/stats"
	"damagecalculator/domain/status"
	"damagecalculator/domain/types"
	"reflect"
	"testing"
)

// TODO:correctorsテスト方法(nilを返さないようにする)
// TODO:AttackCount実装後、AttackCountがCalculateに影響を与えること
func Test_攻撃回数(t *testing.T) {
	t.Fail()
}

// 補正などかけないこと
func Test_わざ_NoMethod(t *testing.T) {
	d := &SkillData{
		types:    []types.Type{types.Bug},
		power:    100,
		countMin: 1,
		countMax: 1,
		category: category.Special,
		method:   None,
	}
	s, err := NewSkill(d)
	if err != nil {
		t.Error()
	}
	if s == nil {
		t.Error()
	}

	st := DummySituation()

	if s.Correctors(st) != nil {
		t.Error()
	}
	if !s.Types(st).Has(types.Bug) {
		t.Error()
	}
	if s.Power(st) != 100 {
		t.Error()
	}
	at, df := s.PickStats(st)
	if at.Value() != 200 {
		t.Error()
	}
	if df.Value() != 10 {
		t.Error()
	}

	dmgs := s.Calculate(1, 1, 1, 1)
	if len(dmgs) != 16 {
		t.Error()
	}
}

func Test_WeatherBall(t *testing.T) {
	d := &SkillData{
		types:    []types.Type{types.Bug},
		power:    150,
		countMin: 1,
		countMax: 1,
		category: category.Special,
		method:   WeatherBall,
	}
	s, _ := NewSkill(d)
	st := WithWeather(field.NoWeather)

	// 天候なしではskillに設定されているタイプと威力をそのまま使うこと
	if !s.Types(st).Has(types.Bug) {
		t.Error()
	}
	c := s.Correctors(st)
	if c != nil {
		t.Error()
	}

	correctorCheck := func(cs []corrector.Corrector) {
		if cs == nil {
			t.Error()
			return
		}
		if len(cs) != 1 {
			t.Error()
		}
		if cs[0].Caterogy() != corrector.Power {
			t.Error()
		}
		if cs[0].Correct(100) != 200 {
			t.Error()
		}
	}

	// 晴れの時ほのおタイプで威力が倍になること
	st = WithWeather(field.Sunny)
	if !s.Types(st).Has(types.Fire) {
		t.Error()
	}
	c = s.Correctors(st)
	correctorCheck(c)

	// 雨の時水タイプで威力が倍になること
	st = WithWeather(field.Rainy)
	if !s.Types(st).Has(types.Water) {
		t.Error()
	}
	c = s.Correctors(st)
	correctorCheck(c)

	// あられのときこおりタイプで威力が倍になること
	st = WithWeather(field.Snow)
	if !s.Types(st).Has(types.Ice) {
		t.Error()
	}
	c = s.Correctors(st)
	correctorCheck(c)

	// 砂嵐のときいわタイプで威力が倍になること
	st = WithWeather(field.SandStorm)
	if !s.Types(st).Has(types.Rock) {
		t.Error()
	}
	c = s.Correctors(st)
	correctorCheck(c)
}

/*
	ダメージはレベル固定であること
	それ以外はメソッドなしと変化ないこと
*/
func Test_わざ_ちきゅうなげ(t *testing.T) {
	d := &SkillData{
		types:    []types.Type{types.Fire},
		power:    100,
		countMin: 1,
		countMax: 1,
		category: category.FoulPlay,
		method:   SeismicToss,
	}
	s, _ := NewSkill(d)
	st := DummySituation()

	if s.Correctors(st) != nil {
		t.Error()
	}
	if !s.Types(st).Has(types.Fire) {
		t.Error()
	}
	if s.Power(st) != 100 {
		t.Error()
	}
	at, df := s.PickStats(st)
	if at.Value() != 20 {
		t.Error()
	}
	if df.Value() != 40 {
		t.Error()
	}

	dmgs := s.Calculate(50, 1, 1, 1)
	if !reflect.DeepEqual(dmgs, []uint{50}) {
		t.Error()
	}

}

func WithWeather(w field.Weather) SituationChecker {
	return &testSituation{
		at: status.NewStatus(50, []types.Type{types.Fire}, 999, 100, 20, 200, 40, 180, 0, 0, 0, 0, 0),
		df: status.NewStatus(50, []types.Type{types.Water}, 500, 20, 40, 60, 10, 120, 0, 0, 0, 0, 0),
		w:  w,
	}
}

func DummySituation() SituationChecker {
	return &testSituation{
		at: status.NewStatus(50, []types.Type{types.Fire}, 999, 100, 20, 200, 40, 180, 0, 0, 0, 0, 0),
		df: status.NewStatus(50, []types.Type{types.Water}, 500, 20, 40, 60, 10, 120, 0, 0, 0, 0, 0),
	}
}

//type
type testSituation struct {
	at, df status.StatusChecker
	w      field.Weather
}

func (s *testSituation) Attacker() status.StatusChecker {
	return s.at
}
func (s *testSituation) Defender() status.StatusChecker {
	return s.df
}
func (s *testSituation) IsWeather(f field.Weather) bool {
	return s.w == f
}

//-------------------------------------------
// 継承実験
type Hoger interface {
	add(x, y int) int
	print() string
}
type Fugar interface {
	Hoger
	mul(x, y int) int
}

type hoge struct {
}

func (h *hoge) add(x, y int) int {
	return x + y
}
func (h *hoge) print() string {
	return "hoge"
}

type fuga struct {
	hoge
}

func (f *fuga) mul(x, y int) int {
	return x * y
}
func (f *fuga) print() string {
	return "fuga"
}

func Test_継承(t *testing.T) {
	var f Fugar = &fuga{}

	if f.mul(2, 3) != 6 {
		t.Error()
	}
	if f.add(2, 3) != 5 {
		t.Error()
	}
	if f.print() != "fuga" {
		t.Error()
	}
}

//-------------------------------------------
