package damage

import (
	"damagecalculator/domain/types"
)

// 技に債務が集まりすぎているので少しずつ分割していく

// TODO:ジャイロボール対応
type SkillPower uint
type CorrectFunc func(d Damage) Damage

type Skill struct {
	types    *types.Types // 攻撃タイプ
	power    SkillPower   // 威力
	category CategoryFunc
	// 命中はひとまず考慮しない。

	// calculatorをインターフェイス、あるいはfuncとして持つ
	// CalcDamage(a, d *Stats) []Damage
	// これなら大体対応できるはず
}

// 2～5回攻撃
// 必ず急所
// ジャイロボールなど(skillPowerに含めるか)
// HPの差額をそのままダメージに
