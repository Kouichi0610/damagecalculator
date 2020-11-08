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
)

// 仮想敵のダメージ結果リスト作成
// virtual?
//TODO:Domain層に作っておくべき

/*
	TODO:仮想敵リストをDomainに

	args types, moves


	戦える相手
	戦えない相手

*/

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
	defenders, err := maker.Defenders(level, attacker.Name, move)
	if err == nil {
		return nil
	}
	rateService := NewService(s.sp, s.mv, s.ab, s.it)

	if rateService == nil || defenders == nil {
	}
	/*
		for _, d := range defenders {
			damages, rates, err := rateService.Calculate(level, attacker, attacker, move, condition)

		}
	*/
	return nil
}

// attackertypes -> (pokedata)[]

// desc, pokedata
/*
	Normal Type = iota
	Fire
	Water
	Electric
	Grass
	Ice
	Fighting
	Poison
	Ground
	Flying
	Psychic
	Bug
	Rock
	Ghost
	Dragon
	Dark
	Steel
	Fairy
*/
