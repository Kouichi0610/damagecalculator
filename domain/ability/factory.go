package ability

import (
	"damagecalculator/domain/field"
	"damagecalculator/domain/types"
)

type AbilityBuilder interface {
	Create() Ability
}

type NoEffectData struct {
}
type MoldBreakerData struct {
}
type NewTralizingGasData struct {
}
type WonderGuardData struct {
}
type StatusCorrectorData struct {
	Types                                               []types.Type
	Attack, Defense, SpAttack, SpDefense, Speed, Weight float64
}
type FieldStatusCorrectData struct {
	Field                                               field.Field
	Types                                               []types.Type
	Attack, Defense, SpAttack, SpDefense, Speed, Weight float64
}
type WeatherStatusCorrectData struct {
	Weather                                             field.Weather
	Types                                               []types.Type
	Attack, Defense, SpAttack, SpDefense, Speed, Weight float64
}
type ProteanData struct {
}

func (d *NoEffectData) Create() Ability {
	return &ability{}
}
func (d *MoldBreakerData) Create() Ability {
	return &moldBreaker{}
}
func (d *NewTralizingGasData) Create() Ability {
	return &newtralizingGas{}
}
func (d *WonderGuardData) Create() Ability {
	return &wonderGuard{}
}
func (d *StatusCorrectorData) Create() Ability {
	return &statusCorrector{
		ty: d.Types,
		at: d.Attack,
		df: d.Defense,
		sa: d.SpAttack,
		sd: d.SpDefense,
		sp: d.Speed,
		wt: d.Weight,
	}
}
func (d *FieldStatusCorrectData) Create() Ability {
	st := statusCorrector{
		ty: d.Types,
		at: d.Attack,
		df: d.Defense,
		sa: d.SpAttack,
		sd: d.SpDefense,
		sp: d.Speed,
		wt: d.Weight,
	}
	return &fieldStatusCorrector{
		fl:              d.Field,
		statusCorrector: st,
	}
}
func (d *WeatherStatusCorrectData) Create() Ability {
	st := statusCorrector{
		ty: d.Types,
		at: d.Attack,
		df: d.Defense,
		sa: d.SpAttack,
		sd: d.SpDefense,
		sp: d.Speed,
		wt: d.Weight,
	}
	return &weatherStatusCorrector{
		wt:              d.Weather,
		statusCorrector: st,
	}
}

func (d *ProteanData) Create() Ability {
	return &protean{}
}
