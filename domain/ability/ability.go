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

	OwnerSpeedCorrectorGenerator interface {
		Create(name string) OwnerSpeedCorrector
	}
	OtherSpeedCorrectorGenerator interface {
		Create(name string) OtherSpeedCorrector
	}

	// TODO:presentation/serverのstruct定義をdomain or usecaseに任せる
	OwnerSpeedCorrector struct {
		Comment            string  `json:"comment"`
		Rank               int     `json:"rank"`
		StatsMagnification float64 `json:"magnification"`
	}
	OtherSpeedCorrector struct {
		Comment            string  `json:"comment"`
		Rank               int     `json:"rank"`
		StatsMagnification float64 `json:"magnification"`
	}
)

func NewAbilityField(at, df detail.AbilityBuilder) AbilityField {
	a := at.Create()
	d := df.Create()
	return detail.NewAbilityField(a, d).(AbilityField)
}

func NewOwnerSpeedCorrectorGenerator() OwnerSpeedCorrectorGenerator {
	return new(ownerSpeedCorrectorGenerator)
}
func NewOtherSpeedCorrectorGenerator() OtherSpeedCorrectorGenerator {
	return new(otherSpeedCorrectorGenerator)
}
