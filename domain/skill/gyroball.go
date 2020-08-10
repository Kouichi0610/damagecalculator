package skill

/*
	ジャイロボール
	相手よりすばやさが低いほど威力が上がる
	最大150まで
*/

type gyroBall struct {
	skill
}

func (s *gyroBall) Power(st SituationChecker) uint {
	my := st.Attacker().Speed().Value()
	en := st.Defender().Speed().Value()
	return gyroBallPower(my, en)
}

func gyroBallPower(my, en uint) uint {
	res := uint(25.0*float64(en)/float64(my)) + 1
	if res >= 150 {
		return 150
	}
	return res
}
