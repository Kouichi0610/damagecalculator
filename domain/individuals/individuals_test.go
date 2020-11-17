package individuals

import (
	"reflect"
	"testing"
)

func Test_StringToIndividuals(t *testing.T) {
	if !reflect.DeepEqual(ToIndividuals("Max"), New(31, 31, 31, 31, 31, 31)) {
		t.Error()
	}
	if !reflect.DeepEqual(ToIndividuals("Slowest"), New(31, 31, 31, 31, 31, 0)) {
		t.Error()
	}
	if !reflect.DeepEqual(ToIndividuals("Weakest"), New(31, 0, 31, 31, 31, 31)) {
		t.Error()
	}
	if !reflect.DeepEqual(ToIndividuals("WeakestSlowest"), New(31, 0, 31, 31, 31, 0)) {
		t.Error()
	}
	if !reflect.DeepEqual(ToIndividuals("その他"), New(31, 31, 31, 31, 31, 31)) {
		t.Error()
	}
}

func Test_Types(t *testing.T) {
	if !reflect.DeepEqual(Max.ToIndividuals(), New(31, 31, 31, 31, 31, 31)) {
		t.Error()
	}
	if !reflect.DeepEqual(Slowest.ToIndividuals(), New(31, 31, 31, 31, 31, 0)) {
		t.Error()
	}
	if !reflect.DeepEqual(Weakest.ToIndividuals(), New(31, 0, 31, 31, 31, 31)) {
		t.Error()
	}
	if !reflect.DeepEqual(WeakestSlowest.ToIndividuals(), New(31, 0, 31, 31, 31, 0)) {
		t.Error()
	}
}

func Test_Individuals(t *testing.T) {
	i := New(1, 2, 3, 4, 5, 6)
	if i.HP() != 1 {
		t.Error()
	}
	if i.Attack() != 2 {
		t.Error()
	}
	if i.Defense() != 3 {
		t.Error()
	}
	if i.SpAttack() != 4 {
		t.Error()
	}
	if i.SpDefense() != 5 {
		t.Error()
	}
	if i.Speed() != 6 {
		t.Error()
	}
}

func Test_Clamp(t *testing.T) {
	res := clamp(32)
	if res != 31 {
		t.Error()
	}
}
