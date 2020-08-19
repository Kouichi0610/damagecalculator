package move

import (
	"testing"
)

func Test_CreateMove(t *testing.T) {
	f := &MoveFactory{}
	mv := f.Create()
	if mv == nil {
		t.Error()
	}

	// TODO:パラメータ確認

}
