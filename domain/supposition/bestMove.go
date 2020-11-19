package supposition

import (
	"damagecalculator/domain/ability"
	"damagecalculator/domain/damage"
	"damagecalculator/domain/item"
	"damagecalculator/domain/move"
	"damagecalculator/domain/move/category"
	"damagecalculator/domain/species"
)

type (
	BestAbilityService interface {
		BestMove(attacker, defender, defenderAbility string) BestMoveResult
	}
	// 最大威力となる技をと最低ダメージを取得
	BestCategoryService interface {
		BestMove(attacker, attackerAbility, defender, defenderAbility string) BestMoveResult
	}
	// 同じカテゴリから最大威力となるわざと最低ダメージを取得
	BestMoveService interface {
		BestMove(attacker, attackerAbility, defender, defenderAbility string, category category.DamageCategory) BestMoveResult
	}
	BestMoveResult interface {
		Ability() string
		Move() string
		Categoty() category.DamageCategory
		Damage() uint
		HasEffective() bool
	}
)

type baService struct {
	sp species.Repository
	mv move.Repository
	ab ability.Repository
	it item.Repository
}

type bcService struct {
	sp species.Repository
	mv move.Repository
	ab ability.Repository
	it item.Repository
}

func (b *baService) BestMove(attacker, defender, defenderAbility string) BestMoveResult {
	sv := newBestCategoryService(b.sp, b.mv, b.ab, b.it)
	sp, _ := b.sp.Get(attacker)

	damages := make(map[string]BestMoveResult)
	for _, ability := range sp.Abilities {
		damages[ability] = sv.BestMove(attacker, ability, defender, defenderAbility)
	}

	bestAbility := ""
	max := NoEffective()

	for ability, dmg := range damages {
		if max.Damage() < dmg.Damage() {
			max = dmg
			bestAbility = ability
		}
	}

	return &bmResult{
		move:     max.Move(),
		ability:  bestAbility,
		damage:   max.Damage(),
		category: max.Categoty(),
	}
}

func (b *bcService) BestMove(attacker, attackerAbility, defender, defenderAbility string) BestMoveResult {
	sv := newBestMoveService(b.sp, b.mv, b.ab, b.it)
	categories := category.List()
	damages := make(map[category.DamageCategory]BestMoveResult)
	for _, c := range categories {
		best := sv.BestMove(attacker, attackerAbility, defender, defenderAbility, c)
		damages[c] = best
	}
	var max uint = 0
	var bestCategory category.DamageCategory
	for _, r := range damages {
		if max < r.Damage() {
			max = r.Damage()
			bestCategory = r.Categoty()
		}
	}
	return damages[bestCategory]
}

func newBestAbilityService(sp species.Repository, mv move.Repository, ab ability.Repository, it item.Repository) BestAbilityService {
	return &baService{
		sp: sp,
		mv: mv,
		ab: ab,
		it: it,
	}
}

func newBestCategoryService(sp species.Repository, mv move.Repository, ab ability.Repository, it item.Repository) BestCategoryService {
	return &bcService{
		sp: sp,
		mv: mv,
		ab: ab,
		it: it,
	}
}

func newBestMoveService(sp species.Repository, mv move.Repository, ab ability.Repository, it item.Repository) BestMoveService {
	return &bmService{
		sp: sp,
		mv: mv,
		ab: ab,
		it: it,
	}
}

type bmService struct {
	sp species.Repository
	mv move.Repository
	ab ability.Repository
	it item.Repository
}

func (b *bmService) BestMove(attacker, attackerAbility, defender, defenderAbility string, category category.DamageCategory) BestMoveResult {
	sv := newMovesService(b.sp, b.mv)
	moves := sv.Moves(attacker, category)

	damages := make(map[string]uint)
	ds := damage.NewSimpleService(b.sp, b.mv, b.ab, b.it)
	for _, move := range moves {
		if isExclude(move) {
			continue
		}
		damages[move] = ds.Calculate(attacker, defender, attackerAbility, defenderAbility, move)
	}
	maxMove := ""
	var max uint = 0
	for move, dmg := range damages {
		if max < dmg {
			max = dmg
			maxMove = move
		}
	}
	if max == 0 {
		return NoEffective()
	}
	return &bmResult{move: maxMove, damage: max, category: category}

}

func NoEffective() BestMoveResult {
	return &bmResult{move: "", damage: 0, category: category.Physical}
}

type bmResult struct {
	move     string
	ability  string
	damage   uint
	category category.DamageCategory
}

func (r *bmResult) Ability() string {
	return r.ability
}
func (r *bmResult) Move() string {
	return r.move
}
func (r *bmResult) Categoty() category.DamageCategory {
	return r.category
}
func (r *bmResult) Damage() uint {
	return r.damage
}
func (r *bmResult) HasEffective() bool {
	return r.Damage() > 0
}

func isExclude(move string) bool {
	for _, exclude := range excludeMoves {
		if exclude == move {
			return true
		}
	}
	return false
}

var excludeMoves []string

func init() {
	excludeMoves = []string{
		"ギガインパクト",
		"はかいこうせん",
		"じばく",
		"だいばくはつ",
	}
}
