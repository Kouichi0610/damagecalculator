/*
	ダメージ計算に必要な場の状態
*/
package situation

import (
	"damagecalculator/domain/status"
)

type SituationChecker interface {
	Attacker() status.StatusChecker
	Defender() status.StatusChecker
}

func NewSituationChecker() SituationChecker {
	return &situation{}
}

type situation struct {
	at *status.Status
	df *status.Status
}

func (s *situation) Attacker() status.StatusChecker {
	return s.at
}
func (s *situation) Defender() status.StatusChecker {
	return s.df
}
