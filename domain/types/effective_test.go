package types

import (
	"testing"
)

func Test_Effective(t *testing.T) {
	// 通常
	flat := FlatEffective()
	if flat != 1.0 {
		t.Error()
	}
	// 効果はばつぐん
	super := SuperEffective()
	if super != 2.0 {
		t.Error()
	}
	// いまひとつ
	notVery := NotVeryEffective()
	if notVery != 0.5 {
		t.Error()
	}
	// こうかなし
	noEffective := NoEffective()
	if noEffective != 0.0 {
		t.Error()
	}

	calc := super * notVery
	if calc != 1.0 {
		t.Error()
	}
}
