package stats

type statsCalculator func(l Level, s Species, i Individual, b BasePoint) uint

func calcShadinjaHP(l Level, s Species, i Individual, b BasePoint) uint {
	return 1
}
func calcHP(l Level, s Species, i Individual, b BasePoint) uint {
	x := uint(s)*2 + uint(i) + uint(b)/4
	y := float32(x)*float32(l)/100.0 + float32(l) + 10
	return uint(y)
}
func calcFlat(l Level, s Species, i Individual, b BasePoint) uint {
	return uint(flatFloat(l, s, i, b))
}
func calcUpper(l Level, s Species, i Individual, b BasePoint) uint {
	return uint(flatFloat(l, s, i, b) * 1.1)
}
func calcLower(l Level, s Species, i Individual, b BasePoint) uint {
	return uint(flatFloat(l, s, i, b) * 0.9)
}

func flatFloat(l Level, s Species, i Individual, b BasePoint) float32 {
	x := uint(s)*2 + uint(i) + uint(b)/4
	y := float32(x)*float32(l)/100.0 + 5
	return y
}
