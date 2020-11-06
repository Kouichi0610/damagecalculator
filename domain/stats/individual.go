package stats

/*
	個体値
	HP、こうげき、ぼうぎょ、とくこう、とくぼう、すばやさそれぞれに設定
	個別の範囲は0～31
*/
var individualMax uint = 31

const (
	IndividualTypeMax IndividualType = iota
	IndividualTypeSlowest
	IndividualTypeWeakest
)

type (
	IndividualType  uint
	Individual      uint
	IndividualStats struct {
		hp, at, df, sa, sd, sp Individual
	}
)

func ToIndividualType(str string) IndividualType {
	if str == "Slowest" {
		return IndividualTypeSlowest
	}
	if str == "Weakest" {
		return IndividualTypeWeakest
	}
	return IndividualTypeMax
}

func (t IndividualType) Create() *IndividualStats {
	if t == IndividualTypeSlowest {
		return newIndividualStats(31, 31, 31, 31, 31, 0)
	}
	if t == IndividualTypeWeakest {
		return newIndividualStats(31, 0, 31, 31, 31, 31)
	}
	return newIndividualStats(31, 31, 31, 31, 31, 31)
}

func newIndividualStats(hp, at, df, sa, sd, sp uint) *IndividualStats {
	res := new(IndividualStats)
	res.hp = newIndividual(hp)
	res.at = newIndividual(at)
	res.df = newIndividual(df)
	res.sa = newIndividual(sa)
	res.sd = newIndividual(sd)
	res.sp = newIndividual(sp)
	return res
}

func (i *IndividualStats) HP() Individual {
	return i.hp
}
func (i *IndividualStats) Attack() Individual {
	return i.at
}
func (i *IndividualStats) Defense() Individual {
	return i.df
}
func (i *IndividualStats) SpAttack() Individual {
	return i.sa
}
func (i *IndividualStats) SpDefense() Individual {
	return i.sd
}
func (i *IndividualStats) Speed() Individual {
	return i.sp
}

func newIndividual(value uint) Individual {
	if value > individualMax {
		value = individualMax
	}
	res := Individual(value)
	return res
}
