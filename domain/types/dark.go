package types

type dark struct {
}

func (t dark) Magnification(defense magnifier) Effective {
	switch defense.(type) {
	case *fighting, *dark, *fairy:
		return notVeryEffective()
	case *psychic, *ghost:
		return superEffective()
	}
	return flatEffective()
}
