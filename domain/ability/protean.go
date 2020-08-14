package ability

import (
	"damagecalculator/domain/status"
)

// へんげんじざい
// 攻撃時、技のタイプと同じになる
type protean struct {
	ability
}

func (s *protean) CorrectStatus(st situationChecker) *status.StatsCorrectors {
	c := status.NewStatsCorrectors()
	if !s.isAttacker {
		return c
	}
	ty := st.SkillTypes()
	c.Types(ty.TypeArray())
	return c
}
