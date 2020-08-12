package item

import (
	"damagecalculator/domain/corrector"
	"damagecalculator/domain/types"
)

type powerCorrector interface {
	CorrectPower(at, df, sk *types.Types) corrector.Corrector
}

func defaultPowerCorrector() powerCorrector {
	return new(defaultPowerCorrectorImpl)
}
func (t *TypeCorrectData) createTypeCorrector() powerCorrector {
	return &typeCorrectorImpl{
		ty: t.Type,
		c:  corrector.NewPower(t.Scale),
	}
}

type defaultPowerCorrectorImpl struct {
}

func (t *defaultPowerCorrectorImpl) CorrectPower(at, df, sk *types.Types) corrector.Corrector {
	return corrector.NewPower(1.0)
}

type typeCorrectorImpl struct {
	ty types.Type
	c  corrector.Corrector
}

// わざのタイプが対象であれば補正を掛ける
func (t *typeCorrectorImpl) CorrectPower(at, df, sk *types.Types) corrector.Corrector {
	if sk.Has(t.ty) {
		return t.c
	}
	return corrector.NewPower(1.0)
}
