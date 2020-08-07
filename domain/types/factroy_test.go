package types

import "testing"

// エラーを投げるためのダミーインターフェイス
type dummy struct {
}

func (t dummy) Magnification(defense magnifier) Effective {
	return flatEffective()
}

// 正常範囲チェック intとDamageMagnifier間で相互変換できること
func Test_Factory(t *testing.T) {
	for i := Normal; i <= Fairy; i++ {
		d, _ := fromType(i)
		id, _ := toType(d)
		if id != i {
			t.Errorf("ToInt failed %d", id)
		}
	}
}

// エラーチェック
func Test_IllegalFromInt(t *testing.T) {
	_, errorUnder := fromType(0)
	if errorUnder == nil {
		t.Errorf("failed")
	}
	_, errorOver := fromType(19)
	if errorOver == nil {
		t.Errorf("failed")
	}
}

// エラーチェック
func Test_IllegalToInt(t *testing.T) {
	_, error := toType(new(dummy))
	if error == nil {
		t.Errorf("failed")
	}
}
