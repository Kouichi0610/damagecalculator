// ダメージ計算API
package damage

import (
	DOMAIN "damagecalculator/domain/damage"
)

// damagerateservice
// enemylistservice

// ランク変動するとくせい

//
type (
	DamageRateService interface {
		Calculate(
			level int,
			name string,

		) DOMAIN.DamageRate
	}
)
