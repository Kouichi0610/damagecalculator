package types

type grass struct {
}

func (t grass) Magnification(defense magnifier) Effective {
	switch defense.(type) {
	case *fire, *grass, *poison, *flying, *bug, *dragon, *steel:
		return notVeryEffective()
	case *water, *ground, *rock:
		return superEffective()
	}
	return flatEffective()
}
