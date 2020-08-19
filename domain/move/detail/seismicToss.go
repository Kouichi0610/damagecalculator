package detail

// ちきゅうなげ(レベル固定ダメージ)
type seismicToss struct {
	*defaultMove
}

func (s *seismicToss) Calculate(level, power, attack, defense uint) []uint {
	return []uint{level}
}
