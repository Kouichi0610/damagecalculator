package supposition

import (
	"damagecalculator/domain/ability"
	_ "damagecalculator/domain/condition"
	_ "damagecalculator/domain/damage"
	"damagecalculator/domain/item"
	"damagecalculator/domain/move"
	"damagecalculator/domain/pokenames"
	"damagecalculator/domain/situation"
	"damagecalculator/domain/species"
	_ "damagecalculator/domain/stats"
	"damagecalculator/domain/types"
)

/*
	仮想敵リストの作成、
	仮想敵から受けるダメージ一覧を取得

*/
type (
	Attacker interface {
		Param() situation.PokeParams
		Move() string
	}
	Attackers []Attacker

	AttackersMaker interface {
		// TODO:PokeParams & Move
		// TODO:相性のいい技を考慮する必要がある
		Attackers(defender, defendersAbility string) (Attackers, error)
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

func (d *amaker) Attackers(defender, defendersAbility string) (Attackers, error) {
	sv := newBestAbilityService(d.sp, d.mv, d.ab, d.it)
	// TODO:リストごとに最善の

	/*
		res := make(Attackers, 0)
		df, err := d.sp.Get(defender)
		if err != nil {
			return nil, err
		}
		ty := types.NewTypes(df.Types...)
		return res, nil
	*/
	return nil, nil

}

/*
	有効打を持つ(タイプ一致優先)
	攻撃力高い順で

	物理特殊分けるか
*/
func (d *amaker) effectiveAttackers(ty *types.Types) Attackers {
	res := make(Attackers, 0)

	return res
}

/*
	攻撃力の高いポケモンを絞り込んでいく
	それぞれ一番有効なわざ(+とくせい)を取得

*/

// 攻撃側から最も有効なわざ(&とくせい)を選択
func (d *amaker) bestMove(attacker, defender, defenderAbility string) (Attacker, error) {
	/*
		sv := damage.NewSimpleService(d.sp, d.mv, d.ab, d.it)
		at, err := d.sp.Get(attacker)
		if err != nil {
			return nil, err
		}

			ある程度決め打ちで絞り込んでおく
			もちもの、性格
			1.タイプ一致&&物理or特殊の強いほうを選ぶ
			2.総当たり

		for _, ability := range at.Abilities {
			m, _ := d.mv.Get("")
			//sv.Calculate(attacker, defender, ability, defenderAbility, "")

			// 代表(category, ability) Best(move)
			// 物理、特殊強いほうで
		}
		// 特性込みで実際に有効か
		//	damage.DamageService
	*/
	return nil, nil
}

type attacker struct {
	param situation.PokeParams
	move  string
}

func (a *attacker) Param() situation.PokeParams {
	return a.param
}
func (a *attacker) Move() string {
	return a.move
}

// 暫定候補一覧
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
