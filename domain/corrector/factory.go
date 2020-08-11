package corrector

func NewPower(numer, denom uint) Corrector {
	return newCorrector(Power, drop4_pick5, numer, denom)
}

func NewAttack(numer, denom uint) Corrector {
	return newCorrector(Attack, drop4_pick5, numer, denom)
}

func NewDefense(numer, denom uint) Corrector {
	return newCorrector(Defense, drop4_pick5, numer, denom)
}

func NewDamage(numer, denom uint) Corrector {
	return newCorrector(Damage, drop5_pick5over, numer, denom)
}

func NewTypeMatch(numer, denom uint) Corrector {
	return newCorrector(TypeMatch, drop5_pick5over, numer, denom)
}
func NewCritical(numer, denom uint) Corrector {
	return newCorrector(Critical, drop5_pick5over, numer, denom)
}
