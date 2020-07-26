package types

import "reflect"

// タイプ相性
type Types struct {
	t []magnifier
}

func NewTypes(args ...Type) *Types {
	res := new(Types)
	res.t = make([]magnifier, 0)

	if len(args) == 0 {
		return res
	}

	for _, n := range args {
		ty, _ := fromType(n)
		if ty == nil {
			break
		}
		res.t = append(res.t, ty)
	}
	return res
}

// ダメージ倍率を計算
func (at *Types) Magnification(df *Types) Effective {
	if len(at.t) == 0 || len(df.t) == 0 {
		return FlatEffective()
	}
	res := FlatEffective()
	for _, a := range at.t {
		for _, d := range df.t {
			m := a.Magnification(d)
			res *= m
		}
	}
	return res
}

func (t *Types) PartialMatch(o *Types) bool {
	for _, a := range t.t {
		for _, b := range o.t {
			if reflect.TypeOf(a) == reflect.TypeOf(b) {
				return true
			}
		}
	}
	return false
}
