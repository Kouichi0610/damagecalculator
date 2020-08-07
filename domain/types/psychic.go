package types

type psychic struct {
}

func (t psychic) Magnification(defense magnifier) Effective {
	switch defense.(type) {
	case *psychic, *steel:
		return notVeryEffective()
	case *fighting, *poison:
		return superEffective()
	case *dark:
		return noEffective()
	}
	return flatEffective()
}
