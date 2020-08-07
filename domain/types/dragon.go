package types

type dragon struct {
}

func (t dragon) Magnification(defense magnifier) Effective {
	switch defense.(type) {
	case *steel:
		return notVeryEffective()
	case *dragon:
		return superEffective()
	case *fairy:
		return noEffective()
	}
	return flatEffective()
}
