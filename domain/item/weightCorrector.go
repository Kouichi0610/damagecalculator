package item

import (
	"damagecalculator/domain/status"
)

// 重さ補正
func (s *WeightCorrectData) createStatsCorrector() statsCorrector {
	c := status.NewStatsCorrectors()
	c.Weight(s.Scale)
	return &statsCorrectorImpl{
		c: c,
	}
}
