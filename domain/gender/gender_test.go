package gender

import "testing"

func Test_Relations(t *testing.T) {
	actuals := []struct {
		a, b                                            Gender
		unrelated, couldbe, alwaysSame, alwaysDifferent bool
	}{
		{MaleFemale, MaleFemale, false, true, false, false},
		{MaleFemale, MaleOnly, false, true, false, false},
		{MaleFemale, FemaleOnly, false, true, false, false},
		{MaleFemale, Unknown, true, false, false, false},
		{MaleOnly, MaleFemale, false, true, false, false},
		{MaleOnly, MaleOnly, false, false, true, false},
		{MaleOnly, FemaleOnly, false, false, false, true},
		{MaleOnly, Unknown, true, false, false, false},
		{FemaleOnly, MaleFemale, false, true, false, false},
		{FemaleOnly, MaleOnly, false, false, false, true},
		{FemaleOnly, FemaleOnly, false, false, true, false},
		{FemaleOnly, Unknown, true, false, false, false},
		{Unknown, MaleFemale, true, false, false, false},
		{Unknown, MaleOnly, true, false, false, false},
		{Unknown, FemaleOnly, true, false, false, false},
		{Unknown, Unknown, true, false, false, false},
	}

	for _, actual := range actuals {
		expect := NewRelation(actual.a, actual.b)
		if expect.Unrelated() != actual.unrelated {
			t.Errorf("%v %v", expect, actual)
		}
		if expect.CouldBeDifferent() != actual.couldbe {
			t.Errorf("%v %v", expect, actual)
		}
		if expect.AlwaysSame() != actual.alwaysSame {
			t.Errorf("%v %v", expect, actual)
		}
		if expect.AlwaysDifferent() != actual.alwaysDifferent {
			t.Errorf("%v %v", expect, actual)
		}
	}
}

func Test_Gender_生成(t *testing.T) {
	r := &relation{
		unrelated:        false,
		couldBeDifferent: true,
		alwaysSame:       true,
		alwaysDifferent:  false,
	}
	if r.Unrelated() {
		t.Error()
	}
	if !r.CouldBeDifferent() {
		t.Error()
	}
	if !r.AlwaysSame() {
		t.Error()
	}
	if r.AlwaysDifferent() {
		t.Error()
	}
}
