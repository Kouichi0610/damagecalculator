package item

import (
	"damagecalculator/domain/corrector"
	"damagecalculator/domain/types"
)

// TypeMatchCorrectData
type superEffectiveCorrectorImpl struct {
	c corrector.Corrector
}

func (t *SuperEffectiveCorrectData) createPowerCorrector() powerCorrector {
	return &superEffectiveCorrectorImpl{
		c: corrector.NewPower(t.Scale),
	}
}

func (t *superEffectiveCorrectorImpl) CorrectPower(at, df, sk *types.Types) corrector.Corrector {
	ef := sk.Magnification(df)
	if ef.IsSuper() {
		return t.c
	}
	return corrector.NewPower(1.0)
}
