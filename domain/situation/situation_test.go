package situation

import (
	"damagecalculator/domain/ability/correct"
	AbilityDetail "damagecalculator/domain/ability/detail"
	"damagecalculator/domain/corrector"
	"damagecalculator/domain/field"
	"damagecalculator/domain/item"
	"damagecalculator/domain/move"
	"damagecalculator/domain/move/attribute"
	"damagecalculator/domain/move/category"
	"damagecalculator/domain/move/detail"
	"damagecalculator/domain/move/target"
	"damagecalculator/domain/status"
	"damagecalculator/domain/types"
	"testing"
)

func Test_NewSituation(t *testing.T) {
	d := NewSituationData()
	d.Attacker = &status.StatusData{50, []types.Type{types.Water}, 100, 100, 100, 100, 100, 100, 0, 0, 0, 0, 0, 20.0}
	d.Defender = &status.StatusData{50, []types.Type{types.Grass}, 100, 100, 100, 100, 100, 100, 0, 0, 0, 0, 0, 20.0}
	d.Move = testMoveFactory()

	st, err := d.Create()
	if st == nil {
		t.Error()
	}
	if err != nil {
		t.Error()
	}
}
func Test_ステータス補正(t *testing.T) {
	d := NewSituationData()
	d.Attacker = &status.StatusData{50, []types.Type{types.Water}, 100, 100, 100, 100, 100, 100, 0, 0, 0, 0, 0, 20.0}
	d.Defender = &status.StatusData{50, []types.Type{types.Grass}, 100, 100, 100, 100, 100, 100, 0, 0, 0, 0, 0, 20.0}
	d.Move = testMoveFactory()
	d.AttackerAbility = &AbilityDetail.StatusCorrectorBuilder{[]types.Type{types.Dragon}, 2.0, 1.0, 1.0, 1.0, 1.0, 2.0}
	st, _ := d.Create()

	at := st.Attacker()

	if at.Attack().Value() != 200 {
		t.Error()
	}
	if at.Weight() != 40.0 {
		t.Error()
	}
	if !at.Types().Has(types.Dragon) {
		t.Error()
	}
}

func Test_ダメージ威力補正(t *testing.T) {
	d := NewSituationData()
	d.Attacker = &status.StatusData{50, []types.Type{types.Water}, 100, 100, 100, 100, 100, 100, 0, 0, 0, 0, 0, 20.0}
	d.Defender = &status.StatusData{50, []types.Type{types.Grass}, 100, 100, 100, 100, 100, 100, 0, 0, 0, 0, 0, 20.0}
	d.Move = testMoveFactory()
	d.AttackerAbility = &AbilityDetail.PowerCorrectorBuilder{
		[]correct.PowerCorrectorBuilder{
			&correct.TypeAttackData{[]types.Type{types.Fire}, 2.0},
		},
	}
	d.DefenderAbility = &AbilityDetail.PowerCorrectorBuilder{
		[]correct.PowerCorrectorBuilder{
			&correct.ActionDefenseData{attribute.Knuckle, 0.5},
		},
	}
	d.AttackerItem = &item.TypeCorrectData{types.Fire, 3.0}
	// 防御側に持たせても効果ない事
	d.DefenderItem = &item.TypeCorrectData{types.Fire, 2.9}
	st, _ := d.Create()

	c := st.Correctors()

	dmgs := 0
	pwrs := 0
	corrected := make([]uint, 0)
	for _, cc := range c {
		if cc.Category() == corrector.Damage {
			dmgs++
		}
		if cc.Category() == corrector.Power {
			pwrs++
		}
		corrected = append(corrected, cc.Correct(100))
	}
	if dmgs != 1 {
		t.Error()
	}
	if pwrs != 3 {
		t.Errorf("%d", pwrs)
	}
}

func testMoveFactory() *move.MoveFactory {
	return &move.MoveFactory{
		Name:      "unknown",
		Power:     20,
		Type:      types.Fire,
		Accuracy:  100,
		Category:  category.Physical,
		CountMin:  2,
		CountMax:  5,
		Target:    target.Select,
		Attribute: attribute.NewAttribute(attribute.Knuckle, attribute.NoAttribute),
		Detail:    detail.Default,
	}
}

func create() (SituationChecker, error) {
	s := &move.MoveFactory{
		Name:      "unknown",
		Power:     120,
		Type:      types.Fire,
		Accuracy:  100,
		Category:  category.Physical,
		CountMin:  1,
		CountMax:  1,
		Target:    target.Select,
		Attribute: attribute.NewAttribute(attribute.Knuckle, attribute.NoAttribute),
		Detail:    detail.Default,
	}
	at := &status.StatusData{
		Lv:    50,
		Types: []types.Type{types.Water},
		HP:    150, Attack: 120, Defense: 80, SpAttack: 100, SpDefense: 75, Speed: 20,
		AttackRank: 0, DefenseRank: 0, SpAttackRank: 0, SpDefenseRank: 0, SpeedRank: 0,
		Weight: 100.0,
	}
	df := &status.StatusData{
		Lv:    50,
		Types: []types.Type{types.Grass},
		HP:    200, Attack: 55, Defense: 130, SpAttack: 65, SpDefense: 160, Speed: 140,
		AttackRank: 0, DefenseRank: 0, SpAttackRank: 0, SpDefenseRank: 0, SpeedRank: 0,
		Weight: 100.0,
	}

	atItem := &item.NoItem{}
	dfItem := &item.NoItem{}

	atAbility := &AbilityDetail.NoEffectBuilder{}
	dfAbility := &AbilityDetail.NoEffectBuilder{}

	d := &SituationData{
		Attacker:        at,
		Defender:        df,
		Move:            s,
		Weather:         field.Sunny,
		Field:           field.ElectricField,
		AttackerAbility: atAbility,
		DefenderAbility: dfAbility,
		AttackerItem:    atItem,
		DefenderItem:    dfItem,
		IsCritical:      false,
	}
	return d.Create()
}

func Test_Situation機能(t *testing.T) {
	st, err := create()
	if err != nil {
		t.Error()
	}
	at := st.Attacker()
	if at.Attack().Value() != 120 {
		t.Error()
	}
	df := st.Defender()
	if df.SpAttack().Value() != 65 {
		t.Error()
	}

	if at.Weight() != 100.0 {
		t.Errorf("%f", at.Weight())
	}
	if df.Weight() != 100.0 {
		t.Errorf("%f", at.Weight())
	}

	sk := st.Move()
	if !sk.Types(st).Has(types.Fire) {
		t.Error()
	}

	if !st.IsWeather(field.Sunny) {
		t.Error()
	}

	if !st.IsField(field.ElectricField) {
		t.Error()
	}

	if st.IsCritical() {
		t.Error()
	}
	attr := st.MoveAttribute()
	if !attr.HasAction(attribute.Knuckle) {
		t.Error()
	}
	if !st.MoveEffective().IsSuper() {
		t.Error()
	}
}

// Skillインターフェイスに渡しても問題なく動作すること
func Test_Situation_for_Skill(t *testing.T) {
	s, _ := testMoveFactory().Create()
	st, _ := create()

	if s.Power(st) != 20 {
		t.Error()
	}
}
