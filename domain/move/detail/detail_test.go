package detail

import (
	"damagecalculator/domain/field"
	"damagecalculator/domain/move/attribute"
	"damagecalculator/domain/move/category"
	"damagecalculator/domain/move/count"
	"damagecalculator/domain/move/target"
	"damagecalculator/domain/status"
	"damagecalculator/domain/types"
	"reflect"
	"testing"
)

func Test_生成型チェック(t *testing.T) {
	expects := map[Detail]reflect.Type{
		Default:     reflect.TypeOf(new(defaultMove)),
		SeismicToss: reflect.TypeOf(new(seismicToss)),
		WeatherBall: reflect.TypeOf(new(weatherBall)),
		GyroBall:    reflect.TypeOf(new(gyroBall)),
		HeavySlam:   reflect.TypeOf(new(heavySlam)),
	}
	for detail, expect := range expects {
		mv := newTestMove(detail)
		actual := reflect.TypeOf(mv)
		if actual != expect {
			t.Errorf("actual[%v] expect[%v]", actual, expect)
		}
	}
}

func Test_DefaultMove機能テスト(t *testing.T) {
	st := dummySituation()
	mv := newTestMove(Default).(*defaultMove)

	c := mv.Correctors(st)
	if len(c) != 1 {
		t.Error()
	}
	if c[0].Correct(100) != 100 {
		t.Error()
	}
	ty := mv.Types(st)
	if !ty.Equal(types.NewTypes(types.Fire)) {
		t.Errorf("%s", ty.String())
	}
	if mv.Power(st) != 100 {
		t.Error()
	}
	at, df := mv.PickStats(st)
	if at.Value() != 100 {
		t.Errorf("%d", at.Value())
	}
	if df.Value() != 20 {
		t.Errorf("%d", df.Value())
	}

	dmgs := mv.Calculate(50, 150, 185, 70)
	if !reflect.DeepEqual(dmgs, []uint{149, 151, 153, 154, 156, 158, 160, 161, 163, 165, 167, 168, 170, 172, 174, 176}) {
		t.Errorf("%v", dmgs)
	}
	attr := mv.Attribute()
	if !attr.IsContact() {
		t.Error()
	}
	if !attr.HasRecoil() {
		t.Error()
	}
}

func Test_ヘビーボンバー(t *testing.T) {
	attackCnt, _ := count.NewAttackCount(1, 1)
	res, _ := NewMove(
		0,
		types.Steel,
		100,
		category.Physical,
		attackCnt,
		target.Select,
		attribute.NewAttribute(attribute.Contact, attribute.NoAttribute),
		HeavySlam)
	mv := res.(*heavySlam)
	st := dummySituation()
	if mv.Power(st) != 120 {
		t.Errorf("%d", mv.Power(st))
	}
}

func Test_heavySlamPower(t *testing.T) {
	type weightData struct {
		a, d   float64
		expect uint
	}

	tests := []weightData{
		{100, 20, 120},
		{99, 20, 100},
		{100, 25, 100},
		{99, 25, 80},
		{90, 30, 80},
		{89, 30, 60},
		{100, 50, 60},
		{99, 50, 40},
	}
	for _, d := range tests {
		a := status.NewWeight(d.a)
		b := status.NewWeight(d.d)
		act := heavySlamPower(a, b)
		if act != d.expect {
			t.Errorf("failed %v act:%d", d, act)
		}
	}
}

func Test_ジャイロボール(t *testing.T) {
	attackCnt, _ := count.NewAttackCount(1, 1)
	res, _ := NewMove(
		0,
		types.Steel,
		100,
		category.Physical,
		attackCnt,
		target.Select,
		attribute.NewAttribute(attribute.Contact, attribute.NoAttribute),
		GyroBall)
	mv := res.(*gyroBall)
	st := dummySituation()

	if mv.Power(st) != 17 {
		t.Errorf("%d", mv.Power(st))
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

func Test_ウェザーボール(t *testing.T) {
	attackCnt, _ := count.NewAttackCount(1, 1)
	res, _ := NewMove(
		100,
		types.Normal,
		80,
		category.Special,
		attackCnt,
		target.Select,
		attribute.NewAttribute(attribute.Remote, attribute.NoAttribute),
		WeatherBall)
	mv := res.(*weatherBall)

	const baseDmg = 100
	type expectData struct {
		correctDmg uint
		ty         *types.Types
	}
	expects := map[field.Weather]*expectData{
		field.NoWeather: &expectData{100, types.NewTypes(types.Normal)},
		field.Sunny:     &expectData{200, types.NewTypes(types.Fire)},
		field.Rainy:     &expectData{200, types.NewTypes(types.Water)},
		field.Snow:      &expectData{200, types.NewTypes(types.Ice)},
		field.SandStorm: &expectData{200, types.NewTypes(types.Rock)},
	}

	for w, expect := range expects {
		st := weatherSituation(w)
		if !mv.Types(st).Equal(expect.ty) {
			t.Errorf("%v", expect)
		}
		c := mv.Correctors(st)
		if len(c) != 1 {
			t.Error()
		}
		if c[0].Correct(baseDmg) != expect.correctDmg {
			t.Errorf("%v", expect)
		}
	}
}

func Test_ちきゅうなげ(t *testing.T) {
	attackCnt, _ := count.NewAttackCount(1, 1)
	res, _ := NewMove(
		100,
		types.Fire,
		80,
		category.Physical,
		attackCnt,
		target.Select,
		attribute.NewAttribute(attribute.Contact, attribute.Recoil),
		SeismicToss)
	mv := res.(*seismicToss)

	// ダメージ計算はレベル固定となること
	dmgs := mv.Calculate(50, 1, 1, 1)
	if !reflect.DeepEqual(dmgs, []uint{50}) {
		t.Error()
	}
}

func Test_エラーチェック_Detail(t *testing.T) {
	attackCnt, _ := count.NewAttackCount(1, 1)
	mv, err := NewMove(
		100,
		types.Fire,
		80,
		category.Physical,
		attackCnt,
		target.Select,
		attribute.NewAttribute(attribute.Contact, attribute.Recoil),
		Detail(100))
	if mv != nil {
		t.Error()
	}
	if err == nil {
		t.Error()
	}
}

func Test_エラーチェック_Category(t *testing.T) {
	attackCnt, _ := count.NewAttackCount(1, 1)
	mv, err := NewMove(
		100,
		types.Fire,
		80,
		category.DamageCategory(100),
		attackCnt,
		target.Select,
		attribute.NewAttribute(attribute.Contact, attribute.Recoil),
		Default)
	if mv != nil {
		t.Error()
	}
	if err == nil {
		t.Error()
	}
}

func newTestMove(d Detail) interface{} {
	attackCnt, _ := count.NewAttackCount(1, 1)
	res, _ := NewMove(
		100,
		types.Fire,
		80,
		category.Physical,
		attackCnt,
		target.Select,
		attribute.NewAttribute(attribute.Contact, attribute.Recoil),
		d)
	return res
}

func weatherSituation(w field.Weather) SituationChecker {
	st := dummySituation().(*testSituation)
	st.w = w
	return st
}
func dummySituation() SituationChecker {
	at := &status.StatusData{
		Lv:            50,
		Types:         []types.Type{types.Fire},
		HP:            999,
		Attack:        100,
		Defense:       200,
		SpAttack:      300,
		SpDefense:     400,
		Speed:         180,
		AttackRank:    0,
		DefenseRank:   0,
		SpAttackRank:  0,
		SpDefenseRank: 0,
		SpeedRank:     0,
		Weight:        100,
	}
	df := &status.StatusData{
		Lv:            50,
		Types:         []types.Type{types.Water},
		HP:            500,
		Attack:        10,
		Defense:       20,
		SpAttack:      30,
		SpDefense:     40,
		Speed:         120,
		AttackRank:    0,
		DefenseRank:   0,
		SpAttackRank:  0,
		SpDefenseRank: 0,
		SpeedRank:     0,
		Weight:        20,
	}
	return &testSituation{
		at: at.Create(),
		df: df.Create(),
	}
}

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
