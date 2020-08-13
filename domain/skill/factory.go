package skill

import (
	"damagecalculator/domain/skill/category"
	"damagecalculator/domain/skill/count"
	"damagecalculator/domain/types"
	"fmt"
)

type Method uint

const (
	NoMethod    Method = iota
	SeismicToss        // ちきゅうなげ
	WeatherBall        // ウェザーボール
	GyroBall           // ジャイロボール
	HeavySlam          // ヘビーボンバー
)

// 攻撃回数
// 対象(1or複数)
type SkillData struct {
	Types     []types.Type
	Power     uint
	CountMin  uint
	CountMax  uint
	Category  category.Category
	Method    Method
	Action    Action
	Attribute Attribute
}

func (d *SkillData) Create() (Skill, error) {
	s, err := newSkill(d)
	if err != nil {
		return nil, err
	}
	switch d.Method {
	case NoMethod:
		return &s, nil
	case SeismicToss:
		return &seismicToss{
			skill: s,
		}, nil
	case WeatherBall:
		return &weatherBall{
			skill: s,
		}, nil
	case GyroBall:
		return &gyroBall{
			skill: s,
		}, nil
	case HeavySlam:
		return &heavySlam{
			skill: s,
		}, nil
	}
	return nil, fmt.Errorf("%d not supported.", d.Method)
}

// create default skill.
func newSkill(d *SkillData) (skill, error) {
	c, err := count.NewAttackCount(d.CountMin, d.CountMax)
	if err != nil {
		return skill{}, err
	}
	p, err := category.NewCategory(d.Category)
	if err != nil {
		return skill{}, err
	}
	t := types.NewTypes(d.Types...)
	return skill{
		types:  t,
		power:  d.Power,
		count:  c,
		picker: p,
		action: d.Action,
		attr:   d.Attribute,
	}, nil
}
