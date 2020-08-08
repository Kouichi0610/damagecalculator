package skill

// ちきゅうなげ(レベル固定ダメージ)
type seismicToss struct {
	skill
}

func (s *seismicToss) Calculate(level, power, attack, defense uint) []uint {
	return []uint{level}
}
