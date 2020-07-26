package stats

/*
	能力値計算

*/
type Stats struct {
	hp, attack, defense, spAttack, spDefense, speed uint
}

func (s *Stats) HP() uint {
	return s.hp
}
func (s *Stats) Attack() uint {
	return s.attack
}
func (s *Stats) Defense() uint {
	return s.defense
}
func (s *Stats) SpAttack() uint {
	return s.spAttack
}
func (s *Stats) SpDefense() uint {
	return s.spDefense
}
func (s *Stats) Speed() uint {
	return s.speed
}
