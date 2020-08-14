package ability

import (
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
