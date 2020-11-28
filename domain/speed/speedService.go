package speed

import (
	"damagecalculator/domain/ability"
	"damagecalculator/domain/item"
	"damagecalculator/domain/move"
	"damagecalculator/domain/situation"
	"damagecalculator/domain/species"
	"damagecalculator/domain/stats"
)

// 天候、フィールド、もちもの、とくせいから素早さを取得

type (
	ServiceArgs struct {
		Name        string
		Level       uint
		Individuals string
		BasePoint   uint
		Ability     string
		Nature      string
		Item        string
		Condition   string
		Weather     string
		Field       string
	}
	SpeedService interface {
		Speed(args *ServiceArgs) uint
	}
)

func NewService(sp species.Repository, ab ability.Repository, mv move.Repository, it item.Repository) SpeedService {
	return &speedService{
		sp: sp,
		ab: ab,
		mv: mv,
		it: it,
	}
}

func (s *speedService) Speed(args *ServiceArgs) uint {
	builder := situation.NewBuilder(s.sp, s.ab, s.mv, s.it)
	fields := &situation.FieldCondition{
		Weather:      args.Weather,
		Field:        args.Field,
		HasReflector: false,
	}
	attacker := &situation.PokeParams{
		Name:        args.Name,
		Individuals: args.Individuals,
		BasePoints:  []uint{0, 0, 0, 0, 0, args.BasePoint},
		Ranks:       []int{0, 0, 0, 0, 0},
		Ability:     args.Ability,
		Item:        args.Item,
		Nature:      args.Nature,
		Condition:   args.Condition,
	}
	defender := &situation.PokeParams{
		Name:        "シェルダー",
		Individuals: "Max",
		BasePoints:  []uint{0, 0, 0, 0, 0, 0},
		Ranks:       []int{0, 0, 0, 0, 0},
		Ability:     "げきりゅう",
		Item:        "None",
		Condition:   "なし",
	}
	move := "ボルテッカー"
	level := stats.NewLevel(args.Level)
	st, _ := builder.ToSituation(level, attacker, defender, move, fields)
	return st.Attacker().Speed().Value()
}

type speedService struct {
	sp species.Repository
	ab ability.Repository
	mv move.Repository
	it item.Repository
}
