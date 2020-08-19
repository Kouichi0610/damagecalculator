package species

import (
	"damagecalculator/domain/gender"
	"damagecalculator/domain/stats"
	"damagecalculator/domain/status"
	"damagecalculator/domain/types"
)

/*
	種族に関するリポジトリ
	名前からStatus, わざ、とくせいリストを取得
*/
type (

	// statusに持たせて status, []abilities, []moves (Gender)
	Species struct {
		Name string

		Stats  *stats.SpeciesStats
		Types  []types.Type
		Weight float64
		Gender gender.Gender

		Abilities []string // 選択可能なとくせい一覧(id intでいいか)
		Moves     []string // 選択可能なわざ一覧
	}

	Repository interface {
		Get(name string) (*Species, error)
	}

	// statuschecker生成
	StatusFactory interface {
		Create(*StatusFactoryArgs) (st status.StatusChecker, abilities, moves []string, g gender.Gender, err error)
	}
)

func NewStatusFactory(r Repository) StatusFactory {
	return &statusFactory{
		r: r,
	}
}

type statusFactory struct {
	r Repository
}

type StatusFactoryArgs struct {
	Name       string
	Level      uint
	Ranks      [5]int
	Nature     stats.NatureType
	Individual *stats.IndividualStats
	BasePoint  *stats.BasePointStats
}

// name(id), Level, Ranks Nature, Individual, BasePoint
func (s *statusFactory) Create(args *StatusFactoryArgs) (st status.StatusChecker, abilities, moves []string, g gender.Gender, err error) {
	sp, err := s.r.Get(args.Name)
	if err != nil {
		return nil, nil, nil, gender.Unknown, err
	}

	lv := stats.NewLevel(args.Level)
	nt := stats.NewNature(args.Nature)
	stat := stats.NewStats(lv, sp.Stats, args.Individual, args.BasePoint, nt)

	a := &status.StatusData{
		Lv:            args.Level,
		Types:         sp.Types,
		HP:            stat.HP,
		Attack:        stat.Attack,
		Defense:       stat.Defense,
		SpAttack:      stat.SpAttack,
		SpDefense:     stat.SpDefense,
		Speed:         stat.Speed,
		AttackRank:    args.Ranks[0],
		DefenseRank:   args.Ranks[1],
		SpAttackRank:  args.Ranks[2],
		SpDefenseRank: args.Ranks[3],
		SpeedRank:     args.Ranks[4],
		Weight:        sp.Weight,
	}

	return a.Create(), sp.Abilities, sp.Moves, sp.Gender, nil

}
