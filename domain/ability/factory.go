package ability

import (
	"damagecalculator/domain/ability/correct"
	"damagecalculator/domain/field"
	"damagecalculator/domain/types"
)

type AbilityBuilder interface {
	Create() Ability
}

type PowerCorrectorData struct {
	Builders []correct.PowerCorrectorBuilder
}

type NoAbilityData struct {
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

func (d *NoAbilityData) Create() Ability {
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

func (d *SkillLinkData) Create() Ability {
	return &skillLink{}
}

func (d *PowerCorrectorData) Create() Ability {
	res := &powerCorrector{
		c: make([]correct.PowerCorrector, 0),
	}
	for _, c := range d.Builders {
		res.c = append(res.c, c.Create())
	}
	return res
}
