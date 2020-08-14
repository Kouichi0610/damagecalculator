package ability

import "damagecalculator/domain/skill/count"

type skillLink struct {
	ability
}

// スキルリンク
// 攻撃回数2～5の技を5にする
func (a *skillLink) AttackCount(cnt *count.AttackCount) *count.AttackCount {
	if !a.ability.isAttacker {
		return cnt
	}
	if !(cnt.Min() == 2 && cnt.Max() == 5) {
		return cnt
	}
	res, _ := count.NewAttackCount(5, 5)
	return res
}
