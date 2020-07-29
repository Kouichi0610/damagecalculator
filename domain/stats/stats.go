/*
	能力値の計算
*/
package stats

type Stats struct {
	hp, at, df, sa, sd, sp uint
}

func NewStats(l Level, s *SpeciesStats, i *IndividualStats, b *BasePointStats, n *Nature) *Stats {
	res := new(Stats)
	res.hp = n.calcHP(l, s.hp, i.hp, b.hp)
	res.at = n.calcAttack(l, s.at, i.at, b.at)
	res.df = n.calcDefense(l, s.df, i.df, b.df)
	res.sa = n.calcSpAttack(l, s.sa, i.sa, b.sa)
	res.sd = n.calcSpDefense(l, s.sd, i.sd, b.sd)
	res.sp = n.calcSpeed(l, s.sp, i.sp, b.sp)
	return res
}

func (s *Stats) HP() uint {
	return s.hp
}
func (s *Stats) Attack() uint {
	return s.at
}
func (s *Stats) Defense() uint {
	return s.df
}
func (s *Stats) SpAttack() uint {
	return s.sa
}
func (s *Stats) SpDefense() uint {
	return s.sd
}
func (s *Stats) Speed() uint {
	return s.sp
}
