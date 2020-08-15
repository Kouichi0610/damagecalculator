/*
	ダメージ計算に必要な場の状態
*/
package situation

import (
	"damagecalculator/domain/ability"
	"damagecalculator/domain/corrector"
	"damagecalculator/domain/field"
	"damagecalculator/domain/item"
	"damagecalculator/domain/skill"
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

	// TODO:直接出さない方針で
	Skill() skill.Skill

	SkillTypes() *types.Types
	SkillEffective() types.Effective
	SkillAction() skill.Action
}

type situation struct {
	at status.StatusChecker
	df status.StatusChecker
	sk skill.Skill
	fl *field.Fields

	atItem item.Item
	dfItem item.Item

	abilities ability.AbilityField

	isCritical bool
}

func (s *situation) Attacker() status.StatusChecker {
	ability := s.abilities.Attacker().CorrectStatus(s)
	ac, _ := s.fl.StatusCorrector(s.at.Types(), s.df.Types())
	ic := s.atItem.Correct()
	res := ability.Create(s.at)
	res = ac.Create(res)
	res = ic.Create(res)
	return res
}
func (s *situation) Defender() status.StatusChecker {
	ability := s.abilities.Defender().CorrectStatus(s)
	_, dc := s.fl.StatusCorrector(s.at.Types(), s.df.Types())
	ic := s.dfItem.Correct()
	res := ability.Create(s.df)
	res = dc.Create(res)
	res = ic.Create(res)
	return res
}
func (s *situation) Skill() skill.Skill {
	return s.sk
}
func (s *situation) Correctors() []corrector.Corrector {
	res := make([]corrector.Corrector, 0)
	res = append(res, s.sk.Correctors(s)...)
	res = append(res, s.fl.Correctors(s.sk.Types(s))...)
	res = append(res, s.abilities.Correctors(s)...)
	atItem := s.atItem.CorrectPower(s.Attacker().Types(), s.Defender().Types(), s.sk.Types(s))

	res = append(res, atItem)
	return res
}
func (s *situation) IsWeather(f field.Weather) bool {
	return s.fl.HasWeather(f)
}
func (s *situation) IsCritical() bool {
	return s.isCritical
}

func (s *situation) IsField(f field.Field) bool {
	return s.fl.HasField(f)
}
func (s *situation) SkillTypes() *types.Types {
	return s.sk.Types(s)
}
func (s *situation) SkillEffective() types.Effective {
	df := s.Defender().Types()
	return s.SkillTypes().Magnification(df)
}
func (s *situation) SkillAction() skill.Action {
	return s.sk.Action()
}
