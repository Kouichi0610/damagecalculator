package types

type ghost struct {
}

func (t ghost) Magnification(defense magnifier) Effective {
	switch defense.(type) {
	case *dark:
		return notVeryEffective()
	case *psychic, *ghost:
		return superEffective()
	case *normal:
		return noEffective()
	}
	return flatEffective()
}
