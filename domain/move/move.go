package move

import (
	"damagecalculator/domain/corrector"
	"damagecalculator/domain/move/attribute"
	"damagecalculator/domain/move/detail"
	"damagecalculator/domain/status"
	"damagecalculator/domain/types"
)

//　わざ
type Move interface {
	Correctors(detail.SituationChecker) []corrector.Corrector
	Types(detail.SituationChecker) *types.Types
	Power(detail.SituationChecker) uint
	PickStats(detail.SituationChecker) (at, df *status.RankedValue)
	Calculate(level, power, attack, defense uint) []uint
	Attribute() attribute.Attribute
}
