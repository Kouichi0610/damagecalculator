/*
	与えうるダメージ値一覧
*/
package damage

import (
	"fmt"
	"sort"
)

type Damages interface {
	Array() []uint
	Min() uint
	Max() uint
	String() string
}

type damages struct {
	d []uint
}

type damageSlice []uint

func NoDamage() Damages {
	res := new(damages)
	res.d = []uint{0}
	return res
}

func NewDamages(args []uint) Damages {
	res := new(damages)

	res.d = args
	sort.Sort(damageSlice(res.d))

	return res
}

func (d *damages) Array() []uint {
	res := make([]uint, 0)
	res = append(res, d.d...)
	return res
}

func (d *damages) Min() uint {
	return d.d[0]
}

func (d *damages) Max() uint {
	return d.d[len(d.d)-1]
}

func (d *damages) String() string {
	return fmt.Sprintf("%d ～ %d", d.Min(), d.Max())
}

func (p damageSlice) Len() int           { return len(p) }
func (p damageSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p damageSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
