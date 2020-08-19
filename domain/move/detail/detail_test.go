package detail

import (
	"damagecalculator/domain/move/attribute"
	"damagecalculator/domain/move/category"
	"damagecalculator/domain/move/count"
	"damagecalculator/domain/move/target"
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

// TODO:dummy situation用意
