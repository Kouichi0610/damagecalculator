package corrector

import "damagecalculator/domain/fixed"

func NewPower(f float64) Corrector {
	fu, _ := fixed.NewFixPN(f, fixed.Drop4Pick5)
	return &corrector{
		t: Power,
		f: fu,
	}
}

func NewAttack(f float64) Corrector {
	fu, _ := fixed.NewFixPN(f, fixed.Drop4Pick5)
	return &corrector{
		t: Attack,
		f: fu,
	}
}

func NewDefense(f float64) Corrector {
	fu, _ := fixed.NewFixPN(f, fixed.Drop4Pick5)
	return &corrector{
		t: Defense,
		f: fu,
	}
}

func NewDamage(f float64) Corrector {
	fu, _ := fixed.NewFixPN(f, fixed.Drop5Pick5Over)
	return &corrector{
		t: Damage,
		f: fu,
	}
}

func NewTypeMatch(f float64) Corrector {
	fu, _ := fixed.NewFixPN(f, fixed.Drop5Pick5Over)
	return &corrector{
		t: TypeMatch,
		f: fu,
	}
}
func NewCritical(f float64) Corrector {
	fu, _ := fixed.NewFixPN(f, fixed.Drop5Pick5Over)
	return &corrector{
		t: Critical,
		f: fu,
	}
}
func newMultiTarget(f float64) Corrector {
	fu, _ := fixed.NewFixPN(f, fixed.Drop5Pick5Over)
	return &corrector{
		t: MultiTarget,
		f: fu,
	}
}
func newBurn(f float64) Corrector {
	fu, _ := fixed.NewFixPN(f, fixed.Drop5Pick5Over)
	return &corrector{
		t: Burn,
		f: fu,
	}
}
