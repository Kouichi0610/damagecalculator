package supposition

import (
	"damagecalculator/domain/ability"
	_ "damagecalculator/domain/condition"
	"damagecalculator/domain/damage"
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
	res := make(Attackers, 0)
	df, err := d.sp.Get(defender)
	if err != nil {
		return nil, err
	}
	ty := types.NewTypes(df.Types...)

	return res, nil
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
 */
func (d *amaker) calcDamage(attacker, defender, move string) uint {
	return 0
}

// 攻撃側から最も有効なわざを選択
func (d *amaker) bestMove(attacker, defender string) string {
	// TODO:簡易situation作成(situationに持たせる)
	// 特性込みで実際に有効か
	//	damage.DamageService

	return ""
}

// ガブリアス -> not fairy -> げきりん
// ガブリアス -> fairy -> じしん

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
