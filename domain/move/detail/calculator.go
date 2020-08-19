package detail

func calculate(level, power, attack, defense uint) uint {
	a := level*2/5 + 2
	a = a * power * attack / defense
	a = a/50 + 2
	return a
}

func calculateArray(level, power, attack, defense uint) []uint {
	d := calculate(level, power, attack, defense)
	res := make([]uint, 0)
	for i := 0.85; i <= 1.0; i += 0.01 {
		t := uint(i * float64(d))
		res = append(res, t)
	}
	return res
}
