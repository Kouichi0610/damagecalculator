/*
	ダメージ計算に必要な場の状態
*/
package situation

import (
	"damagecalculator/domain/ability"
	"damagecalculator/domain/corrector"
	"damagecalculator/domain/field"
	"damagecalculator/domain/item"
	"damagecalculator/domain/move"
	"damagecalculator/domain/move/attribute"
	"damagecalculator/domain/status"
	"damagecalculator/domain/types"
)

type SituationChecker interface {
	Attacker() status.StatusChecker
	Defender() status.StatusChecker
	Correctors() []corrector.Corrector
	IsWeather(field.Weather) bool
	IsField(field.Field) bool
	IsCritical() bool

	Move() move.Move

	MoveTypes() *types.Types
	MoveEffective() types.Effective
	MoveAttribute() attribute.Attribute
}

type situation struct {
	attacker status.StatusChecker
	defender status.StatusChecker
	move move.Move
	fields *field.Fields

	attackersItem item.Item
	defendersItem item.Item

	abilities ability.AbilityField

	isCritical bool
}

func (s *situation) Attacker() status.StatusChecker {
	ability := s.abilities.CorrectAttackerStatus(s)
	ac, _ := s.fields.StatusCorrector(s.attacker.Types(), s.defender.Types())
	ic := s.attackersItem.Correct()
	res := ability.Create(s.attacker)
	res = ac.Create(res)
	res = ic.Create(res)
	return res
}
func (s *situation) Defender() status.StatusChecker {
	ability := s.abilities.CorrectDefenderStatus(s)
	_, dc := s.fields.StatusCorrector(s.attacker.Types(), s.defender.Types())
	ic := s.defendersItem.Correct()
	res := ability.Create(s.defender)
	res = dc.Create(res)
	res = ic.Create(res)
	return res
}
func (s *situation) Move() move.Move {
	return s.move
}
func (s *situation) Correctors() []corrector.Corrector {
	res := make([]corrector.Corrector, 0)
	res = append(res, s.move.Correctors(s)...)
	res = append(res, s.fields.Correctors(s.move.Types(s))...)
	res = append(res, s.abilities.Correctors(s)...)
	atItem := s.attackersItem.CorrectPower(s.Attacker().Types(), s.Defender().Types(), s.move.Types(s))

	res = append(res, atItem)
	return res
}
func (s *situation) IsWeather(f field.Weather) bool {
	return s.fields.HasWeather(f)
}
func (s *situation) IsCritical() bool {
	return s.isCritical
}

func (s *situation) IsField(f field.Field) bool {
	return s.fields.HasField(f)
}

func (s *situation) MoveTypes() *types.Types {
	return s.move.Types(s)
}
func (s *situation) MoveEffective() types.Effective {
	df := s.Defender().Types()
	return s.MoveTypes().Magnification(df)
}
func (s *situation) MoveAttribute() attribute.Attribute {
	return s.move.Attribute()

}
