package ability

import (
	"damagecalculator/domain/ability/situation"
	"damagecalculator/domain/field"
	"damagecalculator/domain/status"
	"damagecalculator/domain/types"
)

// ぎたい
// フィールドに合わせてタイプが変わる
type mimicry struct {
	ability
}

func (s *mimicry) CorrectStatus(st situation.SituationChecker) *status.StatsCorrectors {
	c := status.NewStatsCorrectors()
	switch {
	case st.IsField(field.ElectricField):
		c.Types([]types.Type{types.Electric})
	case st.IsField(field.PsychoField):
		c.Types([]types.Type{types.Psychic})
	case st.IsField(field.MystField):
		c.Types([]types.Type{types.Fairy})
	case st.IsField(field.GrassField):
		c.Types([]types.Type{types.Grass})
	}
	return c
}
