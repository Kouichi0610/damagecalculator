package types

type electric struct {
}

func (t electric) Magnification(defense magnifier) Effective {
	switch defense.(type) {
	case *electric, *grass, *dragon:
		return notVeryEffective()
	case *water, *flying:
		return superEffective()
	case *ground:
		return noEffective()
	}
	return flatEffective()
}
