package types

type rock struct {
}

func (t rock) Magnification(defense magnifier) Effective {
	switch defense.(type) {
	case *fighting, *ground, *steel:
		return notVeryEffective()
	case *fire, *ice, *flying, *bug:
		return superEffective()
	}
	return flatEffective()
}
