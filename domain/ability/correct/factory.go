package correct

import (
	"damagecalculator/domain/skill"
	"damagecalculator/domain/types"
)

type PowerCorrectorBuilder interface {
	Create() PowerCorrector
}

type TypeAttackData struct {
	Types []types.Type
	Scale float64
}
type ActionAttackData struct {
	Action skill.Action
	Scale  float64
}
type TypeDefenseData struct {
	Types []types.Type
	Scale float64
}
type ActionDefenseData struct {
	Action skill.Action
	Scale  float64
}

func (d *TypeAttackData) Create() PowerCorrector {
	return &typeAttack{
		ty: d.Types,
		sc: d.Scale,
	}
}
func (d *ActionAttackData) Create() PowerCorrector {
	return &actionAttack{
		ac: d.Action,
		sc: d.Scale,
	}
}
func (d *TypeDefenseData) Create() PowerCorrector {
	return &typeDefense{
		ty: d.Types,
		sc: d.Scale,
	}
}
func (d *ActionDefenseData) Create() PowerCorrector {
	return &actionDefense{
		ac: d.Action,
		sc: d.Scale,
	}
}
