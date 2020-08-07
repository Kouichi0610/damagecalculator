package types

type fire struct {
}

func (t fire) Magnification(defense magnifier) Effective {
	switch defense.(type) {
	case *fire, *water, *rock, *dragon:
		return notVeryEffective()
	case *grass, *ice, *bug, *steel:
		return superEffective()
	}
	return flatEffective()
}
