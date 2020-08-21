package correct

import (
	"damagecalculator/domain/ability/situation"
	"damagecalculator/domain/corrector"
	"damagecalculator/domain/types"
)

// 特定のタイプのダメージ補正(防御側)
type typeDefense struct {
	ty []types.Type
	sc float64
}

func (s *typeDefense) Correct(isAttacker bool, st situation.SituationChecker) corrector.Corrector {
	if isAttacker {
		return nil
	}
	ty := st.MoveTypes()
	for _, t := range s.ty {
		if ty.Has(t) {
			return corrector.NewDamage(s.sc)
		}
	}
	return nil
}
