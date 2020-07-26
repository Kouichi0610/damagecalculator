package types

import "testing"

func Test_string(t *testing.T) {
	if typeToString(Normal) != "ノーマル" {
		t.Error()
	}
	if typeToString(Fire) != "ほのお" {
		t.Error()
	}
	if typeToString(Water) != "みず" {
		t.Error()
	}
	if typeToString(Electric) != "でんき" {
		t.Error()
	}
	if typeToString(Grass) != "くさ" {
		t.Error()
	}
	if typeToString(Ice) != "こおり" {
		t.Error()
	}
	if typeToString(Fighting) != "かくとう" {
		t.Error()
	}
	if typeToString(Poison) != "どく" {
		t.Error()
	}
	if typeToString(Ground) != "じめん" {
		t.Error()
	}
	if typeToString(Flying) != "ひこう" {
		t.Error()
	}
	if typeToString(Psychic) != "エスパー" {
		t.Error()
	}
	if typeToString(Bug) != "むし" {
		t.Error()
	}
	if typeToString(Rock) != "いわ" {
		t.Error()
	}
	if typeToString(Ghost) != "ゴースト" {
		t.Error()
	}
	if typeToString(Dragon) != "ドラゴン" {
		t.Error()
	}
	if typeToString(Dark) != "あく" {
		t.Error()
	}
	if typeToString(Steel) != "はがね" {
		t.Error()
	}
	if typeToString(Fairy) != "フェアリー" {
		t.Error()
	}

	if typeToString(None) != "" {
		t.Error()
	}
}
