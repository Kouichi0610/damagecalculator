package status

import (
	"fmt"
)

const RankMax = 6
const RankMin = -6

type rank int

// ランク補正付きステータス
type RankedValue struct {
	v uint
	r rank
}

func NewRankedValue(v uint, r int) *RankedValue {
	res := new(RankedValue)
	res.v = v
	res.r = newRank(r)
	return res
}

func (r *RankedValue) String() string {
	return fmt.Sprintf("%d(%d)", r.Value(), r.ignoreRank())
}

func (r *RankedValue) Value() uint {
	return r.r.RankedStats(r.v)
}

func (r *RankedValue) ignoreRank() uint {
	nr := newRank(0)
	return nr.RankedStats(r.v)
}

// +の補正値を無視する
func (r *RankedValue) IgnorePlusValue() uint {
	nr := r.r
	if nr > 0 {
		nr = newRank(0)
	}
	return nr.RankedStats(r.v)
}
func (r *RankedValue) IgnoreMinusValue() uint {
	nr := r.r
	if nr < 0 {
		nr = newRank(0)
	}
	return nr.RankedStats(r.v)
}

func newRank(r int) rank {
	if r < RankMin {
		r = RankMin
	}
	if r > RankMax {
		r = RankMax
	}
	return rank(r)
}

/*
	ランク補正を掛けた能力値を取得
	+6: 8/2
	+5: 7/2
	+4: 6/2
	+3: 5/2
	+2: 4/2
	+1: 3/2
	+0: 2/2
	-1: 2/3
	-2: 2/4
	-3: 2/5
	-4: 2/6
	-5: 2/7
	-6: 2/8
*/
func (r rank) RankedStats(s uint) uint {
	if r < 0 {
		x := int(-r) + 2
		if x == 0 {
			return s
		}
		res := int(s) * 2 / x
		return uint(res)
	}
	{
		x := int(r) + 2
		res := int(s) * x / 2
		return uint(res)
	}
}
