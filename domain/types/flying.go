package types

type flying struct {
}

func (t flying) Magnification(defense magnifier) Effective {
	switch defense.(type) {
	case *electric, *rock, *steel:
		return notVeryEffective()
	case *grass, *fighting, *bug:
		return superEffective()
	}
	return flatEffective()
}
