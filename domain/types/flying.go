package types

type flying struct {
}

func (t flying) Magnification(defense magnifier) Effective {
	switch defense.(type) {
	case *electric, *rock, *steel:
		return NotVeryEffective()
	case *grass, *fighting, *bug:
		return SuperEffective()
	}
	return FlatEffective()
}
