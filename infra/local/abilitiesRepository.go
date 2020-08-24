package local

import (
	"damagecalculator/domain/ability"
	"damagecalculator/domain/ability/correct"
	D "damagecalculator/domain/ability/detail"
	"damagecalculator/domain/field"
	"damagecalculator/domain/types"
)

// とくせいを取得してAbilityFieldを生成する
type abilityRepository struct {
}

func (r *abilityRepository) Get(at, df string) ability.AbilityField {
	a := newAbilityBuilder(at)
	d := newAbilityBuilder(df)
	return ability.NewAbilityField(a, d)
}

/*
	該当するAbilityBuilderを取得。
	該当しなければNoEffectを返す
*/
func newAbilityBuilder(name string) D.AbilityBuilder {
	res, ok := abilityMap[name]
	if !ok {
		return &D.NoEffectBuilder{}
	}
	return res
}

var abilityMap map[string]D.AbilityBuilder = map[string]D.AbilityBuilder{
	"かたやぶり":    &D.MoldBreakderBuilder{},
	"ぎたい":      &D.MimicryBuilder{},
	"へんげんじざい":  &D.ProteanBuilder{},
	"リベロ":      &D.ProteanBuilder{},
	"かがくへんかガス": &D.NewtraizingGasBuilder{},
	"ふしぎなまもり":  &D.WonderGuardBuilder{},
	"スキルリンク":   &D.SkillLinkBuilder{},
	"てんきや":     &D.ForecastBuilder{},
	"サーフテール":   &D.FieldStatusCorrectorBuilder{Field: field.ElectricField, Speed: 2.0},
	"ようりょくそ":   &D.WeatherStatusCorrectorBuilder{Weather: field.Sunny, Speed: 2.0},
	"すいすい":     &D.WeatherStatusCorrectorBuilder{Weather: field.Rainy, Speed: 2.0},
	"すいほう":     &D.PowerCorrectorBuilder{[]correct.PowerCorrectorBuilder{&correct.TypeAttackData{[]types.Type{types.Water}, 2.0}, &correct.TypeDefenseData{[]types.Type{types.Fire}, 0.5}}},
	"ちからもち":    &D.StatusCorrectorBuilder{Attack: 2.0},
	"ヨガパワー":    &D.StatusCorrectorBuilder{Attack: 2.0},
}
