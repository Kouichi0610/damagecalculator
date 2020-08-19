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

func Test_生成型チェック(t *testing.T) {
	expects := map[Detail]reflect.Type{
		Default: reflect.TypeOf(new(defaultMove)),
	}
	for detail, expect := range expects {
		mv := newTestMove(detail)
		actual := reflect.TypeOf(mv)
		if actual != expect {
			t.Errorf("actual[%v] expect[%v]", actual, expect)
		}
	}
}

/*

func (m *defaultMove) Attribute() attribute.Attribute {
	return attribute.NewAttribute(attribute.Fang, attribute.NoAttribute)
}

*/

func Test_DefaultMove機能テスト(t *testing.T) {
	st := dummySituation()
	mv := newTestMove(Default).(*defaultMove)

	if mv.Correctors(st) != nil {
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

func dummySituation() SituationChecker {
	at := &status.StatusData{
		Lv:            50,
		Types:         []types.Type{types.Fire},
		HP:            999,
		Attack:        100,
		Defense:       200,
		SpAttack:      300,
		SpDefense:     400,
		Speed:         500,
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
		Speed:         50,
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
