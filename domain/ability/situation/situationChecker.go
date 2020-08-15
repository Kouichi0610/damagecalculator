package situation

import (
	"damagecalculator/domain/field"
	"damagecalculator/domain/skill"
	"damagecalculator/domain/types"
)

// 特性が状況判断するために必要なメソッド群
type SituationChecker interface {
	SkillTypes() *types.Types
	SkillEffective() types.Effective
	IsWeather(field.Weather) bool
	IsField(field.Field) bool
	SkillAction() skill.Action
}
