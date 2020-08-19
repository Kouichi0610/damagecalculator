/*
	能力値を計算してstats.Statsを生成する
*/
package stats

type stats struct {
	HP, Attack, Defense, SpAttack, SpDefense, Speed uint
}

func NewStats(l Level, s *SpeciesStats, i *IndividualStats, b *BasePointStats, n *Nature) *stats {
	hp := n.calcHP(l, s.hp, i.hp, b.hp)
	at := n.calcAttack(l, s.at, i.at, b.at)
	df := n.calcDefense(l, s.df, i.df, b.df)
	sa := n.calcSpAttack(l, s.sa, i.sa, b.sa)
	sd := n.calcSpDefense(l, s.sd, i.sd, b.sd)
	sp := n.calcSpeed(l, s.sp, i.sp, b.sp)

	return &stats{
		HP: hp,
		Attack: at,
		Defense: df,
		SpAttack: sa,
		SpDefense: sd,
		Speed: sp,
	}
}
