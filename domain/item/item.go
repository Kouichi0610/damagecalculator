package item

import (
	_ "damagecalculator/domain/corrector"
	_ "damagecalculator/domain/status"
	_ "damagecalculator/domain/types"
)

type (
	Item interface {
		// statsCorrector
		// powerCorrector

		// CorrectStats() []*status.StatsCorrectors
		//	CorrectPower(sk, df *types.Types) corrector.Corrector
	}
)

type item struct {
	isAttacker bool
	//s          statsCorrector
	//p          powerCorrector
}
