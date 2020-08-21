package detail

import (
	"damagecalculator/domain/ability/situation"
	"damagecalculator/domain/status"
)

// へんげんじざい
// 攻撃時、技のタイプと同じになる
type protean struct {
	abilityImpl
}

func (s *protean) CorrectStatus(st situation.SituationChecker) *status.StatsCorrectors {
	c := status.NewStatsCorrectors()
	if !s.isAttacker {
		return c
	}
	ty := st.MoveTypes()
	c.Types(ty)
	return c
}
