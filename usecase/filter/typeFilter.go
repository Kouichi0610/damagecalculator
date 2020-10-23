// タイプによる絞り込み
package filter

import "damagecalculator/domain/types"

// 絞り込み結果
type FilterResult uint
type Type uint

const (
	All Type = iota
	Normal
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

type TypeFilter interface {
	Filter([]types.Type) FilterResult
}

func createTypeFilter(t Type) TypeFilter {
	if t == All {
		return new(allType)
	}
	ty := typeindex[t]
	return &withType{ty: ty}
}

type allType struct {
}

func (f *allType) Filter([]types.Type) FilterResult {
	return Include
}

type withType struct {
	ty types.Type
}

func (f *withType) Filter(t []types.Type) FilterResult {
	for _, ty := range t {
		if f.ty == ty {
			return Include
		}
	}
	return Exclude
}

var typeindex map[Type]types.Type

func init() {
	typeindex = make(map[Type]types.Type, 0)
	typeindex[Normal] = types.Normal
	typeindex[Fire] = types.Fire
	typeindex[Water] = types.Water
	typeindex[Electric] = types.Electric
	typeindex[Grass] = types.Grass
	typeindex[Ice] = types.Ice
	typeindex[Fighting] = types.Fighting
	typeindex[Poison] = types.Poison
	typeindex[Ground] = types.Ground
	typeindex[Flying] = types.Flying
	typeindex[Psychic] = types.Psychic
	typeindex[Bug] = types.Bug
	typeindex[Rock] = types.Rock
	typeindex[Ghost] = types.Ghost
	typeindex[Dragon] = types.Dragon
	typeindex[Dark] = types.Dark
	typeindex[Steel] = types.Steel
	typeindex[Fairy] = types.Fairy
}
