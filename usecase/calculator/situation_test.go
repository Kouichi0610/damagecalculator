package calculator

import (
	"damagecalculator/domain/ability"
	"damagecalculator/domain/damage"
	"damagecalculator/domain/field"
	"damagecalculator/domain/item"
	"damagecalculator/domain/move"
	"damagecalculator/domain/pokenames"
	"damagecalculator/domain/species"
	"damagecalculator/domain/stats"
	"damagecalculator/infra/local"
	"testing"
)

// ダメージ計算まで行えること
func Test_Factory(t *testing.T) {
	builder := NewBuilder(repositories())
	st, err := builder.Build()
	if st != nil || err == nil {
		t.Error()
	}

	pb := builder.NewPokeDataBuilder()
	list := pb.PokeList()
	if len(list) == 0 {
		t.Error()
	}
	pb.SetName("ピカチュウ")
	abilities := pb.AbilityList()
	if len(abilities) == 0 {
		t.Error()
	}
	pb.SetAbility(abilities[0])
	pb.SetLevel(50)
	pb.SetNature(stats.Jolly)
	pb.SetBasePoints(6, 252, 0, 0, 0, 252)
	pb.SetRanks(0, 0, 0, 0, 0)
	pb.SetIndividual(AllMax)
	//pb.ItemList()
	builder.SetAttacker(pb)

	pb = builder.NewPokeDataBuilder()
	pb.SetName("ゼニガメ")
	abilities = pb.AbilityList()
	pb.SetAbility(abilities[0])
	pb.SetLevel(50)
	pb.SetNature(stats.Bold)
	pb.SetBasePoints(252, 0, 252, 0, 6, 0)
	pb.SetRanks(0, 0, 0, 0, 0)
	pb.SetIndividual(AllMax)
	builder.SetDefender(pb)

	moves := builder.MoveList()
	if len(moves) == 0 {
		t.Error()
	}
	builder.SetMove("かみなりパンチ")

	builder.SetFields(field.NoWeather, field.ElectricField)
	builder.SetOptions(true, false, false, false)

	situation, err := builder.Build()
	if err != nil {
		t.Error()
	}
	if situation == nil {
		t.Error()
	}
	cl := damage.NewDamageCalculator()
	dmgs, rate := cl.CreateDamage(situation)
	if dmgs == nil {
		t.Error()
	}
	if rate == nil {
		t.Error()
	}

	if !(dmgs.Min() == 139 && dmgs.Max() == 168) {
		t.Errorf("%v", dmgs)
	}
	if rate.DetermineCount() != 2 {
		t.Errorf("%s", rate.String())
	}

	// Attackerを更新するとMoveが無効になること
	builder.SetAttacker(pb)
	situation, err = builder.Build()
	if err == nil {
		t.Error()
	}
}

func repositories() (pk pokenames.Repository, mv move.Repository, sp species.Repository, ab ability.Repository, it item.Repository) {
	return local.Names(), local.Move(), local.Species(), local.Ability(), local.Item()
}
