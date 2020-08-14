package ability

// かがくへんかガス
// 全てのとくせいを無効にする
type newtralizingGas struct {
	ability
}

func (a *newtralizingGas) onField(f AbilityField) AbilityField {
	return &abilityField{
		at: &ability{true},
		df: &ability{false},
	}
}
