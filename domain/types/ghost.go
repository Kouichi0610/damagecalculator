package types

type ghost struct {
}

func (t ghost) Magnification(defense magnifier) Effective {
	switch defense.(type) {
	case *dark:
		return NotVeryEffective()
	case *psychic, *ghost:
		return SuperEffective()
	case *normal:
		return NoEffective()
	}
	return FlatEffective()
}
