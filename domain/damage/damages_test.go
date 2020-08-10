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

func Test_DamageCalculator(t *testing.T) {
	// TODO:テストの場合分けをどうするか
	sd := defaultArgs()
	st, _ := sd.Create()

	d := NewDamageCalculator()

	dmgs := d.CreateDamage(st)
	if len(dmgs.d) != 16 {
		t.Errorf("%d", len(dmgs.d))
	}
	t.Errorf("%v", dmgs.d)
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
