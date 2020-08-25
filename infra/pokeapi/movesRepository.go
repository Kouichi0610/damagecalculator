package pokeapi

import (
	"damagecalculator/domain/move"
	"damagecalculator/domain/move/accuracy"
	"damagecalculator/domain/move/attribute"
	"damagecalculator/domain/move/category"
	"damagecalculator/domain/move/detail"
	"damagecalculator/domain/move/power"
	"damagecalculator/domain/move/target"

	"github.com/mtslzr/pokeapi-go"
)

type movesRepository struct {
}

func (r *movesRepository) Get(name string) (*move.MoveFactory, error) {
	mv, err := pokeapi.Move(name)
	if err != nil {
		return nil, err
	}
	jname := mv.Name
	for _, n := range mv.Names {
		if n.Language.Name == JpName {
			jname = n.Name
			break
		}
	}

	target := moveTarget(mv.Target.Name)
	min, max := hitCounts(mv.Meta.MinHits, mv.Meta.MaxHits)

	// 追加効果はmove.Meta.AilmentChanceが0でなければとれる

	// TODO:接触、非接触
	// TODO:キバ、パンチ、波動など取れないなら特性に対応わざのリストを載せるか

	// TODO:2～5回ダメージの確率が平等でない
	// Has a 3/8 chance each to hit 2 or 3 times, and a 1/8 chance each to hit 4 or 5 times.  Averages to 3 hits per use
	// スキルリンクでもなければ使わないとしても。

	// TODO:DamageCategory & Attribute
	damage := damageClass(mv.DamageClass.Name)
	action := attribute.Remote
	effect := attribute.NoAttribute

	res := &move.MoveFactory{
		Name:     jname,
		Power:    power.NewPower(uint(mv.Power)),
		Type:     typesMap[mv.Type.Name],
		Accuracy: accuracy.NewAccuracy(uint(mv.Accuracy)),
		CountMin: uint(min),
		CountMax: uint(max),
		Target:   target,
		Category: damageCategory(jname, damage),
		Action:   action,
		Effect:   effect,
		Detail:   getDetail(jname),
	}
	return res, nil
}

func damageCategory(name string, ct category.DamageCategory) category.DamageCategory {
	switch name {
	case "サイコショック":
		return category.PsycoShock
	case "サイコブレイク":
		return category.PsycoShock
	case "ボディプレス":
		return category.BodyPress
	case "イカサマ":
		return category.FoulPlay
	case "シェルアームズ":
		return category.ShellArms
	}
	return ct
}

func getDetail(name string) detail.Detail {
	switch name {
	case "ジャイロボール":
		return detail.GyroBall
	case "ヘビーボンバー":
		return detail.HeavySlam
	case "ちきゅうなげ":
		return detail.SeismicToss
	case "ウェザーボール":
		return detail.WeatherBall
	case "ヒートスタンプ":
		return detail.HeavySlam
	}
	return detail.Default
}

func damageClass(name string) category.DamageCategory {
	switch name {
	case "physical":
		return category.Physical
	case "special":
		return category.Special
		/*
				現時点では対応しない
			case "status":
				return category.Physical
		*/
	}
	return category.Physical
}
func hitCounts(minHits, maxHits interface{}) (min, max int) {
	if minHits == nil || maxHits == nil {
		return 1, 1
	}
	if _, ok := minHits.(float64); !ok {
		return 1, 1
	}
	if _, ok := maxHits.(float64); !ok {
		return 1, 1
	}
	min = int(minHits.(float64))
	max = int(maxHits.(float64))
	return
}

// selected-pokemon 単体
// all-other-pokemon 他全て
// all-opponents　相手全て
// user 自身
func moveTarget(name string) target.MoveTarget {
	switch name {
	case "user":
		return target.User
	case "selected-pokemon":
		return target.Select
	case "all-other-pokemon":
		return target.AllOther
	case "all-opponents":
		return target.AllOpponents
	}
	return target.Select
}
