package damage

import (
	"damagecalculator/domain/stats"
	"reflect"
	"testing"
)

func Test_CalcDamage(t *testing.T) {
	l := stats.NewLevel(50)
	s := SkillPower(100)
	var at uint = 182
	var df uint = 55

	actual := calcDamage(l, s, at, df)
	expect := Damage(147)
	if actual != expect {
		t.Errorf("Damage Error %d %d", expect, actual)
	}

}
func Test_CalcDamageArray(t *testing.T) {
	l := stats.NewLevel(50)
	s := SkillPower(100)
	var at uint = 182
	var df uint = 55
	a := calcDamageArray(l, s, at, df)
	e := []Damage{124, 126, 127, 129, 130, 132, 133, 135, 136, 138, 139, 141, 142, 144, 145, 147}

	if len(a) != 16 {
		t.Errorf("Illegal Length %d\n", len(a))
	}
	if !reflect.DeepEqual(a, e) {
		t.Errorf("Actual %v", a)
		t.Errorf("Expect %v", e)
	}
}
