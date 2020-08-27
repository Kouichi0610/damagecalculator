package damage

import (
	"fmt"
)

/*
	倒せる攻撃確定数を取得
*/
type DamageRate interface {
	RateMin() float64
	RateMax() float64
	DetermineCount() uint
	String() string
}

type damageRate struct {
	hp  float64
	dmg *Damages
}

func NewDamageRate(hp uint, dmg *Damages) DamageRate {
	return &damageRate{
		hp:  float64(hp),
		dmg: dmg,
	}
}

func (d *damageRate) RateMin() float64 {
	return float64(d.dmg.Min()) / d.hp * 100
}

func (d *damageRate) RateMax() float64 {
	return float64(d.dmg.Max()) / d.hp * 100
}

func (d *damageRate) DetermineCount() uint {
	hp := uint(d.hp)
	min := d.dmg.Min()
	res := hp / min
	if hp%min > 0 {
		res++
	}
	return res
}

func (d *damageRate) String() string {
	if d.dmg.Min() == d.dmg.Max() {
		return fmt.Sprintf("%0.1f%% 確定数%d", d.RateMin(), d.DetermineCount())
	}
	return fmt.Sprintf("%0.1f%% ～ %0.1f%% 確定数%d", d.RateMin(), d.RateMax(), d.DetermineCount())
}
