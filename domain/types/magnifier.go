package types

type magnifier interface {
	Magnification(defense magnifier) Effective
}
