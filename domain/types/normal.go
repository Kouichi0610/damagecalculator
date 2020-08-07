package types

type normal struct {
}

func (t normal) Magnification(defense magnifier) Effective {
	switch defense.(type) {
	case *rock, *steel:
		return notVeryEffective()
	case *ghost:
		return noEffective()
	}
	return flatEffective()
}
