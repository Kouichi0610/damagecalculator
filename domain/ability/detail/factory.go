package detail

import (
	"damagecalculator/domain/ability/correct"
	"damagecalculator/domain/field"
	"damagecalculator/domain/types"
	_ "damagecalculator/domain/types"
)

type (
	// とくせいを生成するためのインターフェイス
	AbilityBuilder interface {
		Create() ability
	}

	// implements AbilityBuilder
	NoEffectBuilder        struct{}
	MoldBreakderBuilder    struct{}
	NewtraizingGasBuilder  struct{}
	ForecastBuilder        struct{}
	MimicryBuilder         struct{}
	ProteanBuilder         struct{}
	SkillLinkBuilder       struct{}
	StatusCorrectorBuilder struct {
		Types                                               []types.Type
		Attack, Defense, SpAttack, SpDefense, Speed, Weight float64
	}
	WeatherStatusCorrectorBuilder struct {
		Weather                                             field.Weather
		Types                                               []types.Type
		Attack, Defense, SpAttack, SpDefense, Speed, Weight float64
	}
	FieldStatusCorrectorBuilder struct {
		Field                                               field.Field
		Types                                               []types.Type
		Attack, Defense, SpAttack, SpDefense, Speed, Weight float64
	}
	WonderGuardBuilder struct{}

	PowerCorrectorBuilder struct {
		Builders []correct.PowerCorrectorBuilder
	}
)

func (b *NoEffectBuilder) Create() ability {
	return new(abilityImpl)
}
func (b *MoldBreakderBuilder) Create() ability {
	return new(moldBreaker)
}
func (b *NewtraizingGasBuilder) Create() ability {
	return new(newtralizingGas)
}
func (b *ForecastBuilder) Create() ability {
	return new(forecast)
}
func (b *MimicryBuilder) Create() ability {
	return new(mimicry)
}
func (b *ProteanBuilder) Create() ability {
	return new(protean)
}
func (b *SkillLinkBuilder) Create() ability {
	return new(skillLink)
}
func (b *StatusCorrectorBuilder) Create() ability {
	return &statusCorrector{
		ty: b.Types,
		at: b.Attack,
		df: b.Defense,
		sa: b.SpAttack,
		sd: b.SpDefense,
		sp: b.Speed,
		wt: b.Weight,
	}
}
func (b *WeatherStatusCorrectorBuilder) Create() ability {
	st := statusCorrector{
		ty: b.Types,
		at: b.Attack,
		df: b.Defense,
		sa: b.SpAttack,
		sd: b.SpDefense,
		sp: b.Speed,
		wt: b.Weight,
	}
	return &weatherStatusCorrector{
		statusCorrector: st,
		weather:         b.Weather,
	}
}

func (b *FieldStatusCorrectorBuilder) Create() ability {
	st := statusCorrector{
		ty: b.Types,
		at: b.Attack,
		df: b.Defense,
		sa: b.SpAttack,
		sd: b.SpDefense,
		sp: b.Speed,
		wt: b.Weight,
	}
	return &fieldStatusCorrector{
		statusCorrector: st,
		fl:              b.Field,
	}
}

func (d *WonderGuardBuilder) Create() ability {
	return &wonderGuard{}
}

func (d *PowerCorrectorBuilder) Create() ability {
	res := &powerCorrector{
		c: make([]correct.PowerCorrector, 0),
	}
	for _, c := range d.Builders {
		res.c = append(res.c, c.Create())
	}
	return res
}
