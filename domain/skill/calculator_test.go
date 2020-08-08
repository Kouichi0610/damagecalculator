package skill

import (
	"reflect"
	"testing"
)

func Test_Calculator(t *testing.T) {
	dmg := calculate(50, 100, 182, 189)
	if dmg != 44 {
		t.Errorf("%d", dmg)
	}
	dmg = calculate(50, 120, 85, 70)
	if dmg != 66 {
		t.Errorf("%d", dmg)
	}
	dmg = calculate(50, 150, 185, 70)
	if dmg != 176 {
		t.Errorf("%d", dmg)
	}
}

func Test_CalculatorArray(t *testing.T) {
	d := calculateArray(50, 150, 185, 70)
	if !reflect.DeepEqual(d, []uint{149, 151, 153, 154, 156, 158, 160, 161, 163, 165, 167, 168, 170, 172, 174, 176}) {
		t.Errorf("%v", d)
	}
}
