package speed

import (
	"damagecalculator/infra/local"
	"testing"
)

func Test_Speed(t *testing.T) {
	rp := local.Species()
	g := NewGenerator()

	res := g.Generate(rp)

	if len(res) == 0 {
		t.Error()
	}

	if res[0].Speed() != 5 {
		t.Error()
	}
	if res[0].Info() != "ツボツボ " {
		t.Error()
	}
}
