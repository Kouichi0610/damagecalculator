package corrector

import (
	"testing"
)

func Test_Factories(t *testing.T) {
	c := NewPower(1, 2)
	if c.Caterogy() != Power {
		t.Error()
	}
	c = NewAttack(1, 2)
	if c.Caterogy() != Attack {
		t.Error()
	}
	c = NewDefense(1, 2)
	if c.Caterogy() != Defense {
		t.Error()
	}
	c = NewDamage(1, 2)
	if c.Caterogy() != Damage {
		t.Error()
	}
}
