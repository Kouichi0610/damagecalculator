package types

type fairy struct {
}

func (t fairy) Magnification(defense magnifier) Effective {
	switch defense.(type) {
	case *fire, *poison, *steel:
		return notVeryEffective()
	case *fighting, *dragon, *dark:
		return superEffective()
	}
	return flatEffective()
}
