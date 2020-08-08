package skill

import (
	"damagecalculator/domain/skill/category"
	"damagecalculator/domain/skill/count"
	"damagecalculator/domain/types"
	"fmt"
)

type Method uint

const (
	None        Method = iota
	SeismicToss        // ちきゅうなげ
	WeatherBall        // ウェザーボール
)

// 攻撃回数
// 対象(1or複数)
type SkillData struct {
	types    []types.Type
	power    uint
	countMin uint
	countMax uint
	category category.Category
	method   Method
}

func NewSkill(d *SkillData) (Skill, error) {
	c, err := count.NewAttackCount(d.countMin, d.countMax)
	if err != nil {
		return nil, err
	}
	p, err := category.NewCategory(d.category)
	if err != nil {
		return nil, err
	}
	t := types.NewTypes(d.types...)
	switch d.method {
	case None:
		return &skill{
			types:  t,
			power:  d.power,
			count:  c,
			picker: p,
		}, nil
	case SeismicToss:
		return &seismicToss{
			skill: skill{
				types:  t,
				power:  d.power,
				count:  c,
				picker: p,
			},
		}, nil
	case WeatherBall:
		return &weatherBall{
			skill: skill{
				types:  t,
				power:  d.power,
				count:  c,
				picker: p,
			},
		}, nil
	}

	return nil, fmt.Errorf("%d not supported.", d.method)
}
