package types

type grass struct {
}

func (t grass) Magnification(defense magnifier) Effective {
	switch defense.(type) {
	case *fire, *grass, *poison, *flying, *bug, *dragon, *steel:
		return NotVeryEffective()
	case *water, *ground, *rock:
		return SuperEffective()
	}
	return FlatEffective()
}
