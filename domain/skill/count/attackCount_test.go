package count

import (
	"math"
	"reflect"
	"testing"
)

func Test_AttackCount(t *testing.T) {
	dmgs := []uint{1, 2, 3, 4, 5, 6, 7, 8}
	ac, _ := NewAttackCount(1, 1)

	expects := ac.Correct(dmgs)
	if !reflect.DeepEqual(dmgs, expects) {
		t.Errorf("%v", expects)
	}

	ac, _ = NewAttackCount(2, 5)
	// TODO:結構重くなるのでアルゴリズムの改善
	expects = ac.Correct(dmgs)
	var l float64
	for i := ac.min; i <= ac.max; i++ {
		l += math.Pow(float64(8), float64(i))
	}

	// 要素数と最小、最大値だけで良しとする
	length := len(expects)
	if length != int(l) {
		t.Errorf("%d", length)
	}
	if expects[0] != 2 {
		t.Error()
	}
	if expects[length-1] != 40 {
		t.Error()
	}
}

func Test_検証(t *testing.T) {
	expect := make([]uint, 0)
	dm := []uint{1, 2}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				for l := 0; l < 2; l++ {
					expect = append(expect, dm[i]+dm[j]+dm[k]+dm[l])
				}
			}
		}
	}
	actual := correct([]uint{1, 2}, 4)

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("%v", actual)
	}
}

func Test_xDigit(t *testing.T) {
	x := newDigit(2, 10)
	// "1010"
	res := x.ToArray(4)
	if !reflect.DeepEqual(res, []int{1, 0, 1, 0}) {
		t.Errorf("%v", res)
	}

	// 0x00c8
	x = newDigit(16, 200)
	res = x.ToArray(4)
	if !reflect.DeepEqual(res, []int{0, 0, 12, 8}) {
		t.Errorf("%v", res)
	}
}
func Test_xDigitString(t *testing.T) {
	expects := []string{
		"0000",
		"0001",
		"0010",
		"0011",
		"0100",
		"0101",
		"0110",
		"0111",
		"1000",
		"1001",
		"1010",
		"1011",
		"1100",
		"1101",
		"1110",
		"1111",
	}
	x := newDigit(2, 0)
	for i := 0; i < 16; i++ {
		s := x.String(4)
		if s != expects[i] {
			t.Errorf("%s", s)
		}
		x.Increment()
	}

	x = newDigit(16, 4281)
	s := x.String(4)
	if s != "10b9" {
		t.Errorf("%s", s)
	}
}

func Test_AttackCount_エラーチェック(t *testing.T) {
	a, e := NewAttackCount(0, 1)
	if a != nil {
		t.Error()
	}
	if e == nil {
		t.Error()
	}
	a, e = NewAttackCount(1, 0)
	if a != nil {
		t.Error()
	}
	if e == nil {
		t.Error()
	}

	a, e = NewAttackCount(1, 1)
	if a == nil {
		t.Error()
	}
	if e != nil {
		t.Error()
	}
}
