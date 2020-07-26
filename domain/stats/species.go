package stats

/*
	種族値
	・HP,こうげき,ぼうぎょ,とくこう,とくぼう,すばやさ それぞれに設定
	・個別の範囲はひとまず1～255としておく
*/
var SpeciesMax uint = 255
var SpeciesMin uint = 1

type Species uint

type SpeciesStats struct {
	hp, at, df, sa, sd, sp Species
}

func NewSpeciesStats(hp, at, df, sa, sd, sp uint) *SpeciesStats {
	res := new(SpeciesStats)
	res.hp = newSpecies(hp)
	res.at = newSpecies(at)
	res.df = newSpecies(df)
	res.sa = newSpecies(sa)
	res.sd = newSpecies(sd)
	res.sp = newSpecies(sp)
	return res
}
func (s *SpeciesStats) HP() Species {
	return s.hp
}
func (s *SpeciesStats) Attack() Species {
	return s.at
}
func (s *SpeciesStats) Defense() Species {
	return s.df
}
func (s *SpeciesStats) SpAttack() Species {
	return s.sa
}
func (s *SpeciesStats) SpDefense() Species {
	return s.sd
}
func (s *SpeciesStats) Speed() Species {
	return s.sp
}

func newSpecies(value uint) Species {
	if value < SpeciesMin {
		value = SpeciesMin
	}
	if value > SpeciesMax {
		value = SpeciesMax
	}
	res := Species(value)
	return res
}
