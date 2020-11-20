package supposition

import (
	"damagecalculator/domain/ability"
	"damagecalculator/domain/item"
	"damagecalculator/domain/move"
	"damagecalculator/domain/pokenames"
	"damagecalculator/domain/situation"
	"damagecalculator/domain/species"
)

/*
	仮想敵リストの作成、
	仮想敵から受けるダメージ一覧を取得

*/
type (
	Defender        string
	DefenderAbility string
	Attacker        interface {
		Param() *situation.PokeParams
		Move() string
	}
	Attackers []Attacker

	AttackersMaker interface {
		Attackers(defender Defender, defendersAbility DefenderAbility) Attackers
	}
)

func NewAttackersMaker(nm pokenames.Repository, sp species.Repository, mv move.Repository, ab ability.Repository, it item.Repository) AttackersMaker {
	return &amaker{
		nm: nm,
		sp: sp,
		mv: mv,
		ab: ab,
		it: it,
	}
}

type amaker struct {
	nm pokenames.Repository
	sp species.Repository
	mv move.Repository
	ab ability.Repository
	it item.Repository
}

func (d *amaker) Attackers(defender Defender, defendersAbility DefenderAbility) Attackers {
	sv := newBestAbilityService(d.sp, d.mv, d.ab, d.it)
	res := make(Attackers, 0)

	for _, at := range attackers {
		best := sv.BestMove(at, string(defender), string(defendersAbility))
		if !best.HasEffective() {
			continue
		}
		a := &attacker{
			move:  best.Move(),
			param: best.PokeParams(at),
		}
		res = append(res, a)
	}
	return res
}

type attacker struct {
	param *situation.PokeParams
	move  string
}

func (a *attacker) Param() *situation.PokeParams {
	return a.param
}
func (a *attacker) Move() string {
	return a.move
}

// 暫定候補一覧
// (TODO:性格、基礎ポイント含め決めておくべきか)
var attackers []string

func init() {
	attackers = []string{
		"リザードン",
		"フーディン",
		"カイリキー",
		"ゲンガー",
		"ギャラドス",
		"カイリュー",
		"ヘラクロス",
		"バンギラス",
		"ガブリアス",
		"サザンドラ",
		"ヒヒダルマ",
		"シャンデラ",
		"フェローチェ",
		"カミツルギ",
		"ニンフィア",
	}

}
