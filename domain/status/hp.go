package status

type HP struct {
	value uint
}

func NewHP(v uint) *HP {
	return &HP{
		value: v,
	}
}

func (hp *HP) Value() uint {
	return hp.value
}
