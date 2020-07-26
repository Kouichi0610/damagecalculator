package types

type steel struct {
}

func (t steel) Magnification(defense magnifier) Effective {
	switch defense.(type) {
	case *fire, *water, *electric, *steel:
		return NotVeryEffective()
	case *ice, *rock, *fairy:
		return SuperEffective()
	}
	return FlatEffective()
}
