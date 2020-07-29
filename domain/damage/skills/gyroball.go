package skills

import (
	"damagecalculator/domain/damage"
)

type GyroBall struct {
	Skill
}

func (s *GyroBall) Calculate(a, b *Stats) Damage {
	return Damage(10)
}
