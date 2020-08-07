package types

type ground struct {
}

func (t ground) Magnification(defense magnifier) Effective {
	switch defense.(type) {
	case *grass, *bug:
		return notVeryEffective()
	case *fire, *electric, *poison, *rock, *steel:
		return superEffective()
	case *flying:
		return noEffective()
	}
	return flatEffective()
}
