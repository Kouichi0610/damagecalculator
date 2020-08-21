package accuracy

import "testing"

func Test_Accuracy(t *testing.T) {
	a := NewAccuracy(100)
	if a != 100 {
		t.Error()
	}
	a = NewAccuracy(0)
	if a != 0 {
		t.Error()
	}
	a = NewAccuracy(101)
	if a != 100 {
		t.Error()
	}
}
