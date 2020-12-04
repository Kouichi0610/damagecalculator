package status

import "fmt"

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

func (hp *HP) String() string {
	return fmt.Sprintf("%d", hp.value)
}
