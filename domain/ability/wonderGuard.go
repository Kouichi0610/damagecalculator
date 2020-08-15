package ability

import (
	"damagecalculator/domain/ability/situation"
	"damagecalculator/domain/corrector"
)

// ふしぎなまもり
// 効果抜群以外のダメージを受けない
type wonderGuard struct {
	ability
}

func (a *wonderGuard) Correctors(st situation.SituationChecker) []corrector.Corrector {
	// 防御側でのみ有効
	if a.ability.isAttacker {
		return nil
	}

	ef := st.SkillEffective()
	if ef.IsSuper() {
		return nil
	}
	return []corrector.Corrector{corrector.NewDamage(0)}
}
