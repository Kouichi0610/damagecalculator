package usableItem

import (
	"damagecalculator/domain/item"
	"damagecalculator/infra/local"
	"testing"
)

func Test_使用可能アイテム一覧(t *testing.T) {
	sv := NewService(local.Item())
	list := sv.List("ピカチュウ")
	if !contains(list, "しんかのきせき") {
		t.Error()
	}
	list = sv.List("ライチュウ")
	if contains(list, "しんかのきせき") {
		t.Error()
	}
}

func contains(list []item.ItemInfo, item string) bool {
	for _, n := range list {
		if n.Name() == item {
			return true
		}
	}
	return false
}
