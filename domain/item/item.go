package item

import (
	"damagecalculator/domain/corrector"
	"damagecalculator/domain/status"
	"damagecalculator/domain/types"
)

type (
	Item interface {
		statsCorrector
		powerCorrector
	}
)

type item struct {
	s statsCorrector
	p powerCorrector
}

func (i *item) Correct() *status.StatsCorrectors {
	return i.s.Correct()
}
func (i *item) CorrectPower(at, df, sk *types.Types) corrector.Corrector {
	return i.p.CorrectPower(at, df, sk)
}
