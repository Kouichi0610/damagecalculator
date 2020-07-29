package damage

import (
	"damagecalculator/domain/field"
	"damagecalculator/domain/types"
)

type SkillPower uint

// TODO:Damages をstructとして

type DamageCalculator interface {
	Calculate(a, b *Stats, f *fields.Fields) Damage
}

type Skill struct {
	types    *types.Types // 攻撃タイプ
	power    SkillPower   // 威力
	category CategoryFunc
	// 命中はひとまず考慮しない。
}

// 2～5回攻撃
// 必ず急所
// ジャイロボールなど(skillPowerに含めるか)
// HPの差額をそのままダメージに
