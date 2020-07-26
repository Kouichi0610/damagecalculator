package types

/*
 * magnifierインスタンスとenum Typeの相互変換
 *
 */
import (
	"fmt"
	"reflect"
)

type Type uint

const (
	None Type = iota
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

var magnifiers map[Type]magnifier

func init() {
	magnifiers = map[Type]magnifier{
		Normal:   new(normal),
		Fire:     new(fire),
		Water:    new(water),
		Electric: new(electric),
		Grass:    new(grass),
		Ice:      new(ice),
		Fighting: new(fighting),
		Poison:   new(poison),
		Ground:   new(ground),
		Flying:   new(flying),
		Psychic:  new(psychic),
		Bug:      new(bug),
		Rock:     new(rock),
		Ghost:    new(ghost),
		Dragon:   new(dragon),
		Dark:     new(dark),
		Steel:    new(steel),
		Fairy:    new(fairy),
	}
}

func fromType(id Type) (magnifier, error) {
	res, ok := magnifiers[id]
	if !ok {
		return nil, fmt.Errorf("%d not supported.", id)
	}
	return res, nil
}

func toType(d magnifier) (Type, error) {
	for key, value := range magnifiers {
		if reflect.TypeOf(value) == reflect.TypeOf(d) {
			return key, nil
		}
	}
	return 0, fmt.Errorf("%T not supported.", d)
}
