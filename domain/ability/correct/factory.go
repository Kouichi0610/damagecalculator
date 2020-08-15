package correct

import (
	"damagecalculator/domain/ability/situation"
	"damagecalculator/domain/corrector"
	"damagecalculator/domain/skill"
	"damagecalculator/domain/types"
)

type (
	/*
		ダメージ、威力に補正を掛ける
	*/
	PowerCorrector interface {
		Correct(isAttacker bool, st situation.SituationChecker) corrector.Corrector
	}

	// PowerCorrector生成インターフェイス
	PowerCorrectorBuilder interface {
		Create() PowerCorrector
	}

	// PowerCorrectorを生成するためのデータクラス
	TypeAttackData struct {
		Types []types.Type
		Scale float64
	}
	ActionAttackData struct {
		Action skill.Action
		Scale  float64
	}
	TypeDefenseData struct {
		Types []types.Type
		Scale float64
	}
	ActionDefenseData struct {
		Action skill.Action
		Scale  float64
	}
)

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
