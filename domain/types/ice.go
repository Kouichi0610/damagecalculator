package types

type ice struct {
}

func (t ice) Magnification(defense magnifier) Effective {
	switch defense.(type) {
	case *fire, *water, *ice, *steel:
		return notVeryEffective()
	case *grass, *ground, *flying, *dragon:
		return superEffective()
	}
	return flatEffective()
}
