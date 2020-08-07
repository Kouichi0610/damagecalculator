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
	x := uint(s)*2 + uint(i) + uint(b)/4
	y := uint(float32(x)*float32(l))/100 + 5
	return y
}
func calcUpper(l Level, s Species, i Individual, b BasePoint) uint {
	return calcFlat(l, s, i, b) * 11 / 10
}
func calcLower(l Level, s Species, i Individual, b BasePoint) uint {
	return calcFlat(l, s, i, b) * 9 / 10
}
