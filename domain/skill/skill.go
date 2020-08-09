package skill

import (
	"damagecalculator/domain/corrector"
	"damagecalculator/domain/skill/category"
	"damagecalculator/domain/skill/count"
	"damagecalculator/domain/status"
	"damagecalculator/domain/types"
)

// わざからダメージ計算に必要なものを取得する
type Skill interface {
	Correctors(SituationChecker) []corrector.Corrector
	Types(SituationChecker) *types.Types
	Power(SituationChecker) uint
	PickStats(SituationChecker) (at, df *status.RankedValue)
	Calculate(level, power, attack, defense uint) []uint
}

// Skill without method
type skill struct {
	types  *types.Types
	power  uint
	count  *count.AttackCount
	picker category.CategoryFunc
	part   Part
}

func (s *skill) Types(SituationChecker) *types.Types {
	return s.types
}
func (s *skill) Power(SituationChecker) uint {
	return s.power
}
func (s *skill) PickStats(st SituationChecker) (at, df *status.RankedValue) {
	return s.picker(st.Attacker(), st.Defender())

}
func (s *skill) Calculate(level, power, attack, defense uint) []uint {
	res := calculateArray(level, power, attack, defense)
	return s.count.Correct(res)
}
func (s *skill) Correctors(SituationChecker) []corrector.Corrector {
	return nil
}
