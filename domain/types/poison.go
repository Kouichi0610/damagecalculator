package types

type poison struct {
}

func (t poison) Magnification(defense magnifier) Effective {
	switch defense.(type) {
	case *poison, *ground, *rock, *ghost:
		return notVeryEffective()
	case *grass, *fairy:
		return superEffective()
	case *steel:
		return noEffective()
	}
	return flatEffective()
}
