package statspattern

import (
	"damagecalculator/domain/stats"
	"damagecalculator/infra/local"
	"testing"
)

func Test_statsPattern(t *testing.T) {
	rp := local.Species()
	l := NewLoader(rp)

	stats, err := l.Get(50, "ツンデツンデ", "いじっぱり", stats.IndividualTypeMax)
	if err != nil {
		t.Error()
	}
	if checkError(stats.HP(), 136, 168) {
		t.Errorf("HP:%v", stats.HP())
	}
	if checkError(stats.Attack(), 166, 201) {
		t.Errorf("AT:%v", stats.Attack())
	}
	if checkError(stats.Defense(), 231, 263) {
		t.Errorf("DF:%v", stats.Defense())
	}
	if checkError(stats.SpAttack(), 65, 94) {
		t.Errorf("SA:%v", stats.SpAttack())
	}
	if checkError(stats.SpDefense(), 121, 153) {
		t.Errorf("SD:%v", stats.SpDefense())
	}
	if checkError(stats.Speed(), 33, 65) {
		t.Errorf("SP:%v", stats.Speed())
	}
}

// HPが必ず1になること
func Test_ヌケニン(t *testing.T) {
	rp := local.Species()
	l := NewLoader(rp)

	stats, err := l.Get(50, "ヌケニン", "いじっぱり", stats.IndividualTypeMax)
	if err != nil {
		t.Error()
	}
	if checkError(stats.HP(), 1, 1) {
		t.Errorf("HP:%v", stats.HP())
	}
}

func checkError(params []uint, start uint, end uint) bool {
	if len(params) != 64 {
		return true
	}
	if params[0] != start {
		return true
	}
	if params[63] != end {
		return true
	}
	return false
}
