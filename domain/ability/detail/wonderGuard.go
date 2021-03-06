package detail

import (
	"damagecalculator/domain/ability/situation"
	"damagecalculator/domain/corrector"
)

// ふしぎなまもり
// 効果抜群以外のダメージを受けない
type wonderGuard struct {
	abilityImpl
}

func (a *wonderGuard) Correctors(st situation.SituationChecker) []corrector.Corrector {
	// 防御側でのみ有効
	if a.isAttacker {
		return nil
	}

	ef := st.MoveEffective()
	if ef.IsSuper() {
		return nil
	}
	return []corrector.Corrector{corrector.NewDamage(0)}
}
