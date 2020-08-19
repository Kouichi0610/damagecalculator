package pokeapi

import (
	"damagecalculator/domain/skill"

	"github.com/mtslzr/pokeapi-go"
)

type movesRepository struct {
}

func (r *movesRepository) Get(name string) (*skill.Move, error) {
	move, err := pokeapi.Move(name)
	if err != nil {
		return nil, err
	}
	jname := move.Name
	for _, n := range move.Names {
		if n.Language.Name == JpName {
			jname = n.Name
			break
		}
	}

	damage := damageClass(move.DamageClass.Name)
	target := moveTarget(move.Target.Name)
	min, max := hitCounts(move.Meta.MinHits, move.Meta.MaxHits)

	// 追加効果はmove.Meta.AilmentChanceが0でなければとれる

	// TODO:接触、非接触
	// TODO:キバ、パンチ、波動など取れないなら特性に対応わざのリストを載せるか

	// TODO:2～5回ダメージの確率が平等でない
	// Has a 3/8 chance each to hit 2 or 3 times, and a 1/8 chance each to hit 4 or 5 times.  Averages to 3 hits per use
	// スキルリンクでもなければ使わないとしても。

	res := &skill.Move{
		Name:     jname,
		Damage:   damage,
		Power:    move.Power,
		Type:     typesMap[move.Type.Name],
		Accuracy: move.Accuracy,
		MinHits:  min,
		MaxHits:  max,
		Target:   target,
	}

	return res, nil
}

func damageClass(name string) skill.DamageClass {
	switch name {
	case "physical":
		return skill.Physical
	case "special":
		return skill.Special
	case "status":
		return skill.Status
	}
	return skill.Physical
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
func moveTarget(name string) skill.MoveTarget {
	switch name {
	case "user":
		return skill.User
	case "selected-pokemon":
		return skill.Select
	case "all-other-pokemon":
		return skill.AllOther
	case "all-opponents":
		return skill.AllOpponents
	}
	return skill.Select
}
