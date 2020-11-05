package statspattern

import (
	"damagecalculator/domain/species"
	"damagecalculator/domain/stats"
)

/*
	名前と性格、個体値から
	基礎ポイント0～252(4刻み)ごとの能力値一覧を取得する
*/
type (
	Loader interface {
		Get(level int, name string, nature string, hp int, at int, df int, sa int, sd int, sp int) (StatsPattern, error)
	}
	StatsPattern interface {
		HP() []uint
		Attack() []uint
		Defense() []uint
		SpAttack() []uint
		SpDefense() []uint
		Speed() []uint
	}
)

func NewLoader(rp species.Repository) Loader {
	return &loader{
		rp: rp,
	}
}

type loader struct {
	rp species.Repository
}
type statsPattern struct {
	hp []uint
	at []uint
	df []uint
	sa []uint
	sd []uint
	sp []uint
}

func (s *statsPattern) HP() []uint {
	return s.hp
}
func (s *statsPattern) Attack() []uint {
	return s.at
}
func (s *statsPattern) Defense() []uint {
	return s.df
}
func (s *statsPattern) SpAttack() []uint {
	return s.sa
}
func (s *statsPattern) SpDefense() []uint {
	return s.sd
}
func (s *statsPattern) Speed() []uint {
	return s.sp
}

func (l *loader) Get(level int, name string, nature string, hp int, at int, df int, sa int, sd int, sp int) (StatsPattern, error) {
	s, err := l.rp.Get(name)
	if err != nil {
		return nil, err
	}
	nt := stats.NameToNature(nature)
	lv := stats.NewLevel(uint(level))

	sstats := stats.NewSpeciesStats(s.HP, s.Attack, s.Defense, s.SpAttack, s.SpDefense, s.Speed)
	istats := stats.NewIndividualStats(uint(hp), uint(at), uint(df), uint(sa), uint(sd), uint(sp))

	hpval := make([]uint, 0)
	atval := make([]uint, 0)
	dfval := make([]uint, 0)
	saval := make([]uint, 0)
	sdval := make([]uint, 0)
	spval := make([]uint, 0)

	for i := 0; i <= 252; i += 4 {
		if name == "ヌケニン" {
			hpval = append(hpval, 1)
		} else {
			hpval = append(hpval, stats.CalcHP(lv, sstats, istats, uint(i), nt))
		}
		atval = append(atval, stats.CalcAttack(lv, sstats, istats, uint(i), nt))
		dfval = append(dfval, stats.CalcDefense(lv, sstats, istats, uint(i), nt))
		saval = append(saval, stats.CalcSpAttack(lv, sstats, istats, uint(i), nt))
		sdval = append(sdval, stats.CalcSpDefense(lv, sstats, istats, uint(i), nt))
		spval = append(spval, stats.CalcSpeed(lv, sstats, istats, uint(i), nt))
	}

	return &statsPattern{
		hp: hpval,
		at: atval,
		df: dfval,
		sa: saval,
		sd: sdval,
		sp: spval,
	}, nil
}
