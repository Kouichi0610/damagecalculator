/*
	ダメージ計算に必要な場の状態
*/
package situation

import (
	"damagecalculator/domain/corrector"
	"damagecalculator/domain/field"
	"damagecalculator/domain/item"
	"damagecalculator/domain/skill"
	"damagecalculator/domain/status"
)

// 状況　and 状況問い合わせ
// TODO:能力値補正(すなあらしで岩タイプ、ようりょくそ->晴れ時素早さ2...)

type SituationChecker interface {
	Attacker() status.StatusChecker
	Defender() status.StatusChecker
	Skill() skill.Skill
	Correctors() []corrector.Corrector
	IsWeather(field.Weather) bool
	IsCritical() bool
}

type situation struct {
	at status.StatusChecker
	df status.StatusChecker
	sk skill.Skill
	fl *field.Fields

	atItem item.Item
	dfItem item.Item

	isCritical bool
}

func (s *situation) Attacker() status.StatusChecker {
	// TODO:dcが無駄になる
	ac, _ := s.fl.StatusCorrector(s.at.Types(), s.df.Types())
	ic := s.atItem.Correct()
	res := ac.Create(s.at)
	res = ic.Create(res)
	return res
}
func (s *situation) Defender() status.StatusChecker {
	_, dc := s.fl.StatusCorrector(s.at.Types(), s.df.Types())
	ic := s.dfItem.Correct()
	res := dc.Create(s.df)
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
