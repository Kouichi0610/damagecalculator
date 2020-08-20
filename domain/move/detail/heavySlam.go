package detail

import (
	"damagecalculator/domain/status"
)

/*
	ヘビーボンバー
	攻撃側の重さが相手より重いほど威力が上がる
*/
type heavySlam struct {
	*defaultMove
}

func (s *heavySlam) Power(st SituationChecker) uint {
	at := st.Attacker().Weight()
	df := st.Defender().Weight()
	return heavySlamPower(at, df)
}

func heavySlamPower(at, df status.Weight) uint {
	m := float64(at / df)
	if m >= 5.0 {
		return 120
	}
	if m >= 4.0 {
		return 100
	}
	if m >= 3.0 {
		return 80
	}
	if m >= 2.0 {
		return 60
	}
	return 40
}
