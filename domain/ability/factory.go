package ability

import (
	"damagecalculator/domain/field"
	"damagecalculator/domain/skill"
	"damagecalculator/domain/types"
)

type AbilityBuilder interface {
	Create() Ability
}

/*
	TODO:
	1個のAbilityBuilder,
	[0...n]のPowerCorrectorBuilder

*/
type PowerCorrectorBuilder interface {
	Create() PowerCorrector
}

type PowerCorrectorData struct {
	Builders []PowerCorrectorBuilder
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
type ForecastData struct {
}
type MimicryData struct {
}
type SkillLinkData struct {
}

type TypePowerCorrectData struct {
	Types []types.Type
	Scale float64
}
type ActionPowerCorrectData struct {
	Action skill.Action
	Scale  float64
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
func (d *ForecastData) Create() Ability {
	return &forecast{}
}
func (d *MimicryData) Create() Ability {
	return &mimicry{}
}

func (d *TypePowerCorrectData) Create() PowerCorrector {
	return &typePowerCorrector{
		ty: d.Types,
		sc: d.Scale,
	}
}

func (d *ActionPowerCorrectData) Create() PowerCorrector {
	return &actionPowerCorrector{
		ac: d.Action,
		sc: d.Scale,
	}
}

func (d *SkillLinkData) Create() Ability {
	return &skillLink{}
}

func (d *PowerCorrectorData) Create() Ability {
	res := &powerCorrector{
		c: make([]PowerCorrector, 0),
	}
	for _, c := range d.Builders {
		res.c = append(res.c, c.Create())
	}
	return res
}
