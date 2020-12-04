package situation

import (
	"damagecalculator/infra/local"
	"testing"
)

func Test_Builder(t *testing.T) {
	builder := NewBuilder(local.Species(), local.Ability(), local.Move(), local.Item())
	if builder == nil {
		t.Error()
	}
	fields := &FieldCondition{
		Weather:      "なし",
		Field:        "なし",
		HasReflector: false,
		IsCritical:   false,
	}
	attacker := &PokeParams{
		Name:        "ピカチュウ",
		Individuals: "Max",
		BasePoints:  []uint{6, 252, 0, 0, 0, 252},
		Ranks:       []int{0, 0, 0, 0, 0},
		Ability:     "せいでんき",
		Item:        "None",
		Nature:      "のんき",
		Condition:   "なし",
	}
	defender := &PokeParams{
		Name:        "シェルダー",
		Individuals: "Max",
		BasePoints:  []uint{252, 0, 252, 0, 0, 0},
		Ranks:       []int{0, 0, 0, 0, 0},
		Ability:     "げきりゅう",
		Item:        "None",
		Condition:   "なし",
	}
	move := "ボルテッカー"

	st, err := builder.ToSituation(50, attacker, defender, move, fields)
	if err != nil {
		t.Error()
	}
	if st == nil {
		t.Error()
	}
	if st.Attacker().Attack().Value() != 107 {
		t.Errorf("%d", st.Attacker().Attack().Value())
	}
	if st.Defender().SpDefense().Value() != 45 {
		t.Errorf("%d", st.Defender().SpDefense().Value())
	}

	// TODO:エラーチェック厳密化
	//t.Errorf("Attacker:%s", st.Attacker())
	//t.Errorf("Defender:%s", st.Defender())
}
