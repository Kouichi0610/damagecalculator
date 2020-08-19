package detail

import (
	"damagecalculator/domain/corrector"
	"damagecalculator/domain/move/accuracy"
	"damagecalculator/domain/move/attribute"
	"damagecalculator/domain/move/category"
	"damagecalculator/domain/move/count"
	"damagecalculator/domain/move/power"
	"damagecalculator/domain/move/target"
	"damagecalculator/domain/status"
	"damagecalculator/domain/types"
)

// 特殊な効果のないわざの基底クラス
type defaultMove struct {
	ty        types.Type
	power     power.Power
	accuracy  accuracy.Accuracy
	count     *count.AttackCount
	picker    category.CategoryFunc
	attribute attribute.Attribute
	target    target.MoveTarget
}

// TODO:[]Correctors -> ?
func (m *defaultMove) Correctors(SituationChecker) []corrector.Corrector {
	return nil
}
func (m *defaultMove) Types(SituationChecker) *types.Types {
	return types.NewTypes(m.ty)
}
func (m *defaultMove) Power(SituationChecker) uint {
	return uint(m.power)
}
func (m *defaultMove) PickStats(st SituationChecker) (at, df *status.RankedValue) {
	attacker := st.Attacker()
	defender := st.Defender()
	return m.picker(attacker, defender)
}
func (m *defaultMove) Calculate(level, power, attack, defense uint) []uint {
	res := calculateArray(level, power, attack, defense)
	return m.count.Correct(res)
}

func (m *defaultMove) Attribute() attribute.Attribute {
	return attribute.NewAttribute(attribute.Fang, attribute.NoAttribute)
}
