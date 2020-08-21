package ability

import (
	"damagecalculator/domain/move"
	"damagecalculator/domain/move/count"
)

type skillLink struct {
	ability
}

// スキルリンク
// 攻撃回数2～5の技を5にする
func (a *skillLink) RewriteMoveFactory(mv move.MoveFactory) *move.MoveFactory {
	res := mv
	if !a.ability.isAttacker {
		return &res
	}

	if !(mv.Count.Min() == 2 && mv.Count.Max() == 5) {
		return &res
	}
	cnt, _ := count.NewAttackCount(5, 5)
	res.Count = cnt
	return &res
}
