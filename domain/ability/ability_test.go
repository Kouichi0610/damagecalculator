package ability

import (
	"damagecalculator/domain/ability/detail"
	"testing"
)

func Test_生成(t *testing.T) {
	a := new(detail.MoldBreakderBuilder)
	b := new(detail.MimicryBuilder)

	af := NewAbilityField(a, b)
	if af == nil {
		t.Error()
	}
}
