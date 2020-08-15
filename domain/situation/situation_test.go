package situation

import (
	"damagecalculator/domain/ability"
	"damagecalculator/domain/ability/correct"
	"damagecalculator/domain/corrector"
	"damagecalculator/domain/field"
	"damagecalculator/domain/item"
	"damagecalculator/domain/skill"
	"damagecalculator/domain/skill/category"
	"damagecalculator/domain/status"
	"damagecalculator/domain/types"
	"testing"
)

func Test_NewSituation(t *testing.T) {
	d := NewSituationData()
	d.Attacker = &status.StatusData{50, []types.Type{types.Water}, 100, 100, 100, 100, 100, 100, 0, 0, 0, 0, 0, 20.0}
	d.Defender = &status.StatusData{50, []types.Type{types.Grass}, 100, 100, 100, 100, 100, 100, 0, 0, 0, 0, 0, 20.0}
	d.Skill = &skill.SkillData{[]types.Type{types.Fire}, 20, 2, 5, category.Physical, skill.NoMethod, skill.Knuckle, skill.NoAttribute}

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
	d.Skill = &skill.SkillData{[]types.Type{types.Fire}, 20, 2, 5, category.Physical, skill.NoMethod, skill.Knuckle, skill.NoAttribute}
	d.AttackerAbility = &ability.StatusCorrectorData{[]types.Type{types.Dragon}, 2.0, 1.0, 1.0, 1.0, 1.0, 2.0}
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
	d.Skill = &skill.SkillData{[]types.Type{types.Fire}, 20, 2, 5, category.Physical, skill.NoMethod, skill.Knuckle, skill.NoAttribute}
	d.AttackerAbility = &ability.PowerCorrectorData{
		[]correct.PowerCorrectorBuilder{
			&correct.TypeAttackData{[]types.Type{types.Fire}, 2.0},
		},
	}
	d.DefenderAbility = &ability.PowerCorrectorData{
		[]correct.PowerCorrectorBuilder{
			&correct.ActionDefenseData{skill.Knuckle, 0.5},
		},
	}
	d.AttackerItem = &item.TypeCorrectData{types.Fire, 3.0}
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
	if pwrs != 2 {
		t.Error()
	}
}

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

	atAbility := &ability.NoAbilityData{}
	dfAbility := &ability.NoAbilityData{}

	d := &SituationData{
		Attacker:        at,
		Defender:        df,
		Skill:           s,
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

	sk := st.Skill()
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
	if st.SkillAction() != skill.Knuckle {
		t.Error()
	}
	if !st.SkillEffective().IsSuper() {
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
