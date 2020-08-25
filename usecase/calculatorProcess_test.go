package usecase

import (
	"damagecalculator/domain/ability"
	"damagecalculator/domain/item"
	"damagecalculator/domain/move"
	"damagecalculator/domain/species"
	"damagecalculator/infra/local"
	"testing"
)

func Test_A(t *testing.T) {
	proc := NewCalculatorProcess(repositories())
	abilities := proc.AttackerAbilities()
	if len(abilities) != 0 {
		t.Error()
	}
	proc.AttackerName("ピカチュウ")
	abilities = proc.AttackerAbilities()
	if len(abilities) == 0 {
		t.Error()
	}
	t.Errorf("%v", abilities)

	proc.DefenderName("メガヤンマ")
	abilities = proc.DefenderAbilities()
	if len(abilities) == 0 {
		t.Error()
	}
	t.Errorf("%v", abilities)

	moves := proc.Moves()
	t.Errorf("%v", moves)

}

func repositories() (mv move.Repository, sp species.Repository, ab ability.Repository, it item.Repository) {
	return local.Move(), local.Species(), local.Ability(), local.Item()
}
