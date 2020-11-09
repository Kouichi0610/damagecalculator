/*
	能力値を計算してstats.Statsを生成する
*/
package stats

import (
	"damagecalculator/domain/basepoints"
	"damagecalculator/domain/individuals"
)

type stats struct {
	HP, Attack, Defense, SpAttack, SpDefense, Speed uint
}

func CalcHP(l Level, s *SpeciesStats, i individuals.Individuals, b uint, n *Nature) uint {
	return n.calcHP(l, s.hp, i.HP(), basePoints.NewBasePoint(b))
}
func CalcAttack(l Level, s *SpeciesStats, i individuals.Individuals, b uint, n *Nature) uint {
	return n.calcAttack(l, s.at, i.Attack(), basePoints.NewBasePoint(b))
}
func CalcDefense(l Level, s *SpeciesStats, i individuals.Individuals, b uint, n *Nature) uint {
	return n.calcDefense(l, s.df, i.Defense(), basePoints.NewBasePoint(b))
}
func CalcSpAttack(l Level, s *SpeciesStats, i individuals.Individuals, b uint, n *Nature) uint {
	return n.calcSpAttack(l, s.sa, i.SpAttack(), basePoints.NewBasePoint(b))
}
func CalcSpDefense(l Level, s *SpeciesStats, i individuals.Individuals, b uint, n *Nature) uint {
	return n.calcSpDefense(l, s.sd, i.SpDefense(), basePoints.NewBasePoint(b))
}
func CalcSpeed(l Level, s *SpeciesStats, i individuals.Individuals, b uint, n *Nature) uint {
	return n.calcSpeed(l, s.sp, i.Speed(), basePoints.NewBasePoint(b))
}

func NewStats(l Level, s *SpeciesStats, i individuals.Individuals, b basePoints.BasePoints, n *Nature) *stats {
	hp := n.calcHP(l, s.hp, i.HP(), b.HP())
	at := n.calcAttack(l, s.at, i.Attack(), b.Attack())
	df := n.calcDefense(l, s.df, i.Defense(), b.Defense())
	sa := n.calcSpAttack(l, s.sa, i.SpAttack(), b.SpAttack())
	sd := n.calcSpDefense(l, s.sd, i.SpDefense(), b.SpDefense())
	sp := n.calcSpeed(l, s.sp, i.Speed(), b.Speed())

	return &stats{
		HP:        hp,
		Attack:    at,
		Defense:   df,
		SpAttack:  sa,
		SpDefense: sd,
		Speed:     sp,
	}
}
