package types

type fighting struct {
}

func (t fighting) Magnification(defense magnifier) Effective {
	switch defense.(type) {
	case *poison, *flying, *psychic, *bug, *fairy:
		return notVeryEffective()
	case *normal, *ice, *rock, *dark, *steel:
		return superEffective()
	case *ghost:
		return noEffective()
	}
	return flatEffective()
}
