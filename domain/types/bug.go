package types

type bug struct {
}

func (t bug) Magnification(defense magnifier) Effective {
	switch defense.(type) {
	case *fire, *fighting, *poison, *flying, *ghost, *steel, *fairy:
		return notVeryEffective()
	case *grass, *psychic, *dark:
		return superEffective()
	}
	return flatEffective()
}
