package count

import (
	"fmt"
	"math"
)

// 攻撃回数からとりうるダメージを生成する
type AttackCount struct {
	min, max uint
}

func (a *AttackCount) Min() uint {
	return a.min
}
func (a *AttackCount) Max() uint {
	return a.max
}

func NewAttackCount(min, max uint) (*AttackCount, error) {
	if min == 0 || max == 0 {
		return nil, fmt.Errorf("0回は指定できない")
	}

	return &AttackCount{
		min: minUint(min, max),
		max: maxUint(min, max),
	}, nil
}

func maxUint(a, b uint) uint {
	if a > b {
		return a
	}
	return b
}
func minUint(a, b uint) uint {
	if a < b {
		return a
	}
	return b
}

/*
	min回の攻撃ダメージ～max回の攻撃ダメージの総当たり配列
*/
func (a *AttackCount) Correct(dmgs []uint) []uint {
	res := make([]uint, 0)
	for i := a.min; i <= a.max; i++ {
		ap := correct(dmgs, i)
		res = append(res, ap...)
	}
	return res
}

// d ダメージ c 回数
func correct(d []uint, c uint) []uint {
	res := make([]uint, 0)
	dg := newDigit(int(len(d)), 0)
	resLength := int(math.Pow(float64(len(d)), float64(c)))

	for dg.Count() != resLength {
		idx := dg.ToArray(int(c))
		a := uint(0)
		for i := 0; i < int(c); i++ {
			a += d[idx[i]]
		}
		dg.Increment()

		res = append(res, a)
	}

	return res
}
