package ability

import (
	"damagecalculator/domain/field"
	"damagecalculator/domain/skill"
	_ "damagecalculator/domain/status"
	"damagecalculator/domain/types"
)

type situationChecker interface {
	// Skill() skill.Skill // 状況の影響を受けるので直接アクセスさせない
	SkillTypes() *types.Types
	SkillEffective() types.Effective
	IsWeather(field.Weather) bool
	IsField(field.Field) bool
	SkillAction() skill.Action
}
