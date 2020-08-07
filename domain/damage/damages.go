/*
	与えうるダメージ値一覧
*/
package damage

import "sort"

type Damages struct {
	d []uint
}

type damageSlice []uint

/*
func NewDamages(args ...uint) *Damages {
	res := new(Damages)
	res.d = make([]uint, 0)
	if len(args) == 0 {
		return res
	}
	for _, n := range args {
		res.d = append(res.d, n)
	}
	sort.Sort(damageSlice(res.d))
	return res
}
*/

func NewDamages(args []uint) *Damages {
	res := new(Damages)

	res.d = args
	sort.Sort(damageSlice(res.d))

	return res
}

func (d *Damages) Min() uint {
	return d.d[0]
}

func (d *Damages) Max() uint {
	return d.d[len(d.d)-1]
}

func (p damageSlice) Len() int           { return len(p) }
func (p damageSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p damageSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
