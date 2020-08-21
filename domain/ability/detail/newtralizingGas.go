package detail

// かがくへんかガス
// 全てのとくせいを無効にする
type newtralizingGas struct {
	abilityImpl
}

func (a *newtralizingGas) onField(er eraser) {
	er.erase()
}
