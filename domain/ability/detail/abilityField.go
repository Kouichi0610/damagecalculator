package detail

import (
	"damagecalculator/domain/ability/situation"
	"damagecalculator/domain/corrector"
	"damagecalculator/domain/move"
	"damagecalculator/domain/status"
)

type eraser interface {
	// 場の特性全て無効にする
	erase()
}

/*
	implements ability.AbilityField, eraser
*/
type abilityField struct {
	at, df ability
}

func NewAbilityField(at, df ability) interface{} {
	at.setAttacker(true)
	df.setAttacker(false)
	res := &abilityField{
		at: at,
		df: df,
	}
	res.at.onField(res)
	res.df.onField(res)
	res.at.onAttack(res)
	res.df.onAttack(res)
	return res
}

func (a *abilityField) CorrectAttackerStatus(st situation.SituationChecker) *status.StatsCorrectors {
	return a.at.CorrectStatus(st)
}
func (a *abilityField) CorrectDefenderStatus(st situation.SituationChecker) *status.StatsCorrectors {
	return a.df.CorrectStatus(st)
}
func (a *abilityField) RewriteMoveFactory(mv move.MoveFactory) *move.MoveFactory {
	res := a.at.RewriteMoveFactory(mv)
	res = a.df.RewriteMoveFactory(*res)
	return res
}
func (a *abilityField) Correctors(st situation.SituationChecker) []corrector.Corrector {
	res := a.at.Correctors(st)
	res = append(res, a.df.Correctors(st)...)
	return res
}
func (a *abilityField) erase() {
	a.at = &abilityImpl{}
	a.at.setAttacker(true)
	a.df = &abilityImpl{}
	a.at.setAttacker(false)
}
