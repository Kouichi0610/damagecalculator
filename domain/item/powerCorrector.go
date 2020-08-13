package item

import (
	"damagecalculator/domain/corrector"
	"damagecalculator/domain/types"
)

type powerCorrectorImpl struct {
	c corrector.Corrector
}

func (t *PowerCorrectData) createPowerCorrector() powerCorrector {
	return &powerCorrectorImpl{
		c: corrector.NewPower(t.Scale),
	}
}

func (t *powerCorrectorImpl) CorrectPower(at, df, sk *types.Types) corrector.Corrector {
	return t.c
}
