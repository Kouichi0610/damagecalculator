package types

type poison struct {
}

func (t poison) Magnification(defense magnifier) Effective {
	switch defense.(type) {
	case *poison, *ground, *rock, *ghost:
		return NotVeryEffective()
	case *grass, *fairy:
		return SuperEffective()
	case *steel:
		return NoEffective()
	}
	return FlatEffective()
}
