package skill

import (
	"damagecalculator/domain/corrector"
	"damagecalculator/domain/field"
	"damagecalculator/domain/skill/category"
	"damagecalculator/domain/status"
	"damagecalculator/domain/types"
	"math"
	"reflect"
	"testing"
)

func Test_攻撃回数(t *testing.T) {
	const min = 2
	const max = 5
	s, _ := (&SkillData{
		Types:    []types.Type{types.Bug},
		Power:    100,
		CountMin: min,
		CountMax: max,
		Category: category.Special,
		method:   None,
	}).Create()
	dmgs := s.Calculate(1, 1, 1, 1)

	var ex int
	for i := min; i <= max; i++ {
		ex += int(math.Pow(16, float64(i)))
	}
	if len(dmgs) != ex {
		t.Error()
	}
	if dmgs[0] != 1*min {
		t.Error()
	}
	if dmgs[ex-1] != 2*max {
		t.Errorf("%d", dmgs[ex-1])
	}
}

// 補正などかけないこと
func Test_わざ_NoMethod(t *testing.T) {
	s, err := (&SkillData{
		Types:    []types.Type{types.Bug},
		Power:    100,
		CountMin: 1,
		CountMax: 1,
		Category: category.Special,
		method:   None,
	}).Create()
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

func Test_gyroBallPower(t *testing.T) {
	var actual uint
	actual = gyroBallPower(150, 100)
	if actual != 17 {
		t.Errorf("%d", actual)
	}
	actual = gyroBallPower(60, 100)
	if actual != 42 {
		t.Errorf("%d", actual)
	}
	actual = gyroBallPower(25, 149)
	if actual != 150 {
		t.Errorf("%d", actual)
	}
	actual = gyroBallPower(25, 150)
	if actual != 150 {
		t.Errorf("%d", actual)
	}
	actual = gyroBallPower(180, 120)
	if actual != 17 {
		t.Errorf("%d", actual)
	}
}
func Test_ジャイロボール(t *testing.T) {
	s, _ := (&SkillData{
		Types:    []types.Type{types.Bug},
		Power:    3,
		CountMin: 1,
		CountMax: 1,
		Category: category.Physical,
		method:   GyroBall,
	}).Create()
	st := DummySituation() // 180 vs 120
	if s.Power(st) != 17 {
		t.Errorf("%d", s.Power(st))
	}
}

func Test_ウェザーボール(t *testing.T) {
	s, _ := (&SkillData{
		Types:    []types.Type{types.Bug},
		Power:    150,
		CountMin: 1,
		CountMax: 1,
		Category: category.Special,
		method:   WeatherBall,
	}).Create()
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
	s, _ := (&SkillData{
		Types:    []types.Type{types.Fire},
		Power:    100,
		CountMin: 1,
		CountMax: 1,
		Category: category.FoulPlay,
		method:   SeismicToss,
	}).Create()
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
	at := &status.StatusData{
		Lv:            50,
		Types:         []types.Type{types.Fire},
		HP:            999,
		Attack:        100,
		Defense:       20,
		SpAttack:      200,
		SpDefense:     40,
		Speed:         180,
		AttackRank:    0,
		DefenseRank:   0,
		SpAttackRank:  0,
		SpDefenseRank: 0,
		SpeedRank:     0,
	}
	df := &status.StatusData{
		Lv:            50,
		Types:         []types.Type{types.Water},
		HP:            500,
		Attack:        20,
		Defense:       40,
		SpAttack:      60,
		SpDefense:     10,
		Speed:         120,
		AttackRank:    0,
		DefenseRank:   0,
		SpAttackRank:  0,
		SpDefenseRank: 0,
		SpeedRank:     0,
	}
	return &testSituation{
		at: at.Create(),
		df: df.Create(),
		w:  w,
	}
}

func DummySituation() SituationChecker {
	at := &status.StatusData{
		Lv:            50,
		Types:         []types.Type{types.Fire},
		HP:            999,
		Attack:        100,
		Defense:       20,
		SpAttack:      200,
		SpDefense:     40,
		Speed:         180,
		AttackRank:    0,
		DefenseRank:   0,
		SpAttackRank:  0,
		SpDefenseRank: 0,
		SpeedRank:     0,
	}
	df := &status.StatusData{
		Lv:            50,
		Types:         []types.Type{types.Water},
		HP:            500,
		Attack:        20,
		Defense:       40,
		SpAttack:      60,
		SpDefense:     10,
		Speed:         120,
		AttackRank:    0,
		DefenseRank:   0,
		SpAttackRank:  0,
		SpDefenseRank: 0,
		SpeedRank:     0,
	}
	return &testSituation{
		at: at.Create(),
		df: df.Create(),
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

func Test_Parts(t *testing.T) {
	if Remote.IsContact() {
		t.Error()
	}
	if WaveMotion.IsContact() {
		t.Error()
	}
	if !Contact.IsContact() {
		t.Error()
	}
	if !Knuckle.IsContact() {
		t.Error()
	}
	if !Fang.IsContact() {
		t.Error()
	}
}
