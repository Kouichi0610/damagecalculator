package detail

import (
	"damagecalculator/domain/move"
)

type skillLink struct {
	abilityImpl
}

// スキルリンク
// 攻撃回数2～5の技を5にする
func (a *skillLink) RewriteMoveFactory(mv move.MoveFactory) *move.MoveFactory {
	res := mv
	if !a.isAttacker {
		return &res
	}

	if !(mv.CountMin == 2 && mv.CountMax == 5) {
		return &res
	}
	res.CountMin = 5
	res.CountMax = 5
	return &res
}
