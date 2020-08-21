package situation

import (
	"damagecalculator/domain/field"
	"damagecalculator/domain/move/attribute"
	"damagecalculator/domain/types"
)

// 特性が状況判断するために必要なメソッド群
type SituationChecker interface {
	MoveTypes() *types.Types
	MoveEffective() types.Effective
	IsWeather(field.Weather) bool
	IsField(field.Field) bool
	MoveAttribute() attribute.Attribute
}
