package status

type HP struct {
	value uint
}

func NewHP(v uint) *HP {
	return &HP{
		value: v,
	}
}
