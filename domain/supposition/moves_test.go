package supposition

import (
	"damagecalculator/domain/move/category"
	"damagecalculator/infra/local"
	"reflect"
	"testing"
)

// 全てのとくせいから最善のわざを
func Test_BestAbility(t *testing.T) {
	sv := newBestAbilityService(local.Species(), local.Move(), local.Ability(), local.Item())
	actual := sv.BestMove("マリルリ", "ヘルガー", "もらいび")
	if actual.Ability() != "ちからもち" {
		t.Errorf("%v", actual)
	}
}

// 全てのわざの最善
func Test_BestCategory(t *testing.T) {
	sv := newBestCategoryService(local.Species(), local.Move(), local.Ability(), local.Item())
	actual := sv.BestMove("オニゴーリ", "てんねん", "フリージオ", "てんねん")
	if actual.Categoty() != category.Physical {
		t.Errorf("%v", actual)
	}

	actual = sv.BestMove("オニゴーリ", "てんねん", "クレベース", "てんねん")
	if actual.Categoty() != category.Special {
		t.Errorf("%v", actual)
	}
}

// カテゴリごとの最大ダメージ
func Test_BestMove(t *testing.T) {
	sv := newBestMoveService(local.Species(), local.Move(), local.Ability(), local.Item())
	actual := sv.BestMove("ミュウツー", "てんねん", "マタドガス", "ふゆう", category.PsycoShock)
	if actual.Move() != "サイコブレイク" {
		t.Errorf("%v", actual)
	}
	if actual.Categoty() != category.PsycoShock {
		t.Error()
	}

	actual = sv.BestMove("ガブリアス", "てんねん", "ニンフィア", "かがくへんかガス", category.Physical)
	if actual.Move() != "アイアンテール" {
		t.Errorf("%v", actual)
	}

	actual = sv.BestMove("コスモッグ", "てんねん", "マタドガス", "ふゆう", category.Physical)
	if actual.HasEffective() {
		t.Errorf("%v", actual)
	}
}

// カテゴリごとに所有する技一覧
func Test_Moves(t *testing.T) {
	sv := newMovesService(local.Species(), local.Move())

	actual := sv.Moves("ミュウツー", category.PsycoShock)
	if !reflect.DeepEqual(actual, []string{"サイコショック", "サイコブレイク"}) {
		t.Errorf("%v", actual)
	}
}
