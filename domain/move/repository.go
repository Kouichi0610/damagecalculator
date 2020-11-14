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
	"fmt"
)

type (
	Repository interface {
		Get(name string) (*MoveFactory, error)
	}

	MoveFactory struct {
		Name     string
		Power    power.Power
		Type     types.Type
		Accuracy accuracy.Accuracy
		Category category.DamageCategory
		CountMin uint
		CountMax uint
		Target   target.MoveTarget
		Action   attribute.Action
		Effect   attribute.Effect
		Detail   detail.Detail
	}
)

func (m *MoveFactory) Types() string {
	return m.Type.String()
}

// 特筆事項
func (m *MoveFactory) Mention() string {
	res := ""
	if m.CountMin > 1 {
		if m.CountMin == m.CountMax {
			res += fmt.Sprintf("%d回攻撃", m.CountMin)
		} else {
			res += fmt.Sprintf("%d～%d回攻撃", m.CountMin, m.CountMax)
		}
	}
	return res
}

// Moveインターフェイスを生成する
func (m *MoveFactory) Create() (Move, error) {
	cnt, err := count.NewAttackCount(m.CountMin, m.CountMax)
	attr := attribute.NewAttribute(m.Action, m.Effect)
	res, err := detail.NewMove(
		m.Power,
		m.Type,
		m.Accuracy,
		m.Category,
		cnt,
		m.Target,
		attr,
		m.Detail)

	if err != nil {
		return nil, err
	}
	return res.(Move), nil
}
