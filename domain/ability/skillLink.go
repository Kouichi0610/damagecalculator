package ability

import (
	"damagecalculator/domain/skill"
)

type skillLink struct {
	ability
}

// スキルリンク
// 攻撃回数2～5の技を5にする
func (a *skillLink) RewriteSkillData(st skill.SkillData) *skill.SkillData {
	res := st
	if !a.ability.isAttacker {
		return &res
	}
	if !(st.CountMin == 2 && st.CountMax == 5) {
		return &res
	}

	res.CountMin = 5
	res.CountMax = 5
	return &res
}
