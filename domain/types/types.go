/*
	タイプ相性
*/
package types

import "strings"

type (
	// タイプ(ノーマル、ほのお、みず...)
	Type uint

	// 複合タイプを扱う
	Types struct {
		ty []Type
	}

	// タイプが他のタイプを攻撃したときの効果(こうかはばつぐん、いまひとつ、...)
	Effective float32
)

const (
	Normal Type = iota
	Fire
	Water
	Electric
	Grass
	Ice
	Fighting
	Poison
	Ground
	Flying
	Psychic
	Bug
	Rock
	Ghost
	Dragon
	Dark
	Steel
	Fairy
)

const (
	Flat        Effective = 1.0
	Super       Effective = 2.0
	NotVery     Effective = 0.5
	NoEffective Effective = 0.0
)

func NewTypes(args ...Type) *Types {
	res := new(Types)
	res.ty = make([]Type, 0)
	res.ty = append(res.ty, args...)
	if len(args) == 0 {
		return res
	}
	return res
}

// ダメージ倍率を取得
func (t Type) Magnification(defense Type) Effective {
	return maps[t](defense)
}
func (t Type) String() string {
	return names[t]
}

func (t Type) Types() *Types {
	return NewTypes(t)
}

// ダメージ倍率を計算
func (at *Types) Magnification(df *Types) Effective {
	if len(at.ty) == 0 || len(df.ty) == 0 {
		return Flat
	}
	res := Flat
	for _, a := range at.ty {
		for _, d := range df.ty {
			m := a.Magnification(d)
			res *= m
		}
	}
	return res
}

func (ty *Types) Has(t Type) bool {
	for _, ty := range ty.ty {
		if ty == t {
			return true
		}
	}
	return false
}

// 部分一致
func (t *Types) PartialMatch(o *Types) bool {
	for _, a := range t.ty {
		for _, b := range o.ty {
			if a == b {
				return true
			}
		}
	}
	return false
}

// 全部一致(順番は問わないこと)
func (t *Types) Equal(o *Types) bool {
	if len(t.ty) != len(o.ty) {
		return false
	}
	for _, ty := range o.ty {
		if !t.Has(ty) {
			return false
		}
	}
	for _, ty := range t.ty {
		if !o.Has(ty) {
			return false
		}
	}
	return true
}

func (t *Types) String() string {
	s := new(strings.Builder)

	if len(t.ty) == 0 {
		return "(none)"
	}

	for i, ty := range t.ty {
		s.WriteString(ty.String())
		if i != len(t.ty)-1 {
			s.WriteString("/")
		}
	}
	return s.String()
}

func (t *Types) Strings() []string {
	res := make([]string, 0)
	for _, ty := range t.ty {
		res = append(res, ty.String())
	}
	return res
}
