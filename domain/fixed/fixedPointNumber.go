// fixed-point-number
// 固定小数点に関する演算
package fixed

import "fmt"

type Category uint

const (
	Drop4Pick5     Category = iota // 小数点以下四捨五入
	Drop5Pick5Over                 // 小数点以下五捨五超入
	Omit                           // 小数点以下切り捨て
)

// numer, denom
type FixPN interface {
	Correct(d uint) uint
}

func NewFixPN(magnification float64, category Category) (FixPN, error) {
	if magnification < 0 {
		return nil, fmt.Errorf("%f out of range.", magnification)
	}
	m := uint(magnification * one)
	switch category {
	case Drop4Pick5:
		return &fixPN_drop4_pick5{m: m}, nil
	case Drop5Pick5Over:
		return &fixPN_drop5_pick5over{m: m}, nil
	case Omit:
		return &fixPN_omit{m: m}, nil
	}
	return nil, fmt.Errorf("%d not supported.", category)
}

type fixPN_drop4_pick5 struct {
	m uint
}
type fixPN_drop5_pick5over struct {
	m uint
}
type fixPN_omit struct {
	m uint
}

func (f *fixPN_drop4_pick5) Correct(d uint) uint {
	res := d * f.m
	if res%one >= 2048 {
		return res/one + 1
	}
	return res / one
}
func (f *fixPN_drop5_pick5over) Correct(d uint) uint {
	res := d * f.m
	if res%one > 2048 {
		return res/one + 1
	}
	return res / one
}
func (f *fixPN_omit) Correct(d uint) uint {
	return d * f.m / one
}

const one = 4096
const half = 2048
