package count

import (
	"math"
	"strings"
)

// x進数
type xDigit struct {
	cnt int
	x   int
}

func newDigit(x int, cnt int) *xDigit {
	return &xDigit{
		cnt: cnt,
		x:   x,
	}
}
func (d *xDigit) Increment() {
	d.cnt++
}
func (d *xDigit) Count() int {
	return d.cnt
}

func (d *xDigit) ToArray(keta int) []int {
	tmp := make([]int, keta)
	n := d.cnt
	for i := 0; i < keta; i++ {
		m := int(math.Pow(float64(d.x), float64(i+1)))
		r := n % m
		n -= r

		for j := 0; j < i; j++ {
			r /= d.x
		}
		tmp[i] = r
	}

	res := make([]int, 0)
	for i := len(tmp) - 1; i >= 0; i-- {
		res = append(res, tmp[i])
	}
	return res
}

func (d *xDigit) String(keta int) string {
	arr := d.ToArray(keta)
	b := strings.Builder{}
	for i := 0; i < keta; i++ {
		n := arr[i]

		var c byte
		if n >= 10 {
			c = 'a' + byte(n-10)

		} else {
			c = '0' + byte(n)
		}
		b.WriteByte(c)
	}
	return b.String()
}
