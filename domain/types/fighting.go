package types

type fighting struct {
}

func (t fighting) Magnification(defense magnifier) Effective {
	switch defense.(type) {
	case *poison, *flying, *psychic, *bug, *fairy:
		return NotVeryEffective()
	case *normal, *ice, *rock, *dark, *steel:
		return SuperEffective()
	case *ghost:
		return NoEffective()
	}
	return FlatEffective()
}
