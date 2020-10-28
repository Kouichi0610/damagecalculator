/*
	能力値を計算してstats.Statsを生成する
*/
package stats

type stats struct {
	HP, Attack, Defense, SpAttack, SpDefense, Speed uint
}

func CalcHP(l Level, s *SpeciesStats, i *IndividualStats, b uint, n *Nature) uint {
	return n.calcHP(l, s.hp, i.hp, newBasePoint(b))
}
func CalcAttack(l Level, s *SpeciesStats, i *IndividualStats, b uint, n *Nature) uint {
	return n.calcAttack(l, s.at, i.at, newBasePoint(b))
}
func CalcDefense(l Level, s *SpeciesStats, i *IndividualStats, b uint, n *Nature) uint {
	return n.calcDefense(l, s.df, i.df, newBasePoint(b))
}
func CalcSpAttack(l Level, s *SpeciesStats, i *IndividualStats, b uint, n *Nature) uint {
	return n.calcSpAttack(l, s.sa, i.sa, newBasePoint(b))
}
func CalcSpDefense(l Level, s *SpeciesStats, i *IndividualStats, b uint, n *Nature) uint {
	return n.calcSpDefense(l, s.sd, i.sd, newBasePoint(b))
}
func CalcSpeed(l Level, s *SpeciesStats, i *IndividualStats, b uint, n *Nature) uint {
	return n.calcSpeed(l, s.sp, i.sp, newBasePoint(b))
}

func NewStats(l Level, s *SpeciesStats, i *IndividualStats, b *BasePointStats, n *Nature) *stats {
	hp := n.calcHP(l, s.hp, i.hp, b.hp)
	at := n.calcAttack(l, s.at, i.at, b.at)
	df := n.calcDefense(l, s.df, i.df, b.df)
	sa := n.calcSpAttack(l, s.sa, i.sa, b.sa)
	sd := n.calcSpDefense(l, s.sd, i.sd, b.sd)
	sp := n.calcSpeed(l, s.sp, i.sp, b.sp)

	return &stats{
		HP:        hp,
		Attack:    at,
		Defense:   df,
		SpAttack:  sa,
		SpDefense: sd,
		Speed:     sp,
	}
}
