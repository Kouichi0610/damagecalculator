package types

type ground struct {
}

func (t ground) Magnification(defense magnifier) Effective {
	switch defense.(type) {
	case *grass, *bug:
		return NotVeryEffective()
	case *fire, *electric, *poison, *rock, *steel:
		return SuperEffective()
	case *flying:
		return NoEffective()
	}
	return FlatEffective()
}
