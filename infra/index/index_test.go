package index

import (
	"testing"
)

func Test_Index(t *testing.T) {
	if indices == nil {
		t.Error()
	}
	if len(indices) == 0 {
		t.Error()
	}
	id, err := Index("pikachu")
	if err != nil {
		t.Errorf("%v", err)
	}
	if len(id) == 0 {
		t.Error()
	}

	id, err = Index("イシヘンジン")
	if err != nil {
		t.Errorf(err.Error())
	}
	if len(id) == 0 {
		t.Error()
	}

	id, err = Index("xxxx")
	if err == nil {
		t.Error()
	}
	if len(id) != 0 {
		t.Error()
	}
}
