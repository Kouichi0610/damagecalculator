package damage

import (
	AbilityDetail "damagecalculator/domain/ability/detail"
	"damagecalculator/domain/field"
	"damagecalculator/domain/item"
	"damagecalculator/domain/move"
	"damagecalculator/domain/move/attribute"
	"damagecalculator/domain/move/category"
	"damagecalculator/domain/move/detail"
	"damagecalculator/domain/move/target"
	"damagecalculator/domain/situation"
	"damagecalculator/domain/status"
	"damagecalculator/domain/types"

	"testing"
)

func defaultArgs() *situation.SituationData {
	s := &move.MoveFactory{
		Name:      "unknown",
		Power:     80,
		Type:      types.Fire,
		Accuracy:  100,
		Category:  category.Physical,
		CountMin:  1,
		CountMax:  1,
		Target:    target.Select,
		Attribute: attribute.NewAttribute(attribute.Contact, attribute.NoAttribute),
		Detail:    detail.Default,
	}
	at := &status.StatusData{
		Lv:    50,
		Types: []types.Type{types.Bug},
		HP:    100, Attack: 180, Defense: 20, SpAttack: 100, SpDefense: 50, Speed: 130,
		AttackRank: 0, DefenseRank: 0, SpAttackRank: 0, SpDefenseRank: 0, SpeedRank: 0,
		Weight: 100.0,
	}
	df := &status.StatusData{
		Lv:    50,
		Types: []types.Type{types.Grass},
		HP:    100, Attack: 25, Defense: 85, SpAttack: 45, SpDefense: 115, Speed: 15,
		AttackRank: 0, DefenseRank: 0, SpAttackRank: 0, SpDefenseRank: 0, SpeedRank: 0,
		Weight: 100.0,
	}

	atItem := &item.NoItem{}
	dfItem := &item.NoItem{}

	atAbility := &AbilityDetail.NoEffectBuilder{}
	dfAbility := &AbilityDetail.NoEffectBuilder{}

	st := &situation.SituationData{
		Move:            s,
		Attacker:        at,
		Defender:        df,
		Weather:         field.NoWeather,
		Field:           field.NoField,
		AttackerAbility: atAbility,
		DefenderAbility: dfAbility,
		AttackerItem:    atItem,
		DefenderItem:    dfItem,
		IsCritical:      false,
	}
	return st
}
func calcDamage(sd *situation.SituationData) *Damages {
	d := NewDamageCalculator()
	st, _ := sd.Create()
	return d.CreateDamage(st)
}

func Test_重さ(t *testing.T) {
	a := defaultArgs()
	a.AttackerItem = &item.WeightCorrectData{2.0}
	a.DefenderItem = &item.WeightCorrectData{0.5}
	st, _ := a.Create()
	atWeight := st.Attacker().Weight()
	dfWeight := st.Defender().Weight()
	if atWeight != 200.0 {
		t.Errorf("%f", atWeight)
	}
	if dfWeight != 50.0 {
		t.Errorf("%f", dfWeight)
	}
}
func Test_ヘビーボンバー(t *testing.T) {
	d := NewDamageCalculator()
	a := defaultArgs()
	a.Move = &move.MoveFactory{
		Name:      "unknown",
		Power:     1,
		Type:      types.Steel,
		Accuracy:  100,
		Category:  category.Physical,
		CountMin:  1,
		CountMax:  1,
		Target:    target.Select,
		Attribute: attribute.NewAttribute(attribute.Contact, attribute.NoAttribute),
		Detail:    detail.HeavySlam,
	}
	st, _ := a.Create()
	dmgs := d.CreateDamage(st)
	if dmgs.Min() != 33 {
		t.Errorf("%v", dmgs)
	}
	// 重さがダメージに影響を与えること
	a.AttackerItem = &item.WeightCorrectData{5.0}
	st, _ = a.Create()
	dmgs = d.CreateDamage(st)
	if dmgs.Min() != 96 {
		t.Errorf("%v", dmgs)
	}
}

func Test_もちもの補正(t *testing.T) {
	d := NewDamageCalculator()
	a := defaultArgs()
	a.Attacker.Types = []types.Type{types.Fire}
	a.Defender.Types = []types.Type{types.Bug}
	a.Move.Type = types.Normal
	a.AttackerItem = &item.TypeCorrectData{types.Bug, 1.5}
	st, _ := a.Create()
	dmgs := d.CreateDamage(st)
	if dmgs.Min() != 64 {
		t.Errorf("%v", dmgs)
	}
	// 補正を掛けること
	a.AttackerItem = &item.TypeCorrectData{types.Normal, 1.5}
	st, _ = a.Create()
	dmgs = d.CreateDamage(st)
	if dmgs.Min() != 96 {
		t.Errorf("%v", dmgs)
	}
	// 防御側の持ち物補正が有効であること
	a.DefenderItem = &item.StatsCorrectData{Defense: 2.0, SpDefense: 2.0}
	st, _ = a.Create()
	dmgs = d.CreateDamage(st)
	if dmgs.Min() != 48 {
		t.Errorf("%v", dmgs)
	}
}

func Test_タイプ相性(t *testing.T) {
	d := NewDamageCalculator()
	a := defaultArgs()
	a.Attacker.Types = []types.Type{types.Fire}
	a.Defender.Types = []types.Type{types.Bug}
	a.Move.Type = types.Normal

	// ノーマル->むし
	st, _ := a.Create()
	dmgs := d.CreateDamage(st)
	if dmgs.Min() != 64 {
		t.Errorf("%v", dmgs)
	}
	// いわ->むし
	a.Move.Type = types.Rock
	st, _ = a.Create()
	dmgs = d.CreateDamage(st)
	if dmgs.Min() != 128 {
		t.Errorf("%v", dmgs)
	}
	// ほのお(タイプ一致)->むし
	a.Move.Type = types.Fire
	st, _ = a.Create()
	dmgs = d.CreateDamage(st)
	if dmgs.Min() != 192 {
		t.Errorf("%v", dmgs)
	}
}

func Test_急所(t *testing.T) {
	a := defaultArgs()

	dmgs := calcDamage(a)
	if dmgs.Min() != 128 {
		t.Errorf("%v", dmgs)
	}

	a.IsCritical = true
	dmgs = calcDamage(a)
	if dmgs.Min() != 192 {
		t.Errorf("%v", dmgs)
	}
}

// 天候補正時、1.5倍
func Test_天候(t *testing.T) {
	a := defaultArgs()

	a.Move.Type = types.Fire
	a.Defender.Types = []types.Type{types.Normal}
	a.Attacker.Types = []types.Type{types.Normal}

	dmgs := calcDamage(a)
	if dmgs.Min() != 64 {
		t.Errorf("%v", dmgs)
	}

	a.Weather = field.Sunny
	dmgs = calcDamage(a)
	if dmgs.Min() != 96 {
		t.Errorf("%v", dmgs)
	}
}

// フィールド補正時、1.3倍
func Test_フィールド(t *testing.T) {
	a := defaultArgs()

	a.Move.Type = types.Electric
	a.Defender.Types = []types.Type{types.Normal}
	a.Attacker.Types = []types.Type{types.Normal}

	dmgs := calcDamage(a)
	if dmgs.Min() != 64 {
		t.Errorf("%v", dmgs)
	}

	a.Field = field.ElectricField
	dmgs = calcDamage(a)
	if dmgs.Min() != 83 {
		t.Errorf("%v", dmgs)
	}
}

func Test_すなあらし岩タイプ補正(t *testing.T) {
	a := defaultArgs()

	a.Move.Type = types.Ghost
	a.Move.Category = category.Special
	a.Defender.Types = []types.Type{types.Rock}
	a.Attacker.Types = []types.Type{types.Normal}

	dmgs := calcDamage(a)
	if dmgs.Min() != 27 {
		t.Errorf("%v", dmgs)
	}

	// すなあらし時、とくぼう1.5倍
	a.Weather = field.SandStorm
	dmgs = calcDamage(a)
	if dmgs.Min() != 18 {
		t.Errorf("%v", dmgs)
	}
}

func Test_Damages_ソートされていること(t *testing.T) {
	d := NewDamages([]uint{5, 4, 2, 10, 8, 3, 1, 6})

	if d.Min() != 1 {
		t.Error()
	}
	if d.Max() != 10 {
		t.Error()
	}
}
