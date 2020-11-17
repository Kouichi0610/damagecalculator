package supposition

import (
	"damagecalculator/infra/local"
	"testing"
)

func Test_AttackersMaker(t *testing.T) {
	maker := newAttackersMaker()
	if maker == nil {
		t.Error()
	}
	/*
		res, err := maker.Defenders("ピカチュウ", "シャドーパンチ")
		if err != nil {
			t.Error()
		}

		if len(res) == 0 {
			t.Error()
			return
		}
		if res[0].Name != "シャンデラ" {
			t.Errorf("%s", res[0].Name)
		}
	*/
}

func newAttackersMaker() AttackersMaker {
	return NewAttackersMaker(local.Names(), local.Species(), local.Move(), local.Ability(), local.Item())
}
