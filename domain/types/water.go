package types

type water struct {
}

func (t water) Magnification(defense magnifier) Effective {
	switch defense.(type) {
	case *water, *grass, *dragon:
		return NotVeryEffective()
	case *fire, *ground, *rock:
		return SuperEffective()
	}
	return FlatEffective()
}
