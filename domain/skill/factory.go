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
	part     Part
}

func newSkill(d *SkillData) (skill, error) {
	c, err := count.NewAttackCount(d.countMin, d.countMax)
	if err != nil {
		return skill{}, err
	}
	p, err := category.NewCategory(d.category)
	if err != nil {
		return skill{}, err
	}
	t := types.NewTypes(d.types...)
	return skill{
		types:  t,
		power:  d.power,
		count:  c,
		picker: p,
		part:   d.part,
	}, nil
}

func NewSkill(d *SkillData) (Skill, error) {
	s, err := newSkill(d)
	if err != nil {
		return nil, err
	}
	switch d.method {
	case None:
		return &s, nil
	case SeismicToss:
		return &seismicToss{
			skill: s,
		}, nil
	case WeatherBall:
		return &weatherBall{
			skill: s,
		}, nil
	}
	return nil, fmt.Errorf("%d not supported.", d.method)
}
