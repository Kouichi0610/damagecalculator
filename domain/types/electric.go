package types

type electric struct {
}

func (t electric) Magnification(defense magnifier) Effective {
	switch defense.(type) {
	case *electric, *grass, *dragon:
		return NotVeryEffective()
	case *water, *flying:
		return SuperEffective()
	case *ground:
		return NoEffective()
	}
	return FlatEffective()
}
