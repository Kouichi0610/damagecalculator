package stats

import (
	"damagecalculator/domain/basepoints"
)

type statsCalculator func(l Level, s Species, i Individual, b basePoints.BasePoint) uint

// 最高値(上昇補正&基礎ポイント+252)
func Highest(l Level, s Species) uint {
	return calcUpper(l, s, Individual(31), basePoints.NewBasePoint(252))
}

// 最低値(下降補正&個体値0&基礎ポイント無振り)
func Lowest(l Level, s Species) uint {
	return calcLower(l, s, Individual(0), basePoints.NewBasePoint(0))
}

// 補正なし無振り
func Raw(l Level, s Species) uint {
	return calcFlat(l, s, Individual(31), basePoints.NewBasePoint(0))
}

// Raw & 基礎ポイント+4
func Raw4(l Level, s Species) uint {
	return calcFlat(l, s, Individual(31), basePoints.NewBasePoint(4))
}

// Raw & 基礎ポイント+252
func Raw252(l Level, s Species) uint {
	return calcFlat(l, s, Individual(31), basePoints.NewBasePoint(252))
}

func calcShadinjaHP(l Level, s Species, i Individual, b basePoints.BasePoint) uint {
	return 1
}
func calcHP(l Level, s Species, i Individual, b basePoints.BasePoint) uint {
	x := uint(s)*2 + uint(i) + uint(b)/4
	y := float32(x)*float32(l)/100.0 + float32(l) + 10
	return uint(y)
}
func calcFlat(l Level, s Species, i Individual, b basePoints.BasePoint) uint {
	x := uint(s)*2 + uint(i) + uint(b)/4
	y := uint(float32(x)*float32(l))/100 + 5
	return y
}
func calcUpper(l Level, s Species, i Individual, b basePoints.BasePoint) uint {
	return calcFlat(l, s, i, b) * 11 / 10
}
func calcLower(l Level, s Species, i Individual, b basePoints.BasePoint) uint {
	return calcFlat(l, s, i, b) * 9 / 10
}
