package move

import (
	"damagecalculator/domain/move/accuracy"
	"damagecalculator/domain/move/attribute"
	"damagecalculator/domain/move/category"
	"damagecalculator/domain/move/count"
	"damagecalculator/domain/move/detail"
	"damagecalculator/domain/move/power"
	"damagecalculator/domain/move/target"
	"damagecalculator/domain/types"
)

// TODO:skillをこっちに移動

type (
	Repository interface {
		Get(name string) (*MoveFactory, error)
	}

	MoveFactory struct {
		Name      string
		Power     power.Power
		Type      types.Type
		Accuracy  accuracy.Accuracy
		Category  category.DamageCategory
		Count     *count.AttackCount
		Target    target.MoveTarget
		Attribute attribute.Attribute
		Detail    detail.Detail
	}
)

// Moveインターフェイスを生成する
func (m *MoveFactory) Create() (Move, error) {
	res, err := detail.NewMove(
		m.Power,
		m.Type,
		m.Accuracy,
		m.Category,
		m.Count,
		m.Target,
		m.Attribute,
		m.Detail)

	if err != nil {
		return nil, err
	}
	return res.(Move), nil
}
