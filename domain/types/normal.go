package types

type normal struct {
}

func (t normal) Magnification(defense magnifier) Effective {
	switch defense.(type) {
	case *rock, *steel:
		return NotVeryEffective()
	case *ghost:
		return NoEffective()
	}
	return FlatEffective()
}
