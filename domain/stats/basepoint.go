package stats

import "fmt"

/*
	基礎ポイント
	・HP、こうげき、ぼうぎょ、とくこう、とくぼう、すばやさそれぞれに設定
	・個別の範囲は0～252
	・合計値は510以内であること
*/
var MinBasePoint uint = 0
var MaxBasePoint uint = 252

type BasePoint uint

var totalBasePoint uint = 510

type BasePointStats struct {
	hp, at, df, sa, sd, sp BasePoint
}

func NewBasePointStats(hp, at, df, sa, sd, sp uint) (*BasePointStats, error) {
	total := hp + at + df + sa + sd + sp
	if total > totalBasePoint {
		return nil, fmt.Errorf("capacity over %d\n", total)
	}
	res := new(BasePointStats)
	res.hp = newBasePoint(hp)
	res.at = newBasePoint(at)
	res.df = newBasePoint(df)
	res.sa = newBasePoint(sa)
	res.sd = newBasePoint(sd)
	res.sp = newBasePoint(sp)
	return res, nil
}

func (b *BasePointStats) HP() BasePoint {
	return b.hp
}
func (b *BasePointStats) Attack() BasePoint {
	return b.at
}
func (b *BasePointStats) Defense() BasePoint {
	return b.df
}
func (b *BasePointStats) SpAttack() BasePoint {
	return b.sa
}
func (b *BasePointStats) SpDefense() BasePoint {
	return b.sd
}
func (b *BasePointStats) Speed() BasePoint {
	return b.sp
}

func newBasePoint(value uint) BasePoint {
	if value > MaxBasePoint {
		value = MaxBasePoint
	}
	res := BasePoint(value)
	return res
}
