/*
	能力値を計算してstats.Statsを生成する
*/
package stats

import (
	"damagecalculator/domain/basepoints"
)

type stats struct {
	HP, Attack, Defense, SpAttack, SpDefense, Speed uint
}

func CalcHP(l Level, s *SpeciesStats, i *IndividualStats, b uint, n *Nature) uint {
	return n.calcHP(l, s.hp, i.hp, basePoints.NewBasePoint(b))
}
func CalcAttack(l Level, s *SpeciesStats, i *IndividualStats, b uint, n *Nature) uint {
	return n.calcAttack(l, s.at, i.at, basePoints.NewBasePoint(b))
}
func CalcDefense(l Level, s *SpeciesStats, i *IndividualStats, b uint, n *Nature) uint {
	return n.calcDefense(l, s.df, i.df, basePoints.NewBasePoint(b))
}
func CalcSpAttack(l Level, s *SpeciesStats, i *IndividualStats, b uint, n *Nature) uint {
	return n.calcSpAttack(l, s.sa, i.sa, basePoints.NewBasePoint(b))
}
func CalcSpDefense(l Level, s *SpeciesStats, i *IndividualStats, b uint, n *Nature) uint {
	return n.calcSpDefense(l, s.sd, i.sd, basePoints.NewBasePoint(b))
}
func CalcSpeed(l Level, s *SpeciesStats, i *IndividualStats, b uint, n *Nature) uint {
	return n.calcSpeed(l, s.sp, i.sp, basePoints.NewBasePoint(b))
}

func NewStats(l Level, s *SpeciesStats, i *IndividualStats, b basePoints.BasePoints, n *Nature) *stats {
	hp := n.calcHP(l, s.hp, i.hp, b.HP())
	at := n.calcAttack(l, s.at, i.at, b.Attack())
	df := n.calcDefense(l, s.df, i.df, b.Defense())
	sa := n.calcSpAttack(l, s.sa, i.sa, b.SpAttack())
	sd := n.calcSpDefense(l, s.sd, i.sd, b.SpDefense())
	sp := n.calcSpeed(l, s.sp, i.sp, b.Speed())

	return &stats{
		HP:        hp,
		Attack:    at,
		Defense:   df,
		SpAttack:  sa,
		SpDefense: sd,
		Speed:     sp,
	}
}
