package detail

// かたやぶり
// 攻撃時、防御側の特性を無視する
type moldBreaker struct {
	abilityImpl
}

func (a *moldBreaker) onAttack(er eraser) {
	if !a.abilityImpl.isAttacker {
		return
	}
	er.erase()
}
