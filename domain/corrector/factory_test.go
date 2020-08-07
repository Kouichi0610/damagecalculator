package corrector

import (
	"testing"
)

func Test_Factories(t *testing.T) {
	c := NewPower(1, 2)
	if c.caterogy() != Power {
		t.Error()
	}
	c = NewAttack(1, 2)
	if c.caterogy() != Attack {
		t.Error()
	}
	c = NewDefense(1, 2)
	if c.caterogy() != Defense {
		t.Error()
	}
	c = NewDamage(1, 2)
	if c.caterogy() != Damage {
		t.Error()
	}
}
