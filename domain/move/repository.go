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
		CountMin  uint
		CountMax  uint
		Target    target.MoveTarget
		Attribute attribute.Attribute
		Detail    detail.Detail
	}
)

// Moveインターフェイスを生成する
func (m *MoveFactory) Create() (Move, error) {
	cnt, err := count.NewAttackCount(m.CountMin, m.CountMax)
	res, err := detail.NewMove(
		m.Power,
		m.Type,
		m.Accuracy,
		m.Category,
		cnt,
		m.Target,
		m.Attribute,
		m.Detail)

	if err != nil {
		return nil, err
	}
	return res.(Move), nil
}
