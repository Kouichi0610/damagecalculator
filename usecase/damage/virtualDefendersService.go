package damage

import (
	"damagecalculator/domain/ability"
	"damagecalculator/domain/damage"
	"damagecalculator/domain/item"
	"damagecalculator/domain/move"
	"damagecalculator/domain/pokenames"
	"damagecalculator/domain/situation"
	"damagecalculator/domain/species"
	"damagecalculator/domain/stats"
	"damagecalculator/domain/supposition"
	"fmt"
)

// 仮想敵のダメージ結果リスト作成

type (
	Result interface {
		Info() string
		Damages() damage.Damages
		Rate() damage.DamageRate
	}
	Results                []Result
	VirtualDefenderService interface {
		Create(level stats.Level, attacker *situation.PokeParams, move string, condition *situation.FieldCondition) Results
	}
)

type vdService struct {
	nm pokenames.Repository
	sp species.Repository
	mv move.Repository
	ab ability.Repository
	it item.Repository
}

func NewVirtualDefendersService(nm pokenames.Repository, sp species.Repository, mv move.Repository, ab ability.Repository, it item.Repository) VirtualDefenderService {
	return &vdService{
		nm: nm,
		sp: sp,
		mv: mv,
		ab: ab,
		it: it,
	}
}

func (s *vdService) Create(level stats.Level, attacker *situation.PokeParams, move string, condition *situation.FieldCondition) Results {
	maker := supposition.NewDefendersMaker(s.nm, s.sp, s.mv, s.ab, s.it)
	defenders, err := maker.Defenders(attacker.Name, move)
	if err != nil {
		return nil
	}
	rateService := damage.NewDamageService(s.sp, s.mv, s.ab, s.it)

	fmt.Printf("Attacker:%s\n", attacker.Name)

	for _, d := range defenders {
		damages, rates, err := rateService.Calculate(level, attacker, &d, move, condition)
		if err != nil {
			fmt.Printf("error.\n")
		}
		if damages != nil {
		}
		// TODO:
		fmt.Printf(" -> %s %s\n", d.Name, rates.String())
	}
	return nil
}
