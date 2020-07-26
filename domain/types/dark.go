package types

type dark struct {
}

func (t dark) Magnification(defense magnifier) Effective {
	switch defense.(type) {
	case *fighting, *dark, *fairy:
		return NotVeryEffective()
	case *psychic, *ghost:
		return SuperEffective()
	}
	return FlatEffective()
}
