package types

type fairy struct {
}

func (t fairy) Magnification(defense magnifier) Effective {
	switch defense.(type) {
	case *fire, *poison, *steel:
		return NotVeryEffective()
	case *fighting, *dragon, *dark:
		return SuperEffective()
	}
	return FlatEffective()
}
