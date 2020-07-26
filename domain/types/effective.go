package types

type Effective float32

func FlatEffective() Effective {
	return 1.0
}
func SuperEffective() Effective {
	return 2.0
}
func NotVeryEffective() Effective {
	return 0.5
}
func NoEffective() Effective {
	return 0.0
}
