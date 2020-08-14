package ability

// かたやぶり
// 攻撃時、防御側の特性を無視する
type moldBreaker struct {
	ability
}

func (a *moldBreaker) onAttack(f AbilityField) AbilityField {
	if !a.ability.isAttacker {
		return f
	}
	return &abilityField{
		at: a,
		df: &ability{false},
	}
}
