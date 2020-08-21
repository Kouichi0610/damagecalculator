package situation

import (
	"damagecalculator/domain/ability"
	"damagecalculator/domain/field"
	"damagecalculator/domain/item"
	"damagecalculator/domain/move"
	"damagecalculator/domain/status"
)

type SituationData struct {
	Move              *move.MoveFactory
	Attacker, Defender *status.StatusData
	Weather            field.Weather
	Field              field.Field

	AttackerAbility ability.AbilityBuilder
	DefenderAbility ability.AbilityBuilder

	AttackerItem item.ItemCreator
	DefenderItem item.ItemCreator

	IsCritical bool
}

func NewSituationData() *SituationData {
	return &SituationData{
		Weather:         field.NoWeather,
		Field:           field.NoField,
		AttackerAbility: &ability.NoAbilityData{},
		DefenderAbility: &ability.NoAbilityData{},
		AttackerItem:    &item.NoItem{},
		DefenderItem:    &item.NoItem{},
		IsCritical:      false,
	}
}

func (s *SituationData) Create() (SituationChecker, error) {
	aa := s.AttackerAbility.Create()
	da := s.DefenderAbility.Create()
	af := ability.NewAbilityField(aa, da)

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
