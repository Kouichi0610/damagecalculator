package skills

import "damagecalculator/domain/damage"

func CreateSample() damage.DamageCalculator {
	return new(GyroBall)
}
