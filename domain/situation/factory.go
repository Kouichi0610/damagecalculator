package situation

import (
	"damagecalculator/domain/ability"
	"damagecalculator/domain/ability/detail"
	"damagecalculator/domain/field"
	"damagecalculator/domain/item"
	"damagecalculator/domain/move"
	"damagecalculator/domain/status"
)

type SituationData struct {
	Move               *move.MoveFactory
	Attacker, Defender *status.StatusData
	Weather            field.Weather
	Field              field.Field

	AttackerAbility detail.AbilityBuilder
	DefenderAbility detail.AbilityBuilder

	AttackerItem item.ItemCreator
	DefenderItem item.ItemCreator

	IsCritical bool
}

func NewSituationData() *SituationData {
	return &SituationData{
		Weather:         field.NoWeather,
		Field:           field.NoField,
		AttackerAbility: new(detail.NoEffectBuilder),
		DefenderAbility: new(detail.NoEffectBuilder),
		AttackerItem:    &item.NoItem{},
		DefenderItem:    &item.NoItem{},
		IsCritical:      false,
	}
}

func (s *SituationData) Create() (SituationChecker, error) {
	af := ability.NewAbilityField(s.AttackerAbility, s.DefenderAbility)

	// とくせいによるわざ情報上書き
	sd := af.RewriteMoveFactory(*s.Move)
	sk, err := sd.Create()
	if err != nil {
		return nil, err
	}
	return &situation{
		at:         s.Attacker.Create(),
		df:         s.Defender.Create(),
		sk:         sk,
		mv:         field.NewFields(s.Field, s.Weather),
		atItem:     s.AttackerItem.Create(true),
		dfItem:     s.DefenderItem.Create(false),
		abilities:  af,
		isCritical: s.IsCritical,
	}, nil
}
