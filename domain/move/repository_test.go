package move

import (
	"damagecalculator/domain/move/attribute"
	"damagecalculator/domain/move/category"
	"damagecalculator/domain/move/detail"
	"damagecalculator/domain/move/target"
	"damagecalculator/domain/types"
	"testing"
)

/*
	TODO:detail.Detailで作成されたインターフェイスがMoveとしての機能を一通りそろえていること
	詳細はDetail側で行っている
*/
func Test_CreateMove(t *testing.T) {
	f := &MoveFactory{
		Name:      "サイコショック",
		Power:     80,
		Type:      types.Psychic,
		Accuracy:  100,
		Category:  category.PsycoShock,
		CountMin:  1,
		CountMax:  1,
		Target:    target.Select,
		Attribute: attribute.NewAttribute(attribute.Remote, attribute.NoAttribute),
		Detail:    detail.Default,
	}
	mv, err := f.Create()
	if mv == nil {
		t.Error()
	}
	if err != nil {
		t.Error()
	}

	// Moveとして機能していること
	if mv.Attribute().IsContact() {
		t.Error()
	}
}
