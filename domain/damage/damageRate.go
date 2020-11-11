package damage

import (
	"fmt"
	"math"
)

/*
	ダメージのHPに対する割合、何回で倒せるか確定数を取得
*/
type DamageRate interface {
	RateMin() float64
	RateMax() float64
	DetermineCount() uint
	String() string
}

type damageRate struct {
	hp  float64
	dmg Damages
}

type noDamageRate struct {
}

func NewNoDamageRate() DamageRate {
	return new(noDamageRate)
}

func NewDamageRate(hp uint, dmg Damages) DamageRate {
	return &damageRate{
		hp:  float64(hp),
		dmg: dmg,
	}
}

func (d *damageRate) RateMin() float64 {
	if d.hp == 0 {
		return 0
	}
	return calcRate(d.dmg.Min(), d.hp)
}

func (d *damageRate) RateMax() float64 {
	if d.hp == 0 {
		return 0
	}
	return calcRate(d.dmg.Max(), d.hp)
}

func calcRate(dmg uint, hp float64) float64 {
	tmp := float64(dmg) / hp * 1000
	return math.Round(tmp) / 10
}

func (d *damageRate) DetermineCount() uint {
	hp := uint(d.hp)
	min := d.dmg.Min()
	if min == 0 {
		return 0
	}
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

func (d *noDamageRate) RateMin() float64 {
	return 0
}
func (d *noDamageRate) RateMax() float64 {
	return 0
}
func (d *noDamageRate) DetermineCount() uint {
	return 0
}
func (d *noDamageRate) String() string {
	return "0%"
}
