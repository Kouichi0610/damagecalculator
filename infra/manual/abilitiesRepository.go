package manual

import (
	"damagecalculator/domain/ability"
	"damagecalculator/domain/ability/correct"
	"damagecalculator/domain/ability/detail"
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
func newAbilityBuilder(name string) detail.AbilityBuilder {
	res, ok := abilityMap[name]
	if !ok {
		return &detail.NoEffectBuilder{}
	}
	return res
}

// TODO:短く。とくにすいほう (生成メソッド用意)
var abilityMap map[string]detail.AbilityBuilder = map[string]detail.AbilityBuilder{
	"かたやぶり":    &detail.MoldBreakderBuilder{},
	"ぎたい":      &detail.MimicryBuilder{},
	"へんげんじざい":  &detail.ProteanBuilder{},
	"リベロ":      &detail.ProteanBuilder{},
	"かがくへんかガス": &detail.NewtraizingGasBuilder{},
	"ふしぎなまもり":  &detail.WonderGuardBuilder{},
	"スキルリンク":   &detail.SkillLinkBuilder{},
	"てんきや":     &detail.ForecastBuilder{},
	"サーフテール":   &detail.FieldStatusCorrectorBuilder{Field: field.ElectricField, Speed: 2.0},
	"ようりょくそ":   &detail.WeatherStatusCorrectorBuilder{Weather: field.Sunny, Speed: 2.0},
	"すいすい":     &detail.WeatherStatusCorrectorBuilder{Weather: field.Rainy, Speed: 2.0},
	"すいほう":     &detail.PowerCorrectorBuilder{[]correct.PowerCorrectorBuilder{&correct.TypeAttackData{[]types.Type{types.Water}, 2.0}, &correct.TypeDefenseData{[]types.Type{types.Fire}, 0.5}}},
	"ちからもち":    &detail.StatusCorrectorBuilder{Attack: 2.0},
	"ヨガパワー":    &detail.StatusCorrectorBuilder{Attack: 2.0},
}
