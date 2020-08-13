package item

import (
	"damagecalculator/domain/status"
)

// 能力値補正
type statsCorrector interface {
	Correct() *status.StatsCorrectors
}

func defaultStatsCorrector() statsCorrector {
	return &statsCorrectorImpl{
		c: status.NewStatsCorrectors(),
	}
}

func (s *StatsCorrectData) createStatsCorrector() statsCorrector {
	c := status.NewStatsCorrectors()
	if s.Attack > 0 {
		c.Attack(s.Attack)
	}
	if s.Defense > 0 {
		c.Defense(s.Defense)
	}
	if s.SpAttack > 0 {
		c.SpAttack(s.SpAttack)
	}
	if s.SpDefense > 0 {
		c.SpDefense(s.SpDefense)
	}
	if s.Speed > 0 {
		c.Speed(s.Speed)
	}
	return &statsCorrectorImpl{
		c: c,
	}
}

type statsCorrectorImpl struct {
	c *status.StatsCorrectors
}

func (s *statsCorrectorImpl) Correct() *status.StatsCorrectors {
	return s.c
}
