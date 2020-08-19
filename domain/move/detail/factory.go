package detail

import (
	"damagecalculator/domain/move/accuracy"
	"damagecalculator/domain/move/attribute"
	"damagecalculator/domain/move/category"
	"damagecalculator/domain/move/count"
	"damagecalculator/domain/move/power"
	"damagecalculator/domain/move/target"
	"damagecalculator/domain/types"
	"fmt"
)

// move.Moveインターフェイスを生成する
func NewMove(
	power power.Power,
	ty types.Type,
	accuracy accuracy.Accuracy,
	cat category.DamageCategory,
	count *count.AttackCount,
	target target.MoveTarget,
	attribute attribute.Attribute,
	detail Detail) (interface{}, error) {

	picker, err := category.NewCategory(cat)
	if err != nil {
		return nil, err
	}
	mv := &defaultMove{
		ty:          ty,
		power:       power,
		accuracy:    accuracy,
		attackCount: count,
		picker:      picker,
		attribute:   attribute,
		target:      target,
	}
	switch detail {
	case Default:
		return mv, nil
	case SeismicToss:
		return &seismicToss{mv}, nil
		/*
			case WeatherBall:
				return nil, nil
			case GyroBall:
				return nil, nil
			case HeavySlam:
				return nil, nil
		*/
	}

	return nil, fmt.Errorf("%d not supported.", target)
}
