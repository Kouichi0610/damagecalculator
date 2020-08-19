package attribute

import "testing"

func Test_Attributes_Action(t *testing.T) {
	type expect struct {
		contact, knuckle, fang, wave, sound bool
	}
	expects := map[Attribute]expect{
		NewAttribute(Remote, NoAttribute):     {false, false, false, false, false},
		NewAttribute(Contact, NoAttribute):    {true, false, false, false, false},
		NewAttribute(Knuckle, NoAttribute):    {true, true, false, false, false},
		NewAttribute(Fang, NoAttribute):       {true, false, true, false, false},
		NewAttribute(WaveMotion, NoAttribute): {false, false, false, true, false},
		NewAttribute(Sound, NoAttribute):      {false, false, false, false, true},
	}

	for a, e := range expects {
		if a.IsContact() != e.contact {
			t.Errorf("%v", a)
		}
		if a.IsKnuckle() != e.knuckle {
			t.Errorf("%v", a)
		}
		if a.IsFang() != e.fang {
			t.Errorf("%v", a)
		}
		if a.IsWaveMotion() != e.wave {
			t.Errorf("%v", a)
		}
		if a.IsSound() != e.sound {
			t.Errorf("%v", a)
		}
	}
}
func Test_Attributes_Effect(t *testing.T) {
	type expect struct {
		appendEffect, recoil bool
	}
	expects := map[Attribute]expect{
		NewAttribute(Remote, NoAttribute):  {false, false},
		NewAttribute(Remote, AppendEffect): {true, false},
		NewAttribute(Remote, Recoil):       {false, true},
	}

	for a, e := range expects {
		if a.HasAppendEffect() != e.appendEffect {
			t.Errorf("%v", a)
		}
		if a.HasRecoil() != e.recoil {
			t.Errorf("%v", a)
		}
	}
}
