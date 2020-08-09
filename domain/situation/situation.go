/*
	ダメージ計算に必要な場の状態
*/
package situation

import (
	"damagecalculator/domain/field"
	"damagecalculator/domain/skill"
	"damagecalculator/domain/status"
)

// 状況　and 状況問い合わせ

type SituationChecker interface {
	Attacker() status.StatusChecker
	Defender() status.StatusChecker
	Skill() skill.Skill
	IsWeather(field.Weather) bool
}

type situation struct {
	at status.StatusChecker
	df status.StatusChecker
	sk skill.Skill
	fl *field.Fields
}

func (s *situation) Attacker() status.StatusChecker {
	return s.at
}
func (s *situation) Defender() status.StatusChecker {
	return s.df
}
func (s *situation) Skill() skill.Skill {
	return s.sk
}
func (s *situation) IsWeather(f field.Weather) bool {
	return s.fl.HasWeather(f)
}
