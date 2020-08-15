package correct

import (
	"damagecalculator/domain/ability/situation"
	"damagecalculator/domain/corrector"
)

type PowerCorrector interface {
	Correct(isAttacker bool, st situation.SituationChecker) corrector.Corrector
}
