package stats

import (
	"reflect"
	"testing"
)

func Test_IndividualRange(t *testing.T) {
	min := newIndividual(0)
	if min != 0 {
		t.Error()
	}
	max := newIndividual(32)
	if max != 31 {
		t.Error()
	}
	value := newIndividual(15)
	if value != 15 {
		t.Error()
	}
}

func Test_IndividualTypes(t *testing.T) {
	actual := IndividualTypeMax.Create()
	if !reflect.DeepEqual(actual, newIndividualStats(31, 31, 31, 31, 31, 31)) {
		t.Error()
	}

	actual = IndividualTypeSlowest.Create()
	if !reflect.DeepEqual(actual, newIndividualStats(31, 31, 31, 31, 31, 0)) {
		t.Error()
	}
	actual = IndividualTypeWeakest.Create()
	if !reflect.DeepEqual(actual, newIndividualStats(31, 0, 31, 31, 31, 31)) {
		t.Error()
	}
}

func Test_IndividualStats(t *testing.T) {
	i := newIndividualStats(0, 1, 2, 3, 4, 5)
	if i.HP() != 0 {
		t.Error()
	}
	if i.Attack() != 1 {
		t.Error()
	}
	if i.Defense() != 2 {
		t.Error()
	}
	if i.SpAttack() != 3 {
		t.Error()
	}
	if i.SpDefense() != 4 {
		t.Error()
	}
	if i.Speed() != 5 {
		t.Error()
	}
}
