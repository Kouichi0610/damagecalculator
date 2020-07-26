package types

type psychic struct {
}

func (t psychic) Magnification(defense magnifier) Effective {
	switch defense.(type) {
	case *psychic, *steel:
		return NotVeryEffective()
	case *fighting, *poison:
		return SuperEffective()
	case *dark:
		return NoEffective()
	}
	return FlatEffective()
}
