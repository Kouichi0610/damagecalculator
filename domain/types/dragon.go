package types

type dragon struct {
}

func (t dragon) Magnification(defense magnifier) Effective {
	switch defense.(type) {
	case *steel:
		return NotVeryEffective()
	case *dragon:
		return SuperEffective()
	case *fairy:
		return NoEffective()
	}
	return FlatEffective()
}
