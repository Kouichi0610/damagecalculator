package power

import "testing"

func Test_Power(t *testing.T) {
	p := NewPower(999)
	if p != 999 {
		t.Error()
	}
	p = NewPower(0)
	if p != 0 {
		t.Error()
	}
}
