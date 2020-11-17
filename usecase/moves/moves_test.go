package moves

import (
	"damagecalculator/infra/local"
	"testing"
)

func Test_Moves(t *testing.T) {
	ld := NewLoader(local.Species(), local.Move())
	if ld == nil {
		t.Error()
	}

	moves := ld.Moves("アンノーン")
	if len(moves) != 1 {
		t.Error()
	}

	if moves[0].Name != "めざめるパワー" {
		t.Error()
	}
}
