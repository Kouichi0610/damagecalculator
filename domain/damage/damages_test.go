package damage

import (
	"damagecalculator/domain/field"
	"damagecalculator/domain/situation"
	"damagecalculator/domain/skill"
	"damagecalculator/domain/skill/category"
	"damagecalculator/domain/status"
	"damagecalculator/domain/types"
	"testing"
)

func defaultArgs() *situation.SituationData {
	s := &skill.SkillData{
		Types:     []types.Type{types.Fire},
		Power:     80,
		CountMin:  1,
		CountMax:  1,
		Category:  category.Physical,
		Method:    skill.NoMethod,
		Action:    skill.Contact,
		Attribute: skill.NoAttribute,
	}
	at := &status.StatusData{
		Lv:    50,
		Types: []types.Type{types.Bug},
		HP:    100, Attack: 180, Defense: 20, SpAttack: 100, SpDefense: 50, Speed: 130,
		AttackRank: 0, DefenseRank: 0, SpAttackRank: 0, SpDefenseRank: 0, SpeedRank: 0,
	}
	df := &status.StatusData{
		Lv:    50,
		Types: []types.Type{types.Grass},
		HP:    100, Attack: 25, Defense: 85, SpAttack: 45, SpDefense: 115, Speed: 15,
		AttackRank: 0, DefenseRank: 0, SpAttackRank: 0, SpDefenseRank: 0, SpeedRank: 0,
	}

	st := &situation.SituationData{
		Skill:      s,
		Attacker:   at,
		Defender:   df,
		Weather:    field.NoWeather,
		Field:      field.NoField,
		IsCritical: false,
	}
	return st
}
func calcDamage(sd *situation.SituationData) *Damages {
	d := NewDamageCalculator()
	st, _ := sd.Create()
	return d.CreateDamage(st)
}

func Test_タイプ相性(t *testing.T) {
	d := NewDamageCalculator()
	a := defaultArgs()
	a.Attacker.Types = []types.Type{types.Fire}
	a.Defender.Types = []types.Type{types.Bug}
	a.Skill.Types = []types.Type{types.Normal}

	// ノーマル->むし
	st, _ := a.Create()
	dmgs := d.CreateDamage(st)
	if dmgs.Min() != 64 {
		t.Errorf("%v", dmgs)
	}
	// いわ->むし
	a.Skill.Types = []types.Type{types.Rock}
	st, _ = a.Create()
	dmgs = d.CreateDamage(st)
	if dmgs.Min() != 128 {
		t.Errorf("%v", dmgs)
	}
	// ほのお(タイプ一致)->むし
	a.Skill.Types = []types.Type{types.Fire}
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
func Test_天候(t *testing.T) {
	a := defaultArgs()

	a.Skill.Types = []types.Type{types.Fire}
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

func Test_フィールド(t *testing.T) {
	a := defaultArgs()

	a.Skill.Types = []types.Type{types.Electric}
	a.Defender.Types = []types.Type{types.Normal}
	a.Attacker.Types = []types.Type{types.Normal}

	dmgs := calcDamage(a)
	if dmgs.Min() != 64 {
		t.Errorf("%v", dmgs)
	}

	a.Field = field.ElectricField
	dmgs = calcDamage(a)
	if dmgs.Min() != 96 {
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
