package damage

import (
	"damagecalculator/infra/local"

	"testing"
)

func Test_SimpleService(t *testing.T) {
	sp := local.Species()
	mv := local.Move()
	ab := local.Ability()
	it := local.Item()
	sv := NewSimpleService(sp, mv, ab, it)
	if sv == nil {
		t.Error()
	}

	dmg := sv.Calculate("ピカチュウ", "ゼニガメ", "せいでんき", "げきりゅう", "かみなりパンチ")
	if dmg != 78 {
		t.Error()
	}
	dmg = sv.Calculate("ピカチュウ", "ゼニガメ", "ちからもち", "げきりゅう", "かみなりパンチ")
	if dmg != 153 {
		t.Error()
	}
	dmg = sv.Calculate("ピカチュウ", "ディグダ", "ちからもち", "げきりゅう", "かみなりパンチ")
	if dmg != 0 {
		t.Error()
	}
}
