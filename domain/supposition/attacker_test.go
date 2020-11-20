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
	actuals := maker.Attackers("フーディン", "いしあたま")
	for _, actual := range actuals {
		t.Errorf("%s %s", actual.Param().Name, actual.Move())
	}
}

func newAttackersMaker() AttackersMaker {
	return NewAttackersMaker(local.Names(), local.Species(), local.Move(), local.Ability(), local.Item())
}
