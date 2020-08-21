package ability

import (
	"damagecalculator/domain/ability/detail"
	"damagecalculator/domain/ability/situation"
	"damagecalculator/domain/corrector"
	"damagecalculator/domain/move"
	"damagecalculator/domain/status"
)

// TODO:リポジトリ

type (
	AbilityField interface {
		CorrectAttackerStatus(situation.SituationChecker) *status.StatsCorrectors
		CorrectDefenderStatus(situation.SituationChecker) *status.StatsCorrectors
		Correctors(situation.SituationChecker) []corrector.Corrector
		RewriteMoveFactory(move.MoveFactory) *move.MoveFactory
	}
)

func NewAbilityField(at, df detail.AbilityBuilder) AbilityField {
	a := at.Create()
	d := df.Create()
	return detail.NewAbilityField(a, d).(AbilityField)
}
