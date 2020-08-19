package move

import (
	"testing"
)

/*
	TODO:detail.Detailで作成されたインターフェイスがMoveとしての機能を一通りそろえていること
	詳細はDetail側で行っている
*/
func Test_CreateMove(t *testing.T) {
	// TODO:これがとおるのは良くないか
	f := &MoveFactory{}
	mv, err := f.Create()
	if mv == nil {
		t.Error()
	}
	if err != nil {
		t.Error()
	}

	// TODO:パラメータ確認

}
