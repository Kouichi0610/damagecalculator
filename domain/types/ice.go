package types

type ice struct {
}

func (t ice) Magnification(defense magnifier) Effective {
	switch defense.(type) {
	case *fire, *water, *ice, *steel:
		return NotVeryEffective()
	case *grass, *ground, *flying, *dragon:
		return SuperEffective()
	}
	return FlatEffective()
}
