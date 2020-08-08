package count

import "fmt"

// 攻撃回数からとりうるダメージを生成する
type AttackCount struct {
	min, max uint
}

func NewAttackCount(min, max uint) (*AttackCount, error) {
	if min == 0 || max == 0 {
		return nil, fmt.Errorf("0回は指定できない")
	}
	return &AttackCount{
		min: min,
		max: max,
	}, nil
}

func (a *AttackCount) Correct(dmgs []uint) []uint {
	// TODO:実装
	return dmgs
	/*
		res := make([]uint, 0)
		for i := a.min; i <= a.max; i++ {
			d := correct(dmgs, i)
			res = append(res, d...)
		}
		return res
	*/
}
