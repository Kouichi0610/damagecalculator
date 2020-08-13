package situation

import (
	"damagecalculator/domain/field"
	"damagecalculator/domain/item"
	"damagecalculator/domain/skill"
	"damagecalculator/domain/skill/category"
	"damagecalculator/domain/status"
	"damagecalculator/domain/types"
	"testing"
)

func create() (SituationChecker, error) {
	s := &skill.SkillData{
		Types:     []types.Type{types.Fire},
		Power:     120,
		CountMin:  1,
		CountMax:  1,
		Category:  category.Physical,
		Method:    skill.NoMethod,
		Action:    skill.Knuckle,
		Attribute: skill.NoAttribute,
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

	d := &SituationData{
		Attacker:     at,
		Defender:     df,
		Skill:        s,
		Weather:      field.Sunny,
		Field:        field.ElectricField,
		AttackerItem: atItem,
		DefenderItem: dfItem,
		IsCritical:   false,
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

	sk := st.Skill()
	if !sk.Types(st).Has(types.Fire) {
		t.Error()
	}

	if !st.IsWeather(field.Sunny) {
		t.Error()
	}

	if st.IsCritical() {
		t.Error()
	}
}

// Skillインターフェイスに渡しても問題なく動作すること
func Test_Situation_for_Skill(t *testing.T) {
	s, _ := (&skill.SkillData{
		Types:     []types.Type{types.Fire},
		Power:     120,
		CountMin:  1,
		CountMax:  1,
		Category:  category.Physical,
		Method:    skill.NoMethod,
		Action:    skill.Knuckle,
		Attribute: skill.NoAttribute,
	}).Create()
	st, _ := create()

	if s.Power(st) != 120 {
		t.Error()
	}
}
