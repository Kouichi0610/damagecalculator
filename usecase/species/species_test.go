package species

import (
	"damagecalculator/infra/local"
	"reflect"
	"testing"
)

func Test_Species(t *testing.T) {
	rp := local.Species()
	loader := NewLoader(rp)
	if loader == nil {
		t.Error()
	}
	sp, err := loader.Get("ジャラランガ")
	if sp == nil || err != nil {
		t.Error()
		return
	}
	if sp.Name() != "ジャラランガ" {
		t.Errorf("%s", sp.Name())
	}
	if !reflect.DeepEqual(sp.Types(), []string{"ドラゴン", "かくとう"}) {
		t.Errorf("%v", sp.Types())
	}
	if !reflect.DeepEqual(sp.Species(), []uint{75, 110, 125, 100, 105, 85}) {
		t.Errorf("%v", sp.Species())
	}
	if !reflect.DeepEqual(sp.Abilities(), []string{"ぼうだん", "ぼうおん", "ぼうじん"}) {
		t.Errorf("%v", sp.Abilities())
	}
	if sp.Weight() != 78.2 {
		t.Errorf("%f", sp.Weight())
	}

	// 例外チェック
	sp, err = loader.Get("unknown")
	if err == nil {
		t.Error()
	}
}

// TODO:計算式
